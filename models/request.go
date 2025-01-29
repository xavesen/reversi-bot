package models

type CreateGameRequest struct {
	PlayerId 	string 	`json:"player_id"`
}

type ErrorContents struct {
	Code 		int		`json:"code"`
	Message 	string	`json:"message"`
}

type CreateGameResult struct {
	GameId	string	`json:"game_id"`
	Color 	string	`json:"color"`
}

type CreateGameResponse struct {
	Status	string				`json:"status"`
	Error	ErrorContents		`json:"error"`
	Result 	CreateGameResult	`json:"result"`
}


