package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/xavesen/reversi-bot/models"
)

type ErrStatusCodeNot200 struct {}

func (err *ErrStatusCodeNot200) Error() string {
	return "response code != 200"
}

func makeRequest(url string, requestStruct any) ([]byte, error) {
	data, err := json.Marshal(requestStruct)
	if err != nil {
		fmt.Printf("Error marshalling request data url=%s: %s\n", url, err)
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		fmt.Printf("Error creating request url=%s: %s\n", url, err)
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Printf("Error making API request url=%s: %s\n", url, err)
		return nil, err
	}

	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading API response body url=%s: %s\n", url, err)
		return nil, err
	}

	if response.StatusCode != 200 {
		var responseStruct models.BasicResponse
		err = json.Unmarshal(respBody, &responseStruct)
		if err != nil {
			fmt.Printf("Error parsing API response body to basic response struct url=%s: %s\n", url, err)
			return nil, err
		}

		fmt.Printf("API responded with an error url=%s; code=%d; message=%s\n", url, responseStruct.Error.Code, responseStruct.Error.Message)
		return nil, &ErrStatusCodeNot200{}
	}
	
	return respBody, nil
}

func CreateGame(url string, playerId string) (*models.CreateGameResponse, error) {
	requestStruct := &models.CreateGameRequest{PlayerId: playerId}
	respBody, err := makeRequest(url, requestStruct)
	if err != nil {
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

func GetGameList(url string, playerId string) (*models.GameListResponse, error) {
	requestStruct := &models.GameListRequest{PlayerId: playerId}
	respBody, err := makeRequest(url, requestStruct)
	if err != nil {
		return nil, err
	}

	var responseStruct models.GameListResponse
	err = json.Unmarshal(respBody, &responseStruct)
	if err != nil {
		fmt.Printf("Error parsing API response body while getting game list: %s\n", err)
		return nil, err
	}

	return &responseStruct, nil
}

func JoinGame(url string, playerId string, gameId string) (*models.JoinGameResponse, error) {
	requestStruct := &models.JoinGameRequest{PlayerId: playerId, GameId: gameId}
	respBody, err := makeRequest(url, requestStruct)
	if err != nil {
		return nil, err
	}

	var responseStruct models.JoinGameResponse
	err = json.Unmarshal(respBody, &responseStruct)
	if err != nil {
		fmt.Printf("Error parsing API response body while joining game: %s\n", err)
		return nil, err
	}

	return &responseStruct, nil
}

func GetGameStatus(url string, playerId string, gameId string) (*models.GameStatusResponse, error) {
	requestStruct := &models.GameStatusRequest{PlayerId: playerId, GameId: gameId}
	respBody, err := makeRequest(url, requestStruct)
	if err != nil {
		return nil, err
	}

	var responseStruct models.GameStatusResponse
	err = json.Unmarshal(respBody, &responseStruct)
	if err != nil {
		fmt.Printf("Error parsing API response body while joining game: %s\n", err)
		return nil, err
	}

	return &responseStruct, nil
}