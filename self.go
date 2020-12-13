package main

import (
	"git.gastrodon.io/imonke/monkebase"
	"git.gastrodon.io/imonke/monketype"

	"net/http"
)

func getSelf(request *http.Request) (code int, r_map map[string]interface{}, err error) {
	var requester string = request.Context().Value("requester").(string)

	var who monketype.User
	if who, _, err = monkebase.ReadSingleUser(requester); err != nil {
		return
	}

	code = 200
	r_map = map[string]interface{}{
		"user": who.Map(),
	}
	return
}
