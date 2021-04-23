package main

import (
	"TestBataraGuru/helper"
	"TestBataraGuru/utils"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

var JSONLink string
var RequestLimit int = 0
var OffsetCount int = 0

type Pokemon struct {
	NextLink string    `json:"next"`
	PrevLink *string   `json:"previous"`
	Results  []Results `json:"results"`
}

type Results struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

var collection *mongo.Collection

func main() {
	collection = helper.ConnectDB()
	router := mux.NewRouter()
	router.HandleFunc("/api/pokemon", SavePokemon)

	fmt.Println("Server listen on : 8000")
	log.Fatal(http.ListenAndServe(":8000", router))

	//SavePokemonToDatabase()
}

func SavePokemon(wrt http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		utils.WrapAPIError(wrt, req, "Method not allowed!", http.StatusMethodNotAllowed)
		return
	}

	authToken := req.Header.Get("Authorization")
	if authToken == "" {
		utils.WrapAPIError(wrt, req, "Invalid auth", http.StatusForbidden)
		return
	}

	if authToken != "izin-masuk-gan" {
		utils.WrapAPIError(wrt, req, "Invalid auth", http.StatusForbidden)
		return
	}

	if RequestLimit == 5 {
		log.Printf("Request %s %s", req.Method, http.StatusText(429))
		time.Sleep(1 * time.Minute)
		RequestLimit = 0
	}
	RequestLimit = RequestLimit + 1

	// Link berisi daftar pokemon dari API pokeapi.co berupa JSON
	JSONLink = fmt.Sprintf("https://pokeapi.co/api/v2/pokemon?offset=%v&limit=20", OffsetCount*10)
	OffsetCount = OffsetCount + 1

	// Membuat objek client
	pokemonClient := http.Client{
		Timeout: time.Second * 2, // Timeout setelah 2 detik
	}

	// Membuat objek request
	req, err := http.NewRequest(http.MethodGet, JSONLink, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Melakukan request dan mengambil responnya
	res, getErr := pokemonClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	// Menutup body jika sudah selesai dibaca
	if res.Body != nil {
		defer res.Body.Close()
	}

	// Membaca Body respon
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	// Parsing body ke dalam bantuk JSON dan assigh ke struct Pokemon
	pokemon := Pokemon{}
	jsonErr := json.Unmarshal(body, &pokemon)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	createPokemon(wrt, req, pokemon.Results)
	// Menampilkan daftar pokemon ke client dalam bentuk JSON
	//utils.WrapAPIData(wrt, req, pokemon, http.StatusOK, "Success!")

	/*result, err := json.Marshal(pokemon))
	fmt.Println(string(result))*/
}

func createPokemon(wrt http.ResponseWriter, req *http.Request, data []Results) {
	// Simpan ke dalam database
	dataPokemon := make([]interface{}, len(data))
	for i, v := range data {
		dataPokemon[i] = v
	}

	_, err := collection.InsertMany(context.TODO(), dataPokemon)
	if err != nil {
		log.Fatal(err)
		helper.GetError(err, wrt)
		return
	}
	utils.WrapAPISuccess(wrt, req, "Success store data!", http.StatusOK)
	//log.Printf("Berhasil menyimpan ke database!")
}
