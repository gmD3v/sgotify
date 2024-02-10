/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gmd3v.com/sgotify/internal/core/services/control"
	controlhandler "gmd3v.com/sgotify/internal/handlers/control_handler"
	"gmd3v.com/sgotify/internal/repositories/controlrepo"
	"gmd3v.com/sgotify/internal/repositories/tokenrepo"
)

func init() {
	tokenRepo := tokenrepo.NewTokenRepository()
	controlRepo := controlrepo.NewWrpControl()
	controlSrv := control.NewService(controlRepo, tokenRepo)
	handler := controlhandler.NewHandler(controlSrv)
	rootCmd.AddCommand(&cobra.Command{
		Use:   "control",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			// get command type from command and call the appropriate function
			if len(args) == 0 {
				fmt.Println("Please provide a command")
				return
			}
			command := args[0]
			if command == "play" {
				handler.PlaySong(cmd, args)
			}
			if command == "pause" {
				handler.PauseSong(cmd, args)
			}
			if command == "next" {
				handler.NextSong(cmd, args)
			}
			if command == "previous" {
				handler.PreviousSong(cmd, args)
			}

		},
	})

}
