package handlers

import (
	"fmt"
	"g-tracker/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type BandData struct {
	Artist    *models.Artist
	Relations *models.Relation
}

var (
	all_data []BandData
	cache    = make(map[string][]BandData)
)

func init() {
	Get_result()
}

func Get_result() {
	limit := len(models.All_Artists)
	for i := 0; i < limit; i++ {
		all_data = append(all_data,
			BandData{
				Artist:    &models.All_Artists[i],
				Relations: models.All_Artists[i].GetDatesLocations(),
			})
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	var data []BandData
	if r.Method != "GET" {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		ErrorHandler(w, http.StatusNotFound)
		return
	}

	query := r.URL.Query().Get("search")
	if len(query) > 0 {
		data = Search(query)
	} else {
		data = all_data
	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Println("Template parsing error")
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}

func Search(query string) []BandData {
	if results, ok := cache[query]; ok {
		return results
	}

	// Perform the search
	results := SearchHandler(query)

	// Cache the results
	cache[query] = results
	return results
}

func SearchHandler(query string) []BandData {
	found := make(map[int]bool)
	result := []BandData{}
	for _, v := range all_data {
		if strings.HasPrefix(strings.ToLower(v.Artist.Name), strings.ToLower(query)) {
			found[v.Artist.ID] = true
			result = append(result, v)
		} else if containsWord(v.Artist.Name, query) {
			found[v.Artist.ID] = true
			result = append(result, v)
		}
	}

	for _, v := range all_data {
		cDate := strconv.Itoa(v.Artist.CreationDate)
		for _, member := range v.Artist.Members {
			if strings.HasPrefix(member, query) {
				if _, ok := found[v.Artist.ID]; !ok {
					found[v.Artist.ID] = true
					result = append(result, v)
				}
			}
		}
		for _, location := range v.Artist.GetLocationsOnly() {
			if containsWord(location, query) {
				if _, ok := found[v.Artist.ID]; !ok {
					found[v.Artist.ID] = true
					result = append(result, v)
				}
			}
		}
		if containsWord(cDate, query) || containsWord(v.Artist.FirstAlbum, query) {
			if _, ok := found[v.Artist.ID]; !ok {
				found[v.Artist.ID] = true
				result = append(result, v)
			}
		}
	}
	fmt.Println(query + " cached!")
	return result
}

func containsWord(str string, query string) bool {
	seps := strings.NewReplacer("-", " ", "|", " ")
	str = seps.Replace(str)
	query = seps.Replace(query)
	str1 := strings.Split(str, " ")
	str2 := strings.Split(query, " ")
	for _, word2 := range str2 {
		for _, word1 := range str1 {
			if strings.HasPrefix(strings.ToLower(word1), strings.ToLower(word2)) {
				return true
			}
		}
	}
	return false
}
