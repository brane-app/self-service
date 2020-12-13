package main

import (
	"git.gastrodon.io/imonke/monkebase"
	"git.gastrodon.io/imonke/monketype"

	"context"
	"net/http"
	"os"
	"testing"
)

const (
	nick  = "self"
	email = "self@imonke.io"
)

var (
	user monketype.User
)

func owned(owner string) (it context.Context) {
	it = context.WithValue(
		context.TODO(),
		"requester",
		owner,
	)

	return
}

func userOK(test *testing.T, fetched map[string]interface{}, target monketype.User) {
	if fetched["id"].(string) != target.ID {
		test.Errorf("id mismatch! have: %s, want: %s", fetched["id"], target.ID)
	}

	if fetched["nick"].(string) != target.Nick {
		test.Errorf("nick mismatch! have: %s, want: %s", fetched["nick"], target.Nick)
	}

	if fetched["email"] == nil {
		test.Errorf("did not get private email!")
	}

	if fetched["created"] == nil {
		test.Errorf("did not get private created!")
	}
}

func TestMain(main *testing.M) {
	monkebase.Connect(os.Getenv("MONKEBASE_CONNECTION"))
	user = monketype.NewUser(nick, "", email)

	var err error
	if err = monkebase.WriteUser(user.Map()); err != nil {
		panic(err)
	}

	var result int = main.Run()
	monkebase.DeleteUser(user.ID)
	os.Exit(result)
}

func Test_getSelf(test *testing.T) {
	var request *http.Request
	var err error
	if request, err = http.NewRequestWithContext(owned(user.ID), "GET", "imonke.io/me", nil); err != nil {
		test.Fatal(err)
	}

	var code int
	var r_map map[string]interface{}
	if code, r_map, err = getSelf(request); err != nil {
		test.Fatal(err)
	}

	if code != 200 {
		test.Errorf("got code %d", code)
	}

	userOK(test, r_map["user"].(map[string]interface{}), user)
}
