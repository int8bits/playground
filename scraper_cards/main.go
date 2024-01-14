package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"ws-kodem/models"
	"ws-kodem/utilities"

	"github.com/anaskhan96/soup"
)

func main() {
	fromLocal := true
	fileName := "kodem-page.html"
	var resp string
	var subImages []models.SubImage
	var subImagesLinks []models.SubImage

	if !fromLocal {
		resp, err := soup.Get("https://kodem-tcg.com/raices-misticas")

		if err != nil {
			log.Printf("Error to get data %s \n", err)
			os.Exit(1)
		}

		f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			log.Println(err.Error())
		}
		defer f.Close()

		f.WriteString(resp + "\n")
	} else {
		fileR, err := os.ReadFile(fileName)

		if err != nil {
			log.Printf("Is not possible read file %s \n", fileName)
			os.Exit(1)
		}

		resp = string(fileR)
	}

	doc := soup.HTMLParse(resp)
	images := doc.FindAll("img")
	fmt.Println(len(images))

	for _, image := range images {
		fmt.Println("alt:", image.Attrs()["alt"])
		name := strings.Replace(image.Attrs()["alt"], ",", "", -1)
		name = strings.Replace(name, " ", "_", -1)
		name = utilities.RemoveAccents(name)

		m := models.Monster{
			Name:   name,
			ImgB64: image.Attrs()["src"],
		}
		path, err := m.SaveImg()

		if err != nil {
			fmt.Println(err)
			continue
		}

		subImage := new(models.SubImage)
		subImage.Image = path
		subImage.NameMonster = name
		subImages = append(subImages, *subImage)
	}

	// fmt.Println(subImages)
	for _, subImage := range subImages {
		subImagesLinks = append(subImagesLinks, models.Split(subImage))
	}

	// fmt.Println(subImagesLinks[0])

	for _, sil := range subImagesLinks {
		models.GetData(sil)
	}
}
