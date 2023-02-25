package main

import (
	"fmt"
	"time"

	"github.com/neovim/go-client/nvim"
)


func main() {

	// create neovim process
	neovim, _ := nvim.NewChildProcess(
		nvim.ChildProcessArgs(
			[]string{
				"--embed",
				"-u",
				"NONE",
				"--cmd",
				"set clipboard=unnamed",
				"./testfile",
			}...,
		),
		nvim.ChildProcessServe(false),
	)

	updateCh := make(chan [][]interface{}, 10000)
	neovim.RegisterHandler("redraw", func(updates ...[]interface{}) {
		updateCh <- updates
	})

	// Serve neovim
	exitCh := make(chan error, 1)
	go func() {
		exitCh <-neovim.Serve()
	}()

	// Attach neovim
	go func() {
		neovim.AttachUI(100, 20, make(map[string]interface{}))
	}()

	// Send some command
	go func() {
		time.Sleep(500 * time.Millisecond)
		neovim.Input("otest<Esc>otest<Esc>")

		// this command reproduces the problem
		neovim.Input("dd<Esc>:wq<Enter>")

		// // this command does not reproduce the problem
		// neovim.Input(":wq<Enter>")
	}()

	<- exitCh
	fmt.Println("program exit!")
}
