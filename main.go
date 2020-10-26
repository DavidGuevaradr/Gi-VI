package main

import (
	"fmt"

	"./multimedia"
)

func main() {
	var (		
		titulo  string
		formato string
		canales   string
		titulo_a  string
		formato_a string
		duracion  string
		titulo_v  string
		formato_v string
		frames    string
	)

	fmt.Println("----IMAGEN----\n\n")

	fmt.Println("Titulo: ")
	fmt.Scan(&titulo)
	fmt.Println("Formato: ")
	fmt.Scan(&formato)
	fmt.Println("Canales: ")
	fmt.Scan(&canales)

	fmt.Println("\n\n----AUDIO----\n\n")

	fmt.Println("Titulo: ")
	fmt.Scan(&titulo_a)
	fmt.Println("Formato: ")
	fmt.Scan(&formato_a)
	fmt.Println("Duracion: ")
	fmt.Scan(&duracion)

	fmt.Println("\n\n----VIDEO----\n\n")

	fmt.Println("Titulo: ")
	fmt.Scan(&titulo_v)
	fmt.Println("Formato: ")
	fmt.Scan(&formato_v)
	fmt.Println("Frames: ")
	fmt.Scan(&frames)

	fmt.Println("\n\n")

	fmt.Println("MOSTRANDO ACONTINUACION LOS REGISTRADOS...")

	inst := multimedia.ContenidoWeb{Multi: []multimedia.Multimedia{

		&multimedia.Imagen{titulo, formato, canales},
		&multimedia.Audio{titulo_a, formato_a, duracion},
		&multimedia.Video{titulo_v, formato_v, frames},
	}}

	fmt.Println(inst.Mostrar())


}
