package handlers

import (
	"crud-repo-2/f"
	"crud-repo-2/responses"
	"net/http"
)

type tokenStatisHandler struct {
	setToken responses.TokenStatisResponse
}

func NewTokenStatisHandler() *tokenStatisHandler {
	return &tokenStatisHandler{}
}

func (t *tokenStatisHandler) GetTokenValue(w http.ResponseWriter, r *http.Request) {
	var headerToken = r.Header.Get("token-statis")

	if headerToken == "" {
		message := responses.TokenStatisNotFoundResponse{
			Message: "Token Not Found",
		}
		f.WriteToJsonError(w, r, message)
	}

	tokenStatis := responses.TokenStatisResponse{
		TokenStatis: headerToken,
	}

	t.setToken = tokenStatis

	f.WriteToJson(w, r, t.setToken)
}
