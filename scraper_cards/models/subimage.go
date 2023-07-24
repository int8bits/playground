package models

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"strings"

	"github.com/disintegration/imaging"
)

var name = []int{80, 75, 990, 170}
var typeMonster = []int{310, 975, 760, 890}
var effects = []int{80, 1270, 990, 980}
var description = []int{80, 1270, 990, 1170}
var attack = []int{415, 1385, 530, 1275}
var rest = []int{540, 1385, 655, 1275}
var nameBio = []int{90, 320, 170, 1150}
var typeBio = []int{710, 650, 755, 850}
var effectBio = []int{755, 100, 920, 1400}
var descriptionBio = []int{900, 100, 1000, 1400}

type SubImage struct {
	Image       string `json:"image"`
	NameMonster string
}

func (subImg *SubImage) Split() {
	mapSplits := map[string][]int{
		"name":            name,
		"type_monster":    typeMonster,
		"effects":         effects,
		"description":     description,
		"attack":          attack,
		"rest":            rest,
		"name_bio":        nameBio,
		"type_bio":        typeBio,
		"effect_bio":      effectBio,
		"description_bio": descriptionBio,
	}

	if err := os.Mkdir("process_image/"+subImg.NameMonster, os.ModePerm); err != nil {
		log.Println(err)
	}

	for k, v := range mapSplits {
		// Cargar la imagen desde un archivo
		fmt.Println(k)
		file, err := os.Open(subImg.Image + ".png")
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
		}).SubImage(image.Rect(v[0], v[1], v[2], v[3])) // La subimagen se ubica desde el punto (50, 50) y tiene un tamaño de 150x150

		if strings.Contains(k, "bio") {
			fmt.Println("rotate image")
			subImage = imaging.Rotate90(subImage)
		}

		// Guardar la subimagen en un archivo
		out, err := os.Create("process_image/" + subImg.NameMonster + "/" + k + "_subimagen.png")
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
}
