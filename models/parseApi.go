package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	url            = "https://groupietrackers.herokuapp.com/api/"
	epointArtist   = "artists"
	epointRelation = "relation"
)

func init() {
	ParseJsonAPI(url+epointArtist, &All_Artists)
	ParseJsonAPI(url+epointRelation, &All_Relations)
	fmt.Println("urls are parsed")
}

func ParseJsonAPI(url string, n interface{}) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Parse error")
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Body read error")
		return
	}
	err = json.Unmarshal(body, &n)
	if err != nil {
		log.Println("Unmarshal error")
		return
	}
}
