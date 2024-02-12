package controlhandler

import (
	"fmt"

	"github.com/spf13/cobra"
	ports "gmd3v.com/sgotify/internal/core/ports/services"
)

type ControlHandler struct {
	ctrlService ports.ControlService
}

func NewHandler(service ports.ControlService) *ControlHandler {
	return &ControlHandler{ctrlService: service}
}

func (h *ControlHandler) ControllerMain(cmd *cobra.Command, args []string) {
	var err error
	var controllers = map[string]func(*cobra.Command, []string){
		"play":     h.PlaySong,
		"pause":    h.PauseSong,
		"next":     h.NextSong,
		"previous": h.PreviousSong,
		"get":      h.GetSong,
	}
	if len(args) == 0 {
		cmd.Help()
		return
	}
	if controller, ok := controllers[args[0]]; ok {
		controller(cmd, args)
	} else {
		err = fmt.Errorf("Invalid command %s", args[0])
	}
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (h *ControlHandler) PlaySong(cmd *cobra.Command, args []string) {
	var err error
	if len(args) >= 2 {
		err = h.ctrlService.PlaySong(args[2])
	} else {
		err = h.ctrlService.PlaySong("")
	}
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Playing song")
}

func (h *ControlHandler) PauseSong(cmd *cobra.Command, args []string) {
	err := h.ctrlService.PauseSong()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Pausing song")
}

func (h *ControlHandler) NextSong(cmd *cobra.Command, args []string) {
	err := h.ctrlService.NextSong()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Next song")
}

func (h *ControlHandler) PreviousSong(cmd *cobra.Command, args []string) {
	err := h.ctrlService.PreviousSong()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Previous song")
}

func (h *ControlHandler) GetSong(cmd *cobra.Command, args []string) {
	song, err := h.ctrlService.GetSong()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Song: %s\nArtist: %s\nAlbum: %s\nRelease Date: %s\nLength: %d\nIs Playing: %t\n", song.Name, song.Artist, song.Album, song.ReleaseDate, song.Length, song.IsPlaying)
}
