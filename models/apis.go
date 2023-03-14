package models

import (
	"strconv"
	"strings"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	RelationsURL string   `json:"relations"`
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Relations struct {
	Index []Relation
}

var (
	All_Relations Relations
	All_Artists   []Artist
)

func (a *Artist) GetDatesLocations() *Relation {
	if a.RelationsURL == "" {
		return &Relation{}
	}
	relationIDstr := getLastID(a.RelationsURL, "/")
	relationID, err := strconv.Atoi(relationIDstr)
	if err != nil {
		return &Relation{}
	}
	return &All_Relations.Index[relationID-1]
}

func (a *Artist) GetLocationsOnly() []string {
	locations := []string{}
	for location := range All_Relations.Index[a.ID-1].DatesLocations {
		locations = append(locations, location)
	}
	return locations
}

func getLastID(url, sep string) string {
	slice := strings.Split(url, sep)
	return slice[len(slice)-1]
}
