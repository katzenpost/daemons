package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"os"

	"github.com/katzenpost/core/crypto/ecdh"
	"github.com/katzenpost/core/crypto/rand"
	cpki"github.com/katzenpost/core/pki"
	"github.com/katzenpost/core/sphinx"
	"github.com/katzenpost/core/sphinx/constants"
	"github.com/katzenpost/core/utils"
	npki "github.com/katzenpost/authority/nonvoting/client"
	"github.com/katzenpost/minclient"
	"github.com/katzenpost/minclient/block"
	"github.com/katzenpost/core/log"
)

var surbKeys = make(map[[constants.SURBIDLength]byte][]byte)

func main() {
	cfgFile := flag.String("f", "client.toml", "Path to the client config file.")
	genOnly := flag.Bool("g", false, "Generate the keys and exit immediately.")
	flag.Parse()

	if *genOnly {
		key, err := ecdh.NewKeypair(rand.Reader)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to to generate the key: %v\n", err)
			os.Exit(-1)
		}
		fmt.Printf("Private key: %v\n", hex.EncodeToString(key.Bytes()))
		fmt.Printf("Public key: %v\n", key.PublicKey().String())
		os.Exit(0)
	}

	cfg, err := LoadFile(*cfgFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load config file '%v': %v\n", *cfgFile, err)
		os.Exit(-1)
	}

	clientLog, err := newLog(cfg.Logging)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create client log: %v\n", err)
		os.Exit(-1)
	}

	pkiClient, err := newPKIClient(cfg.PKI.Nonvoting, clientLog)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create PKI client: %v\n", err)
		os.Exit(-1)
	}

	client, _, err := newMinClient(cfg, pkiClient, clientLog)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create min client: %v\n", err)
		os.Exit(-1)
	}

	client.Wait()
}

func newLog(cfg *Logging) (*log.Backend, error) {
	return log.New(cfg.File, cfg.Level, cfg.Disable)
}

func newPKIClient(cfg *Nonvoting, clientLog *log.Backend) (cpki.Client, error) {
	pubkey, err := cfg.getPublicKey()
	if err != nil {
		return nil, err
	}

	pkiCfg := npki.Config{
		LogBackend: clientLog,
		Address: cfg.Address,
		PublicKey: pubkey,
	}
	return npki.New(&pkiCfg)
}

func newMinClient(cfg *Config, pkiClient cpki.Client, clientLog *log.Backend) (*minclient.Client, chan interface{}, error) {
	lm := clientLog.GetLogger("callbacks:main")
	user := cfg.Account.Name
	provider := cfg.Account.Provider
	privateKey, _ := cfg.Account.getPrivateKey()
	onlineCh := make(chan interface{}, 1)

	clientCfg := &minclient.ClientConfig{
		User:        user,
		Provider:    provider,
		LinkKey: privateKey,
		LogBackend:  clientLog,
		PKIClient:   pkiClient,
		OnConnFn:    func(isConnected bool) {
			lm.Noticef("Peer connection status changed: %v", isConnected)
			select {
			case onlineCh <- isConnected:
			default:
			}
		},
		OnMessageFn: func(b []byte) error {
			lm.Noticef("Received Message: %v", len(b))

			blk, pk, err := block.DecryptBlock(b, privateKey)
			if err != nil {
				lm.Errorf("Failed to decrypt block: %v", err)
				return nil
			}

			lm.Noticef("Sender Public Key: %v", pk)
			lm.Noticef("Message payload: %v", hex.Dump(blk.Payload))

			return nil
		},
		OnACKFn: func(id *[constants.SURBIDLength]byte, b []byte) error {
			lm.Noticef("Received SURB-ACK: %v", len(b))
			lm.Noticef("SURB-ID: %v", hex.EncodeToString(id[:]))

			// surbKeys should have a lock in production code, but lazy.
			k, ok := surbKeys[*id]
			if !ok {
				lm.Errorf("Failed to find SURB SPRP key")
				return nil
			}

			payload, err := sphinx.DecryptSURBPayload(b, k)
			if err != nil {
				lm.Errorf("Failed to decrypt SURB: %v", err)
				return nil
			}
			if utils.CtIsZero(payload) {
				lm.Noticef("SURB Payload: %v bytes of 0x00", len(payload))
			} else {
				lm.Noticef("SURB Payload: %v", hex.Dump(payload))
			}

			return nil
		},
	}
	c, err := minclient.New(clientCfg)
	if err != nil {
		return nil, nil, err
	}

	return c, onlineCh, nil
}
