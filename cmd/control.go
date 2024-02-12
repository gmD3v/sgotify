/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"gmd3v.com/sgotify/internal/core/services/control"
	"gmd3v.com/sgotify/internal/handlers"
	controlhandler "gmd3v.com/sgotify/internal/handlers/control_handler"
	"gmd3v.com/sgotify/internal/repositories/authrepo"
	"gmd3v.com/sgotify/internal/repositories/controlrepo"
	"gmd3v.com/sgotify/internal/repositories/tokenrepo"
)

func init() {
	tokenRepo := tokenrepo.NewTokenRepository()
	controlRepo := controlrepo.NewWrpControl()
	authRepo := authrepo.NewWRPAuth()
	controlSrv := control.NewService(controlRepo, tokenRepo)
	handler := controlhandler.NewHandler(controlSrv)
	// handler :=
	handlerController := handlers.NewHandler(tokenRepo, authRepo, handler.ControllerMain)
	var controlCmd = &cobra.Command{
		Use:   "control",
		Short: "Control your music player from the command line.",
		Long: `Control your music player from the command line. For example:
		- sgotify control --device <device_id> play

		Opciones:
		- play: Reproduce la canción actual.
		- pause: Pausa la canción actual.
		- next: Reproduce la siguiente canción.
		- previous: Reproduce la canción anterior.
		- get: Muestra la canción actual.`,
		Run: handlerController.ControllerProtected,
	}
	controlCmd.PersistentFlags().StringP("device", "d", "", "Device ID")
	controlCmd.PersistentFlags().StringP("song", "s", "", "Song ID")
	controlCmd.PersistentFlags().StringP("playlist", "p", "", "Playlist ID")
	rootCmd.AddCommand(controlCmd)

}
