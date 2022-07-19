package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

type Inventory struct {
	Namening string
	Material string
	Count    uint
}

func main() {
	sweater := Inventory{Material: "wool", Count: 17}
	garderob := []Inventory{
		sweater,
		{
			Namening: "Hate",
			Material: "coton",
			Count:    13,
		},
	}
	for _, v := range garderob {
		tmpl, err := template.New("test").Parse("{{.Namening}} => {{.Count }} items are made of {{.Material}}")
		if err != nil {
			log.Println(err)
		}
		err = tmpl.Execute(os.Stdout, v)
		if err != nil {
			log.Println(err)
		}
		fmt.Println()
	}
	//tmpl, err := template.New("test").Parse("{{.Count }} items are made of {{.Material}}")
	//if err != nil {
	//	log.Println(err)
	//}
	//err = tmpl.Execute(os.Stdout, sweater)
	//if err != nil {
	//	log.Println(err)
	//}
}
