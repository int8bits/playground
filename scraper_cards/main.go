package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/anaskhan96/soup"
)

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
