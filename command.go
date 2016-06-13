package main

import (
	"fmt"
	"os"
	"os/exec"
)

type WallpaperChange struct {
	path string
}

func (wc WallpaperChange) Change() error {
	const command = "gsettings"

	// gnomeCMD := []string{"gsettings", "set",		"org.gnome.desktop.background",		"picture-uri", "file://"+wc.path}
	awesomeCMD := []string{"feh", "--bg-scale", wc.path}

	fmt.Println("Cmd: feh --bg-scale "+ wc.path)
	cmd := exec.Command(awesomeCMD[0], awesomeCMD[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if nil != err {
		return err
	}
	return nil
}
