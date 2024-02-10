package controlhandler

import (
	"github.com/spf13/cobra"
	ports "gmd3v.com/sgotify/internal/core/ports/services"
)

type ControlHandler struct {
	ctrlService ports.ControlService
}

func NewHandler(service ports.ControlService) *ControlHandler {
	return &ControlHandler{ctrlService: service}
}

func (h *ControlHandler) PlaySong(cmd *cobra.Command, args []string) {
	var err error
	if len(args) >= 2 {
		err = h.ctrlService.PlaySong(args[2])
	} else {
		err = h.ctrlService.PlaySong("")
	}
	if err != nil {
		panic(err)
	}
}

func (h *ControlHandler) PauseSong(cmd *cobra.Command, args []string) {
	err := h.ctrlService.PauseSong()
	if err != nil {
		panic(err)
	}
}

func (h *ControlHandler) NextSong(cmd *cobra.Command, args []string) {
	err := h.ctrlService.NextSong()
	if err != nil {
		panic(err)
	}
}

func (h *ControlHandler) PreviousSong(cmd *cobra.Command, args []string) {
	err := h.ctrlService.PreviousSong()
	if err != nil {
		panic(err)
	}
}

func (h *ControlHandler) GetSong(cmd *cobra.Command, args []string) {
	_, err := h.ctrlService.GetSong()
	if err != nil {
		panic(err)
	}
}
