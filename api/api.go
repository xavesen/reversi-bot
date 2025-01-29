package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/xavesen/reversi-bot/models"
)

func CreateGame(url string, playerId string) (*models.CreateGameResponse, error) {
	requestStruct := &models.CreateGameRequest{PlayerId: playerId}
	data, err := json.Marshal(requestStruct)
	if err != nil {
		fmt.Printf("Error marshalling request for game creation: %s\n", err)
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		fmt.Printf("Error creating request for game creation: %s\n", err)
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Printf("Error making API request for game creation: %s\n", err)
		return nil, err
	}

	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading API response body while creating new game: %s\n", err)
		return nil, err
	}

	var responseStruct models.CreateGameResponse
	err = json.Unmarshal(respBody, &responseStruct)
	if err != nil {
		fmt.Printf("Error parsing API response body while creating new game: %s\n", err)
		return nil, err
	}

	return &responseStruct, nil
}