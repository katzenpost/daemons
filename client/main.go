// main.go - mixnet client
// Copyright (C) 2017  David Anthony Stainton
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Package main provides a mixnet client daemon
package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	"unsafe"

	"github.com/katzenpost/client/auth"
	"github.com/katzenpost/client/config"
	"github.com/katzenpost/client/constants"
	"github.com/katzenpost/client/crypto/block"
	"github.com/katzenpost/client/mix_pki"
	"github.com/katzenpost/client/path_selection"
	"github.com/katzenpost/client/proxy"
	"github.com/katzenpost/client/session_pool"
	"github.com/katzenpost/client/storage"
	"github.com/katzenpost/client/user_pki"
	"github.com/katzenpost/core/crypto/rand"
	"github.com/katzenpost/core/wire/server"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("mixclient")

var logFormat = logging.MustStringFormatter(
	"%{level:.4s} %{id:03x} %{message}",
)
var ttyFormat = logging.MustStringFormatter(
	"%{color}%{time:15:04:05} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}",
)

const ioctlReadTermios = 0x5401

func isTerminal(fd int) bool {
	var termios syscall.Termios
	_, _, err := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(fd), ioctlReadTermios, uintptr(unsafe.Pointer(&termios)), 0, 0, 0)
	return err == 0
}

func stringToLogLevel(level string) (logging.Level, error) {

	switch level {
	case "DEBUG":
		return logging.DEBUG, nil
	case "INFO":
		return logging.INFO, nil
	case "NOTICE":
		return logging.NOTICE, nil
	case "WARNING":
		return logging.WARNING, nil
	case "ERROR":
		return logging.ERROR, nil
	case "CRITICAL":
		return logging.CRITICAL, nil
	}
	return -1, fmt.Errorf("invalid logging level %s", level)
}

func setupLoggerBackend(level logging.Level) logging.LeveledBackend {
	format := logFormat
	if isTerminal(int(os.Stderr.Fd())) {
		format = ttyFormat
	}
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	formatter := logging.NewBackendFormatter(backend, format)
	leveler := logging.AddModuleLevel(formatter)
	leveler.SetLevel(level, "mixclient")
	return leveler
}

