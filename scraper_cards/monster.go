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
)

type Monster struct {
	Name             string   `json:"name"`
	Type             string   `json:"type_monster"`
	ImgB64           string   `json:"img_base64"`
	pathFileComplete string   `json:"path_file_complete"`
	pathFile         string   `json:"path_file"`
	subImages        []string `json:"sub_images"`
}

func (mons *Monster) SetVariablesFile() {
	mons.pathFile = "images/" + mons.Name
	mons.pathFileComplete = mons.pathFile + mons.Name
}

func (mons *Monster) SaveImg() error {
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
