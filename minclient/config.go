package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/katzenpost/core/crypto/ecdh"
	"github.com/katzenpost/core/crypto/eddsa"
	"github.com/katzenpost/core/utils"
	"github.com/pelletier/go-toml"
)

const (
	defaultLogLevel         = "NOTICE"
)

var defaultLogging = Logging{
	Disable: false,
	File:    "",
	Level:   defaultLogLevel,
}

// Account is the Katzenpost user configuration.
type Account struct {
	// Name is the user name identifier to connect to the provider.
	Name string

	// Provider is the user's provider identifier.
	Provider string

	// PrivateKey is the user's ecdh identity key.
	PrivateKey string
}

func (aCfg *Account) validate() error {
	if aCfg.Name == "" {
		return fmt.Errorf("config: Account: Name is not set")
	}
	if aCfg.Provider == "" {
		return fmt.Errorf("config: Account: Provider is not set")
	}

	_, err := aCfg.getPrivateKey()
	if err != nil {
 		fmt.Errorf("config: Account: Invalid PrivateKey: %v", err)
	}

	return nil
}

func (aCfg *Account) getPrivateKey() (*ecdh.PrivateKey, error) {
	rawKey, err := hex.DecodeString(aCfg.PrivateKey)
	if err != nil {
		return nil, err
	}
	privateKey := new(ecdh.PrivateKey)
	err = privateKey.FromBytes(rawKey)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

// Logging is the Katzenpost server logging configuration.
type Logging struct {
	// Disable disables logging entirely.
	Disable bool

	// File specifies the log file, if omitted stdout will be used.
	File string

	// Level specifies the log level.
	Level string
}

func (lCfg *Logging) validate() error {
	lvl := strings.ToUpper(lCfg.Level)
	switch lvl {
	case "ERROR", "WARNING", "NOTICE", "INFO", "DEBUG":
	case "":
		lCfg.Level = defaultLogLevel
	default:
		return fmt.Errorf("config: Logging: Level '%v' is invalid", lCfg.Level)
	}
	lCfg.Level = lvl // Force uppercase.
	return nil
}

// PKI is the Katzenpost directory authority configuration.
type PKI struct {
	// Nonvoting is a non-voting directory authority.
	Nonvoting *Nonvoting
}

func (pCfg *PKI) validate() error {
	nrCfg := 0
	if pCfg.Nonvoting != nil {
		if err := pCfg.Nonvoting.validate(); err != nil {
			return err
		}
		nrCfg++
	}
	if nrCfg != 1 {
		return fmt.Errorf("config: Only one authority backend should be configured, got: %v", nrCfg)
	}
	return nil
}

// Nonvoting is a non-voting directory authority.
type Nonvoting struct {
	// Address is the authority's IP/port combination.
	Address string

	// PublicKey is the authority's public key in Base64 or Base16 format.
	PublicKey string
}

func (nCfg *Nonvoting) validate() error {
	if err := utils.EnsureAddrIPPort(nCfg.Address); err != nil {
		return fmt.Errorf("config: PKI/Nonvoting: Address is invalid: %v", err)
	}
	if _, err := nCfg.getPublicKey(); err != nil {
		return fmt.Errorf("config: PKI/Nonvoting: Invalid PublicKey: %v", err)
	}

	return nil
}

func (nCfg *Nonvoting) getPublicKey() (*eddsa.PublicKey, error) {
	var pubKey eddsa.PublicKey
	err := pubKey.FromString(nCfg.PublicKey)
	return &pubKey, err
}

// Config is the top level Katzenpost client configuration.
type Config struct {
	Account    *Account
	Logging    *Logging
	PKI        *PKI
}

// FixupAndValidate applies defaults to config entries and validates the
// supplied configuration.  Most people should call one of the Load variants
// instead.
func (cfg *Config) FixupAndValidate() error {
	// The User and PKI sections are mandatory, everything else is optional.
	if cfg.Account == nil {
		return errors.New("config: No Account block was present")
	}
	if cfg.Logging == nil {
		cfg.Logging = &defaultLogging
	}
	if cfg.PKI == nil {
		return errors.New("config: No PKI block was present")
	}

	// Perform basic validation.
	if err := cfg.Account.validate(); err != nil {
		return err
	}
	if err := cfg.PKI.validate(); err != nil {
		return err
	}
	if err := cfg.Logging.validate(); err != nil {
		return err
	}

	return nil
}

// Load parses and validates the provided buffer b as a config file body and
// returns the Config.
func Load(b []byte) (*Config, error) {
	cfg := new(Config)
	if err := toml.Unmarshal(b, cfg); err != nil {
		return nil, err
	}
	if err := cfg.FixupAndValidate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

// LoadFile loads, parses and validates the provided file and returns the
// Config.
func LoadFile(f string) (*Config, error) {
	b, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	return Load(b)
}
