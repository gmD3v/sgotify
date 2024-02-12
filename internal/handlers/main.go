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

var _controller CobraFunc

func NewHandler(tokenRepo repositories.TokenRepository, authRepo repositories.AuthRepository, controller CobraFunc) *Handler {
	_controller = controller
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
	_controller(cmd, args)
	if r := recover(); r != nil {
		fmt.Print("Debe estar autenticado para realizar esta acción. Por favor, autentíquese usando el comando auth.")
	}
}