func main() {
	var err error
	var level logging.Level

	var configFilePath string
	var keysDirPath string
	var userPKIFile string
	var mixPKIFile string
	var dbFile string
	var logLevel string
	var shouldAutogenKeys bool

	flag.BoolVar(&shouldAutogenKeys, "autogenkeys", false, "auto-generate cryptographic keys specified in configuration file")
	flag.StringVar(&configFilePath, "config", "", "configuration file")
	flag.StringVar(&keysDirPath, "keysdir", "", "the path to the keys directory")
	flag.StringVar(&userPKIFile, "userpkifile", "", "user pki in a json file")
	flag.StringVar(&mixPKIFile, "mixpkifile", "", "consensus file path to use as the mixnet PKI")
	flag.StringVar(&dbFile, "dbfile", "", "incoming and outgoing message DB file path")
	flag.StringVar(&logLevel, "log_level", "INFO", "logging level could be set to: DEBUG, INFO, NOTICE, WARNING, ERROR, CRITICAL")
	flag.Parse()

	level, err = stringToLogLevel(logLevel)
	if err != nil {
		log.Critical("Invalid logging-level specified.")
		os.Exit(1)
	}
	logBackend := setupLoggerBackend(level)
	log.SetBackend(logBackend)

	passphrase := os.Getenv("MIX_CLIENT_VAULT_PASSPHRASE")
	if len(passphrase) == 0 {
		panic("Aborting because bash env var not set: MIX_CLIENT_VAULT_PASSPHRASE")
	}

	if configFilePath == "" {
		log.Error("you must specify a configuration file")
		flag.Usage()
		os.Exit(1)
	}

	if keysDirPath == "" {
		log.Error("you must specify a keys directory file path")
		flag.Usage()
		os.Exit(1)
	}

	if userPKIFile == "" {
		log.Error("you must specify a user-pki json file path")
		flag.Usage()
		os.Exit(1)
	}

	if mixPKIFile == "" {
		log.Error("you must specify a mixnet PKI consensus file path")
		flag.Usage()
		os.Exit(1)
	}

	sigKillChan := make(chan os.Signal, 1)
	signal.Notify(sigKillChan, os.Interrupt, os.Kill)

	cfg, err := config.FromFile(configFilePath)
	if err != nil {
		panic(err)
	}
	if shouldAutogenKeys == true {
		err = cfg.GenerateKeys(keysDirPath, passphrase)
		if err != nil {
			panic(err)
		}
		os.Exit(0)
	}

	accountKeys, err := cfg.AccountsMap(constants.EndToEndKeyType, keysDirPath, passphrase)
	if err != nil {
		panic(err)
	}
	userPKI, err := user_pki.UserPKIFromJsonFile(userPKIFile)
	if err != nil {
		panic(err)
	}
	mixPKI, err := mix_pki.StaticPKIFromFile(mixPKIFile)
	if err != nil {
		panic(err)
	}
	pinnedProviders, err := cfg.GetProviderPinnedKeys()
	if err != nil {
		panic(err)
	}
	peerAuthenticator := auth.ProviderAuthenticator(pinnedProviders)
	providerSessionPool, err := session_pool.New(accountKeys, cfg, peerAuthenticator, mixPKI)
	if err != nil {
		panic(err)
	}
	routeFactory := path_selection.New(mixPKI, constants.HopsPerPath, constants.PoissonLambda)
	store, err := storage.New(dbFile)
	if err != nil {
		panic(err)
	}
	// ensure each account has a boltdb bucket
	identities := cfg.AccountIdentities()
	store.CreateAccountBuckets(identities)
	fetchers := make(map[string]*proxy.Fetcher)
	senders := make(map[string]*proxy.Sender)
	for _, identity := range identities {
		privateKey, err := accountKeys.GetIdentityKey(identity)
		if err != nil {
			panic(err)
		}
		handler := block.NewHandler(privateKey, rand.Reader)
		sender, err := proxy.NewSender(identity, providerSessionPool, store, routeFactory, userPKI, handler)
		if err != nil {
			panic(err)
		}
		senders[identity] = sender
	}
	sendScheduler := proxy.NewSendScheduler(senders)
	for _, identity := range identities {
		privateKey, err := accountKeys.GetIdentityKey(identity)
		if err != nil {
			panic(err)
		}
		handler := block.NewHandler(privateKey, rand.Reader)
		fetcher := proxy.NewFetcher(identity, providerSessionPool, store, sendScheduler, handler)
		fetchers[identity] = fetcher
	}
	smtpProxy := proxy.NewSmtpProxy(accountKeys, rand.Reader, userPKI, store, providerSessionPool, routeFactory, sendScheduler)
	periodicRetriever := proxy.NewFetchScheduler(fetchers, time.Second*7)
	periodicRetriever.Start()

	// create pop3 service
	pop3Service := proxy.NewPop3Service(store)

	var smtpServer, pop3Server *server.Server
	if len(cfg.SMTPProxy.Network) == 0 {
		smtpServer = server.New(constants.DefaultSMTPNetwork, constants.DefaultSMTPAddress, smtpProxy.HandleSMTPSubmission, nil)
	} else {
		smtpServer = server.New(cfg.SMTPProxy.Network, cfg.SMTPProxy.Address, smtpProxy.HandleSMTPSubmission, nil)
	}

	if len(cfg.POP3Proxy.Network) == 0 {
		pop3Server = server.New(constants.DefaultPOP3Network, constants.DefaultPOP3Address, pop3Service.HandleConnection, nil)
	} else {
		pop3Server = server.New(cfg.POP3Proxy.Network, cfg.POP3Proxy.Address, pop3Service.HandleConnection, nil)
	}

	log.Notice("mixclient startup")

	err = smtpServer.Start()
	if err != nil {
		panic(err)
	}
	defer smtpServer.Stop()
	err = pop3Server.Start()
	if err != nil {
		panic(err)
	}
	defer pop3Server.Stop()

	for {
		select {
		case <-sigKillChan:
			log.Notice("mixclient shutdown")
			return
		}
	}
}
