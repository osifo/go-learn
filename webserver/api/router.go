package api

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"osifo.dev/go-learn/web/data"
)

func ApiRouter(response http.ResponseWriter, request *http.Request) {
	reg := regexp.MustCompile(`/api/exhibitions/(?P<id>\d+)?`)
	matches := reg.FindStringSubmatch(request.URL.Path)

	response.Header().Set("Content-Type", "application/json")

	if request.Method == http.MethodGet {
		exhibitionIndex := reg.SubexpIndex("id")

		if exhibitionIndex < 0 {
			json.NewEncoder(response).Encode(data.GetAllExhibitions())
		} else {
			index, _ := strconv.Atoi(matches[exhibitionIndex])
			json.NewEncoder(response).Encode(data.GetExhibitionByIndex(index))
		}
	}

	if request.Method == http.MethodPost {
		var exhibitionParam data.Exhibition
		decodeError := json.NewDecoder(request.Body).Decode(&exhibitionParam)

		if decodeError != nil {
			http.Error(response, "request body must match the shape of an exhibition", http.StatusNotAcceptable)
		}

		data.AddNewExhibition(exhibitionParam)
		response.WriteHeader(http.StatusCreated)
		response.Write([]byte("Created"))
		// json.NewEncoder(response).Encode(exhibitionParam)

	}
}
