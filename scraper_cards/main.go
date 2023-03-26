package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/anaskhan96/soup"
)

type Monster struct {
	Name   string `json:"name"`
	Type   string `json:"type_monster"`
	ImgB64 string `json:"img_base64"`
}

func main() {
	resp, err := soup.Get("https://kodem-tcg.com/raices-misticas")

	if err != nil {
		log.Printf("Error to get data %s \n", err)
		os.Exit(1)
	}

	doc := soup.HTMLParse(resp)
	images := doc.FindAll("img")
	fmt.Println(len(images))

	for _, image := range images {
		fmt.Println("alt:", image.Attrs()["alt"])
		name := strings.Replace(image.Attrs()["alt"], ",", "", -1)
		name = strings.Replace(name, " ", "_", -1)

		m := Monster{
			Name:   name,
			ImgB64: image.Attrs()["src"],
		}
		err := m.SaveImg()

		if err != nil {
			fmt.Println(err)
		}
	}
}

func (m Monster) SaveImg() error {
	data := m.ImgB64

	index := strings.Index(data, ";base64,")

	if index < 0 {
		log.Printf("The %s is not possible create image", m.Name)
		err := errors.New("error with string to covert")

		return err
	}

	// imageType := data[11:index]
	// fmt.Println(imageType)

	unbased, err := base64.StdEncoding.DecodeString(data[index+8:])
	if err != nil {
		log.Println("Cannot decode b64")
	}
	pathSave := "images/" + m.Name + ".png"
	f, err := os.Create(pathSave)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	if _, err := f.Write(unbased); err != nil {
		log.Println(err)
	}
	if err := f.Sync(); err != nil {
		log.Println(err)
	}

	return nil
}
