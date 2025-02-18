package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/xavesen/reversi-bot/config"
	"github.com/xavesen/reversi-bot/models"
)

type ErrStatusCodeNot200 struct {}

func (err *ErrStatusCodeNot200) Error() string {
	return "response code != 200"
}

func makeRequest(url string, requestStruct any, retries int) ([]byte, error) {
	for i := 0; i <= retries; i++ {
		data, err := json.Marshal(requestStruct)
		if err != nil {
			fmt.Printf("Error marshalling request data url=%s: %s\n", url, err)
			if i == retries-1 {
				return nil, err
			} else {
				continue
			}
		}

		request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
		if err != nil {
			fmt.Printf("Error creating request url=%s: %s\n", url, err)
			if i == retries-1 {
				return nil, err
			} else {
				continue
			}
		}

		request.Header.Set("Content-Type", "application/json")
		response, err := http.DefaultClient.Do(request)
		if err != nil {
			fmt.Printf("Error making API request url=%s: %s\n", url, err)
			if i == retries-1 {
				return nil, err
			} else {
				continue
			}
		}

		respBody, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("Error reading API response body url=%s: %s\n", url, err)
			if i == retries-1 {
				return nil, err
			} else {
				continue
			}
		}

		if response.StatusCode != 200 {
			var responseStruct models.BasicResponse
			err = json.Unmarshal(respBody, &responseStruct)
			if err != nil {
				fmt.Printf("Error parsing API response body to basic response struct url=%s: %s\n", url, err)
				if i == retries-1 {
					return nil, err
				} else {
					continue
				}
			}

			fmt.Printf("API responded with an error url=%s; code=%d; message=%s\n", url, responseStruct.Error.Code, responseStruct.Error.Message)
			if i == retries-1 {
				return nil, &ErrStatusCodeNot200{}
			} else {
				continue
			}
		}
		
		return respBody, nil
	}
	return nil, errors.New("unsupported retries value")
}

func CreateGame(conf *config.Config) (*models.CreateGameResult, error) {
	requestStruct := &models.CreateGameRequest{PlayerId: conf.PlayerUUID}
	respBody, err := makeRequest(conf.ApiURL+conf.CreateGameURL, requestStruct, conf.ApiRetries)
	if err != nil {
		return nil, err
	}

	var responseStruct models.CreateGameResponse
	err = json.Unmarshal(respBody, &responseStruct)
	if err != nil {
		fmt.Printf("Error parsing API response body while creating new game: %s\n", err)
		return nil, err
	}

	return &responseStruct.Result, nil
}

func GetGameList(conf *config.Config) ([]models.GameInfo, error) {
	requestStruct := &models.GameListRequest{PlayerId: conf.PlayerUUID}
	respBody, err := makeRequest(conf.ApiURL+conf.GameListUrl, requestStruct, conf.ApiRetries)
	if err != nil {
		return nil, err
	}

	var responseStruct models.GameListResponse
	err = json.Unmarshal(respBody, &responseStruct)
	if err != nil {
		fmt.Printf("Error parsing API response body while getting game list: %s\n", err)
		return nil, err
	}

	return responseStruct.Result, nil
}

func JoinGame(gameId string, conf *config.Config) (*models.JoinGameResult, error) {
	requestStruct := &models.JoinGameRequest{PlayerId: conf.PlayerUUID, GameId: gameId}
	respBody, err := makeRequest(conf.ApiURL+conf.JoinGameURL, requestStruct, conf.ApiRetries)
	if err != nil {
		return nil, err
	}

	var responseStruct models.JoinGameResponse
	err = json.Unmarshal(respBody, &responseStruct)
	if err != nil {
		fmt.Printf("Error parsing API response body while joining game: %s\n", err)
		return nil, err
	}

	return &responseStruct.Result, nil
}

func GetGameStatus(gameId string, conf *config.Config) (*models.GameStatusResult, error) {
	requestStruct := &models.GameStatusRequest{PlayerId: conf.PlayerUUID, GameId: gameId}
	respBody, err := makeRequest(conf.ApiURL+conf.GameStatusURL, requestStruct, conf.ApiRetries)
	if err != nil {
		return nil, err
	}

	var responseStruct models.GameStatusResponse
	err = json.Unmarshal(respBody, &responseStruct)
	if err != nil {
		fmt.Printf("Error parsing API response body while joining game: %s\n", err)
		return nil, err
	}

	return &responseStruct.Result, nil
}

func MakeMove(gameId string, move string, conf *config.Config) (*models.MoveResult, error) {
	requestStruct := &models.MoveRequest{PlayerId: conf.PlayerUUID, GameId: gameId, Move: move}
	respBody, err := makeRequest(conf.ApiURL+conf.MakeMoveURL, requestStruct, conf.ApiRetries)
	if err != nil {
		return nil, err
	}

	var responseStruct models.MoveResponse
	err = json.Unmarshal(respBody, &responseStruct)
	if err != nil {
		fmt.Printf("Error parsing API response body while joining game: %s\n", err)
		return nil, err
	}

	return &responseStruct.Result, nil
}