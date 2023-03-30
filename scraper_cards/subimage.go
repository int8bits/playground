package main

import (
	"image"
	"image/png"
	"log"
	"os"
)

type SubImage struct {
	Image        string `json:"image"`
	nameSubImage string `json:"name_sub_image"`
	x1           int    `json:"x1"`
	x2           int    `json:"x2"`
	y1           int    `json:"y1"`
	y2           int    `json:"y2"`
}

type SubImageName struct {
	Image        string `json:"image"`
	nameSubImage string `json:"name_sub_image"`
	x1           int    `json:"x1"`
	x2           int    `json:"x2"`
	y1           int    `json:"y1"`
	y2           int    `json:"y2"`
}

type SubImageDescription struct {
	Image        string `json:"image"`
	nameSubImage string `json:"name_sub_image"`
	x1           int    `json:"x1"`
	x2           int    `json:"x2"`
	y1           int    `json:"y1"`
	y2           int    `json:"y2"`
}

type SubImageTypes struct {
	Image        string `json:"image"`
	nameSubImage string `json:"name_sub_image"`
	x1           int    `json:"x1"`
	x2           int    `json:"x2"`
	y1           int    `json:"y1"`
	y2           int    `json:"y2"`
}

type SubImageValues struct {
	Image        string `json:"image"`
	nameSubImage string `json:"name_sub_image"`
	x1           int    `json:"x1"`
	x2           int    `json:"x2"`
	y1           int    `json:"y1"`
	y2           int    `json:"y2"`
}

type SubImageBioType struct {
	Image        string `json:"image"`
	nameSubImage string `json:"name_sub_image"`
	x1           int    `json:"x1"`
	x2           int    `json:"x2"`
	y1           int    `json:"y1"`
	y2           int    `json:"y2"`
}

type SubImageBioDescription struct {
	Image        string `json:"image"`
	nameSubImage string `json:"name_sub_image"`
	x1           int    `json:"x1"`
	x2           int    `json:"x2"`
	y1           int    `json:"y1"`
	y2           int    `json:"y2"`
}

func (subImg *SubImage) Split() {
	// Cargar la imagen desde un archivo
	file, err := os.Open("imagen.png")
	if err != nil {
		log.Fatal("Error abriendo archivo:", err)
	}
	defer file.Close()

	// Decodificar la imagen en formato PNG
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal("Error decodificando la imagen:", err)
	}

	// Obtener una subimagen de la imagen original
	subImage := img.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(image.Rect(50, 50, 200, 200)) // La subimagen se ubica desde el punto (50, 50) y tiene un tamaño de 150x150

	// Guardar la subimagen en un archivo
	out, err := os.Create("subimagen.png")
	if err != nil {
		log.Fatal("Error creando archivo:", err)
	}
	defer out.Close()

	// Codificar la subimagen en formato PNG y escribir en el archivo
	err = png.Encode(out, subImage)
	if err != nil {
		log.Fatal("Error escribiendo la subimagen:", err)
	}

	log.Println("La subimagen se ha guardado correctamente.")
}
