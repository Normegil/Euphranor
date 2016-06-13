package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"path/filepath"
)

var supportedExtentions = map[string]struct{}{
	".jpg":  struct{}{},
	".jpeg": struct{}{},
	".png":  struct{}{},
}

func getRandomImagePath(pathToDirs ...string) (string, error) {
	allPics := make([]string, 1)
	fmt.Println(pathToDirs)
	for _, pathToDir := range pathToDirs {
		fmt.Println("Path:" + pathToDir)
		pics, err := getListOfPictures(pathToDir)
		if nil != err {
			return "", err
		}
		allPics = append(allPics, pics...)
	}

	return choosePicture(allPics)
}

func getListOfPictures(pathToDir string) ([]string, error) {
	var pictures []string
	files, err := ioutil.ReadDir(pathToDir)
	if nil != err {
		return nil, err
	}

	for _, file := range files {
		pathToFile := pathToDir + "/" + file.Name()
		if file.IsDir() {
			pics, err := getListOfPictures(pathToFile)
			if nil != err {
				return nil, err
			}
			pictures = append(pictures, pics...)
		} else if checkExtention(pathToFile) {
			pictures = append(pictures, pathToFile)
		}
	}
	return pictures, nil
}

func checkExtention(path string) bool {
	extention := filepath.Ext(path)
	_, ok := supportedExtentions[extention]
	return ok
}

func choosePicture(pics []string) (string, error) {
	if len(pics) <= 0 {
		return "", errors.New("No pics found")
	}
	return pics[rand.Intn(len(pics))], nil
}
