package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Error struct {
	Message string `json:"message"`
}

type ViaCep struct {
	Url         string `json:"url"`
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type FindCep struct {
	Url         string `json:"url"`
	Uf          string `json:"uf"`
	Cidade      string `json:"cidade"`
	Bairro      string `json:"bairro"`
	Logradouro  string `json:"logradouro"`
	Cep         string `json:"cep"`
	Complemento string `json:"complemento"`
	Nome        string `json:"nome"`
	Status      string `json:"status"`
	Tipo        string `json:"tipo"`
	CodigoIbge  string `json:"codigo_ibge"`
}

const urlViaCep = "https://viacep.com.br/ws/%s/json"
const urlFindCep = "https://example.api.findcep.com/v1/cep/%s.json"

func main() {
	http.HandleFunc("/", BuscarCepHandler)
	http.ListenAndServe(":8000", nil)
}

func BuscarCepHandler(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	if cep == "" {
		e := Error{Message: "cep is invalid"}
		json.NewEncoder(w).Encode(e)
		return
	}
	viaCep := make(chan interface{})
	findCep := make(chan interface{})
	go BuscarViaCep(cep, viaCep)
	go BuscarFindCep(cep, findCep)
	select {
	case v := <-viaCep:
		json.NewEncoder(w).Encode(v)
	case v := <-findCep:
		json.NewEncoder(w).Encode(v)
	case <-time.After(1 * time.Second):
		json.NewEncoder(w).Encode(Error{Message: "exceeded timeout of 1 second"})
	}
}

func BuscarViaCep(cep string, ch chan<- interface{}) {
	url := fmt.Sprintf(urlViaCep, cep)
	res, err := http.Get(url)
	if err != nil {
		er := Error{Message: err.Error()}
		ch <- er
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		er := Error{Message: err.Error()}
		ch <- er
		return
	}
	var viaCep ViaCep
	if err := json.Unmarshal(body, &viaCep); err != nil {
		er := Error{Message: err.Error()}
		ch <- er
		return
	}
	viaCep.Url = url
	ch <- viaCep
}

func BuscarFindCep(cep string, ch chan<- interface{}) {
	url := fmt.Sprintf(urlFindCep, cep)
	res, err := http.Get(url)
	if err != nil {
		er := Error{Message: err.Error()}
		ch <- er
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		er := Error{Message: err.Error()}
		ch <- er
		return
	}
	var findCep FindCep
	if err := json.Unmarshal(body, &findCep); err != nil {
		er := Error{Message: err.Error()}
		ch <- er
		return
	}
	findCep.Url = url
	ch <- findCep
}
