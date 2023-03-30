package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
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

		// if i == 2 {
		// 	break
		// }
	}
}

func (mons Monster) SaveImg() error {
	data_full := mons.ImgB64

	index := strings.Index(data_full, ";base64,")

	if index < 0 {
		log.Printf("The %s is not possible create image", mons.Name)
		err := errors.New("error with string to covert")

		return err
	}
	pathSave := "images/" + mons.Name

	// Decodificar la cadena base64
	// fmt.Println(data_full[index+8:])
	imageData, err := base64.StdEncoding.DecodeString(data_full[index+8:])
	if err != nil {
		log.Fatal("Error decodificando la cadena base64:", err)
	}

	// Decodificar la imagen en formato PNG
	img, _, err := image.Decode(strings.NewReader(string(imageData)))
	if err != nil {
		fmt.Println("Error decodificando la imagen en png:", err)

		// Decodificar la imagen en formato JPEG
		img, err := jpeg.Decode(strings.NewReader(string(imageData)))
		if err != nil {
			log.Fatal("Error decodificando la imagen:", err)
		}

		// Guardar la imagen en un archivo
		out, err := os.Create(pathSave + ".jpg")
		if err != nil {
			log.Fatal("Error creando archivo:", err)
		}
		defer out.Close()

		// Codificar la imagen en formato JPEG y escribir en el archivo
		err = jpeg.Encode(out, img, nil)
		if err != nil {
			log.Fatal("Error escribiendo la imagen:", err)
		}

		return nil
	}

	// Guardar la imagen en un archivo
	out, err := os.Create(pathSave + ".png")
	if err != nil {
		log.Fatal("Error creando archivo:", err)
	}
	defer out.Close()

	// Codificar la imagen en formato PNG y escribir en el archivo
	err = png.Encode(out, img)
	if err != nil {
		log.Fatal("Error escribiendo la imagen:", err)
	}

	log.Println("La imagen se ha guardado correctamente.")

	return nil
}
