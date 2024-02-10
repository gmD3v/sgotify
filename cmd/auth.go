/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"gmd3v.com/sgotify/internal/core/services/auth"
	authhdl "gmd3v.com/sgotify/internal/handlers/auth_handler"
	"gmd3v.com/sgotify/internal/repositories/authrepo"
	"gmd3v.com/sgotify/internal/repositories/tokenrepo"
)

func init() {
	authRepo := authrepo.NewWRPAuth()
	tokenRepo := tokenrepo.NewTokenRepository()
	authSrv := auth.NewService(authRepo, tokenRepo)
	authHandler := authhdl.NewHandler(authSrv)
	rootCmd.AddCommand(&cobra.Command{
		Use:   "auth",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: authHandler.Login,
	})
}
