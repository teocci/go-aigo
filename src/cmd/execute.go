// Package cmd
// Created by RTT.
// Author: teocci@yandex.com on 2025-1월-14
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/teocci/aigo/src/core"

	"github.com/teocci/aigo/src/config"
	"github.com/teocci/aigo/src/webserver"
)

var (
	app = &cobra.Command{
		Use:           "aigo",
		Short:         "Aigo is an AI Search Engine implemented in Go",
		Long:          `Aigo is a modular AI Search Engine built with Fiber.`,
		RunE:          runE,
		SilenceErrors: false,
		SilenceUsage:  false,
	}

	errs chan error
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the Fiber webserver",
	Run: func(cmd *cobra.Command, args []string) {
		// Load configuration
		err := config.Load()
		if err != nil {
			log.Fatalf("Error loading configuration: %v", err)
		}

		cfg := config.Get()

		// Start the webserver
		webserver.Start(cfg)
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	config.AddFlags(app)
}

func initConfig() {
	if err := config.Load(); err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	fmt.Printf("Using profile: %v\n", config.Get())
}

func runE(ccmd *cobra.Command, args []string) error {
	fmt.Printf("Arguments: %v\n", args)
	fmt.Printf("\n")
	fmt.Printf("Press Ctrl+C to stop the server\n")
	fmt.Printf("\n")

	// Load command-line options
	configFile, _ := ccmd.Flags().GetString("config")

	// Load configuration file if specified
	if configFile != "" {
		if err := config.Load(); err != nil {
			log.Fatalf("Failed to load config file: %v", err)
		}
	}

	errs = make(chan error)

	// Start the application in a goroutine
	go runApp()

	// Wait for an error to occur
	if err := <-errs; err != nil {
		log.Fatalf("Application error: %v", err)
	}

	return nil
}

func runApp() {
	if err := core.Start(); err != nil {
		errs <- err
		return
	}
	errs <- nil
}

func Execute() {
	if err := app.Execute(); err != nil {
		log.Fatalf("Execution failed: %v", err)
	}
}
