package main

import (
	"github.com/brane-app/database-library"
	"github.com/brane-app/types-library"

	"net/http"
)

func getSelf(request *http.Request) (code int, r_map map[string]interface{}, err error) {
	var requester string = request.Context().Value("requester").(string)

	var who types.User
	if who, _, err = database.ReadSingleUser(requester); err != nil {
		return
	}

	code = 200
	r_map = map[string]interface{}{
		"user": who.Map(),
	}
	return
}
