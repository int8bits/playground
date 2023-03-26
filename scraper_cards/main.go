package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

type Monster struct {
	Name   string `json:"name"`
	Type   string `json:"type_monster"`
	ImgB64 string `json:"img_base64"`
}

func main() {
	// create folder to storage images
	if err := os.Mkdir("images", os.ModePerm); err != nil {
		log.Println(err)
	}

	// raices-misticas
	c := colly.NewCollector(
		colly.AllowedDomains("kodem-tcg.com"),
	)
	cardsMonster := make([]Monster, 0)

	c.OnHTML("img", func(element *colly.HTMLElement) {
		name := strings.Replace(element.Attr("alt"), ",", "", -1)
		name = strings.Replace(name, " ", "_", -1)
		monster := Monster{
			Name:   name,
			ImgB64: element.Attr("src"),
		}

		cardsMonster = append(cardsMonster, monster)
	})

	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})

	c.Visit("https://kodem-tcg.com/raices-misticas")

	for _, monster := range cardsMonster {
		err := monster.SaveImg()

		if err != nil {
			log.Println(err)
		}

		// if i == 0 {
		// 	break
		// }
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

	imageType := data[11:index]
	fmt.Println(imageType)

	unbased, err := base64.StdEncoding.DecodeString(data[index+8:])
	if err != nil {
		panic("Cannot decode b64")
	}
	pathSave := "images/" + m.Name + ".png"
	f, err := os.Create(pathSave)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err := f.Write(unbased); err != nil {
		panic(err)
	}
	if err := f.Sync(); err != nil {
		panic(err)
	}

	return nil
}