package main

import (
	"errors"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
)

type Tamaño struct {
	Anchura int
	Altura  int
}

func ObtenerTamaño(url string) (Tamaño, error) {
	resp, err := http.Get(url)
	if err != nil {
		return Tamaño{}, errors.New("Error al obtener la imagen")
	}
	defer resp.Body.Close()

	m, _, err := image.Decode(resp.Body)
	if err != nil {
		return Tamaño{}, errors.New("Error al obtener la imagen")
	}
	datos := m.Bounds()

	var dimensiones Tamaño
	dimensiones.Altura = datos.Dy()
	dimensiones.Anchura = datos.Dx()

	return dimensiones, nil
}
