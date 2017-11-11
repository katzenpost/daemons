// main.go - Katzenpost server binary.
// Copyright (C) 2017  David Stainton.
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

package main

import (
	"context"
	"encoding/base64"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/jonboulle/clockwork"
	"github.com/katzenpost/authority/config"
	"github.com/katzenpost/authority/server"
	"github.com/katzenpost/core/crypto/eddsa"
	"github.com/katzenpost/core/crypto/rand"
)

func main() {
	cfgFile := flag.String("f", "katzenpost_authority.toml", "Path to the config file.")
	newKey := flag.Bool("n", false, "Generate a new key and print it's base64 encoding")
	flag.Parse()

	if *newKey {
		key, err := eddsa.NewKeypair(rand.Reader)
		if err != nil {
			panic(err)
		}
		fmt.Printf("eddsa private key is %s\n", base64.StdEncoding.EncodeToString(key.Bytes()))
		os.Exit(0)
	}

	// Set the umask to something "paranoid".
	syscall.Umask(0077)

	cfg, err := config.LoadFile(*cfgFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load config file '%v': %v\n", *cfgFile, err)
		os.Exit(-1)
	}

	clock := clockwork.NewRealClock()
	ctx := context.TODO() // XXX
	// Start up the server.
	svr, err := server.New(cfg, ctx, clock)
	fmt.Println("after server new")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to spawn server instance: %v\n", err)
		os.Exit(-1)
	}
	defer svr.Shutdown()

	sigKillChan := make(chan os.Signal, 1)
	signal.Notify(sigKillChan, os.Interrupt, os.Kill)

	for {
		select {
		case <-sigKillChan:
			return
		}
	}

}
