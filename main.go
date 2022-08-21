package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Content struct {
	ID              string `json:"id"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	Characteristics *Info  `json:"characteristics"`
	Price           string `json:"price"`
	Category        string `json:"category"`
	Images          *Image `json:"images"`
	Rating          *Rate  `json:"rating"`
}

type Info struct {
	DisplayType        string `json:"display_type"`
	DisplayDiagonal    string `json:"display_diagonal"`
	DisplayResolution  string `json:"display_resolution"`
	BuildInMemory      string `json:"build_in_memory"`
	RAM                string `json:"ram"`
	CameraType         string `json:"camera_type"`
	CameraMegaPixels   string `json:"camera_mega_pixels"`
	Battery            string `json:"battery"`
	CPU                string `json:"cpu"`
	SIMCard            string `json:"sim_card"`
	OS                 string `json:"os"`
	WirelessInterfaces string `json:"wireless_interfaces"`
	Internet           string `json:"internet"`
	Weight             string `json:"weight"`
}

type Image struct {
	ImageUrl string `json:"image_url"`
}

type Rate struct {
	NumByOneToTen string `json:"num_by_one_to_ten"`
	NumByStoToK   string `json:"num_by_sto_to_k"`
}

var ctx []Content

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ctx)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, item := range ctx {

		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(ctx)
}

func main() {
	r := mux.NewRouter()

	for idx := 0; idx <= 1000; idx++ {

		i := strconv.Itoa(rand.Intn(1000000))

		ctx = append(ctx, Content{
			ID:          strconv.Itoa(idx),
			Title:       "AnyPhone " + strconv.Itoa(idx),
			Description: "It's description by - " + strconv.Itoa(idx) + " phone",
			Characteristics: &Info{
				DisplayType:        "Super Retina XDR",
				DisplayDiagonal:    strconv.Itoa(rand.Intn(15)),
				DisplayResolution:  strconv.Itoa(rand.Intn(3000)) + "*" + strconv.Itoa(rand.Intn(2000)),
				BuildInMemory:      "256",
				RAM:                "4",
				CameraType:         "triple",
				CameraMegaPixels:   strconv.Itoa(idx) + "+" + strconv.Itoa(idx) + "+" + strconv.Itoa(idx),
				Battery:            strconv.Itoa(rand.Intn(10000)),
				CPU:                "Any C" + strconv.Itoa(idx) + " Bionic",
				SIMCard:            "nano-SIM",
				OS:                 "Any OS",
				WirelessInterfaces: "Bluetooth, Wi-Fi, NFC, 4G LTE",
				Internet:           "3G, 4G LTE",
				Weight:             strconv.Itoa(idx),
			},
			Price:    i,
			Category: "smartphone",
			Images: &Image{
				ImageUrl: "https://my-apple-store.ru/wa-data/public/shop/products/48/46/4648/images/8143/8143.750.jpg",
			},
			Rating: &Rate{
				NumByOneToTen: strconv.Itoa(rand.Intn(10)),
				NumByStoToK:   strconv.Itoa(rand.Intn(1000)),
			},
		})
	}

	r.HandleFunc("/", getProducts).Methods("GET")
	r.HandleFunc("/{id}", getProduct).Methods("GET")

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
