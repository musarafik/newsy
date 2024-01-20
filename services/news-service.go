package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/musarafik/newsy/structs"
	"github.com/musarafik/newsy/utilities"
)

func GetHeadlines() ([]structs.Article, error) {
	url := fmt.Sprintf("https://newsapi.org/v2/top-headlines?country=us&apiKey=%s", utilities.GetNewsApiKey())
	var headlinesResponse structs.HeadlinesResponse

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return nil, errors.New("error making GET request")
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&headlinesResponse)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, errors.New("error decoding JSON")
	}

	return headlinesResponse.Articles, nil
}
