package postapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
)

type postAPI struct {
	baseURL string
}

func NewFollowAPI(baseURL string) *postAPI {
	return &postAPI{
		baseURL: baseURL,
	}
}

func (api *postAPI) GetListFollowingByUserId(ctx context.Context, uid uuid.UUID) ([]uuid.UUID, error) {
	type Payload struct {
		Id uuid.UUID `json:"id"`
	}
	client := &http.Client{}
	data := Payload{Id: uid}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", api.baseURL, "/follow/get-following"), bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	type Response struct {
		Data []uuid.UUID `json:"data"`
	}

	response := Response{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return response.Data, nil
}
