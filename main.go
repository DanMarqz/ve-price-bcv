package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
  "encoding/json"

	"github.com/gocolly/colly/v2"
)

type BCVDatata struct {
	USD  string `json:"usd_price"`
	EUR string `json:"eur_price"`
	VDate  string `json:"value_date"`
}

func main() {
	// Crear un cliente HTTP personalizado con verificación TLS deshabilitada
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// Crear un nuevo colector con el cliente HTTP personalizado
	c := colly.NewCollector()
	c.WithTransport(httpClient.Transport)

	var priceUSD, priceEUR, fechaValor string

	// Definir la lógica para manejar los elementos extraídos
	c.OnHTML("#dolar", func(e *colly.HTMLElement) {
		priceUSD = e.DOM.Find("strong").Text()
	})

	c.OnHTML("#euro", func(e *colly.HTMLElement) {
		priceEUR = e.DOM.Find("strong").Text()
	})

	c.OnHTML(".dinpro", func(e *colly.HTMLElement) {
		fechaValor = e.DOM.Find("span").Text()
	})

	// Visitar la URL deseada
	err := c.Visit("https://bcv.org.ve")
	if err != nil {
		log.Fatal(err)
	}

  // Crear una instancia de la estructura BCVDatata con datos
	data := BCVDatata{
		USD:    priceUSD,
		EUR:    priceEUR,
		VDate:  fechaValor,
	}

	// Convertir la estructura a JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error al convertir a JSON:", err)
		return
	}

	// Imprimir el JSON resultante
	fmt.Println(string(jsonData)) 
}
