package responses

type TokenStatisResponse struct {
	TokenStatis string `json:"token_statis"`
}

type TokenStatisNotFoundResponse struct {
	Message string `json:"message"`
}

type TokenJWTResponse struct {
	Token string `json:"token"`
}
