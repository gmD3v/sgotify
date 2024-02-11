package handlers

import (
	"fmt"

	"github.com/spf13/cobra"
	"gmd3v.com/sgotify/internal/core/ports/repositories"
)

type Handler struct {
	tokenRepo repositories.TokenRepository
	authRepo  repositories.AuthRepository
}
type CobraFunc func(cmd *cobra.Command, args []string)

var controllers = map[string]CobraFunc{}

func NewHandler(tokenRepo repositories.TokenRepository, authRepo repositories.AuthRepository, controller map[string]CobraFunc) *Handler {
	controllers = controller
	return &Handler{
		tokenRepo: tokenRepo,
		authRepo:  authRepo,
	}
}

func (h *Handler) ControllerProtected(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.Help()
		return
	}
	token, refresh, err := h.tokenRepo.GetSavedToken()
	if err != nil {
		panic(err)
	}
	if !h.authRepo.IsTokenValid(token) {
		fmt.Println("El token ha expirado. Refrescando...")
		token, refresh, err = h.authRepo.RefreshToken(refresh)
		if err != nil {
			panic(err)
		}
		err := h.tokenRepo.SaveToken(token, refresh)
		if err != nil {
			panic(err)
		}
	}
	if f, ok := controllers[args[0]]; ok {
		f(cmd, args[1:])
	} else {
		cmd.Help()
	}

	if r := recover(); r != nil {
		fmt.Print("Debe estar autenticado para realizar esta acción. Por favor, autentíquese usando el comando auth.")
	}
}
