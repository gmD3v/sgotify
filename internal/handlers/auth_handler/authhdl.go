package authhdl

import (
	"log"
	"net/http"

	"github.com/spf13/cobra"
	ports "gmd3v.com/sgotify/internal/core/ports/services"
)

type AuthHandler struct {
	authService ports.AuthService
}

func NewHandler(authService ports.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(cmd *cobra.Command, args []string) {
	ch := make(chan string)
	print("Starting server")
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.FormValue("code")
		ch <- code
		w.Write([]byte("You can now close this window"))
	})
	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()
	log.Println("\nWaiting for authorization... \n", "Please log in to Spotify by visiting the following page in your browser: ", h.authService.RedirectUrl())

	code := <-ch
	err := h.authService.Login(code)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Logged in successfully")
}

func (h *AuthHandler) Refresh(accessToken, refresh string) error {
	return h.authService.Refresh(accessToken, refresh)
}

func (h *AuthHandler) Logout() error {
	return h.authService.Logout()
}
