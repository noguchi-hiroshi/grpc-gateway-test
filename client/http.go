package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/noguchi-hiroshi/grpc-gateway-test/user"
	"io/ioutil"
	"net/http"
	"strconv"
)

const UserEndpoint = "http://localhost:5000"

type httpUserClient struct {
}

type userBody struct {
	ID string
	Email string
	Password string
}

func (h *httpUserClient) Create(email string, password string) (int64, error) {
	values, err := json.Marshal(userBody{Email: email, Password: password})

	resp, err := http.Post(UserEndpoint + "/users", "application/json", bytes.NewBuffer(values))
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return 0, err
	}

	u := userBody{}

	if err := json.Unmarshal(body, &u); err != nil {
		return 0, err
	}

	return h.parseInt64(u.ID)
}

func (h *httpUserClient) Find(id int64) (*user.Entity, error) {
	url := fmt.Sprintf("%s/users/%d", UserEndpoint, id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	u := userBody{}
	if err := json.Unmarshal(body, &u); err != nil {
		return nil, err
	}

	id, err = h.parseInt64(u.ID)
	if err != nil {
		return nil, err
	}

	return &user.Entity{
		ID: id,
		Email: u.Email,
		Password: u.Password,
	}, err
}

func (h *httpUserClient) parseInt64(str string) (int64, error) {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return int64(i), nil
}

func NewHttpUserClient() UserClient {
	return &httpUserClient{}
}