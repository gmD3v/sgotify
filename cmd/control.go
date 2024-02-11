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
	handlerController := handlers.NewHandler(tokenRepo, authRepo, map[string]handlers.CobraFunc{
		"play":     handler.PlaySong,
		"pause":    handler.PauseSong,
		"next":     handler.NextSong,
		"previous": handler.PreviousSong,
	})
	rootCmd.AddCommand(&cobra.Command{
		Use:   "control",
		Short: "Control your music player from the command line.",
		Long: `Control your music player from the command line. For example:
		- sgotify control play
		- sgotify control pause
		- sgotify control next
		- sgotify control previous

	Remember to authenticate first using the auth command.`,
		Run: handlerController.ControllerProtected,
	})

}
