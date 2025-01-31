package models

type ErrorContents struct {
	Code 		int		`json:"code"`
	Message 	string	`json:"message"`
}

type BasicResponse struct {
	Status	string			`json:"status"`
	Error	ErrorContents	`json:"error"`
	Result 	any				`json:"result"`
}

type CreateGameRequest struct {
	PlayerId 	string 	`json:"player_id"`
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

type GameListRequest struct {
	PlayerId 	string 	`json:"player_id"`
}

type GameInfo struct {
	GameId		string	`json:"game_id"`
	FirstPlayer	string	`json:"first_player"`
}

type GameListResponse struct {
	Status	string			`json:"status"`
	Error	ErrorContents	`json:"error"`
	Result 	[]GameInfo		`json:"result"`
}

type JoinGameRequest struct {
	PlayerId 	string	`json:"player_id"`
	GameId		string	`json:"game_id"`
}

type JoinGameResult struct {
	Result 	bool	`json:"result"`
	Color 	string	`json:"color"`
}

type JoinGameResponse struct {
	Status	string				`json:"status"`
	Error	ErrorContents		`json:"error"`
	Result 	JoinGameResult		`json:"result"`
}

type GameStatusRequest struct {
	PlayerId 	string	`json:"player_id"`
	GameId		string	`json:"game_id"`
}

type GameStatusResult struct {
	Status		string	`json:"status"`
	LastMove	string	`json:"last_move"`
}

type GameStatusResponse struct {
	Status	string				`json:"status"`
	Error	ErrorContents		`json:"error"`
	Result 	GameStatusResult	`json:"result"`
}