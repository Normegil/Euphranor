package main

import (
	"flag"
	"fmt"
	"math/rand"
	"path/filepath"
	"time"
)

const safe = "/home/normegil/Images/.wallpapers/s"
const unsafe = "/home/normegil/Images/.wallpapers/us"

var minutes = flag.Int("t", 5, "Interval in minutes between 2 wallpaper change")

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	flag.Parse()
	var cmdPaths []string
	for _, path := range flag.Args() {
		pathToGive, err := filepath.Abs(path)
		if nil != err {
			panic(err)
		}
		cmdPaths = append(cmdPaths, pathToGive)
	}

	for {
		if len(cmdPaths) != 0 {
			changeWallpaper(cmdPaths...)
		} else {
			if triggerSafeWall() {
				changeWallpaper(safe)
			} else {
				changeWallpaper(safe, unsafe)
			}
		}

		time.Sleep(time.Duration(*minutes) * time.Minute)
	}
}

func triggerSafeWall() bool {
	now := time.Now()

	return time.Saturday == now.Weekday() || time.Sunday == now.Weekday() ||
		(now.Hour() >= 17 && now.Hour() < 8)
}

func changeWallpaper(paths ...string) {
	path, err := getRandomImagePath(paths...)
	if nil != err {
		panic(err)
	}
	fmt.Println("Change wallpaper by: " + path)
	wc := WallpaperChange{path}
	err = wc.Change()
	if err != nil {
		panic(err)
	}
}
