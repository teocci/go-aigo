// Package core
// Created by RTT.
// Author: teocci@yandex.com on 2025-1월-15
package core

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/teocci/aigo/src/config"
	"github.com/teocci/aigo/src/webserver"
)

func Start() error {
	pid := os.Getpid()
	fmt.Println("PID:", pid)

	cfg := config.Get()

	go webserver.Start(cfg)

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		log.Println(sig)
		done <- true
	}()
	log.Println("Server start awaiting signal")
	<-done
	log.Println("Server stop working by signal")

	return nil
}
