package reactionpostapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
)

type reactionPostApi struct {
	baseUrl string
}

func NewReactionPostApi(baseUrl string) *reactionPostApi {
	return &reactionPostApi{
		baseUrl: baseUrl,
	}
}

func (api *reactionPostApi) GetListFollwingByUserId(ctx context.Context, userId uuid.UUID) ([]uuid.UUID, error) {
	type Payload struct {
		Id uuid.UUID `json:"id"`
	}

	client := &http.Client{}
	data := Payload{Id: userId}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", api.baseUrl, "/follow/get-following"), bytes.NewBuffer(jsonData))
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
