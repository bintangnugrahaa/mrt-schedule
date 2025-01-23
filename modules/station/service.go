package station

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/bintangnugrahaa/mrt-schedule/common/client"
)

type Service interface {
	GetAllStation() ([]StationResponse, error)
}

type service struct {
	client *http.Client
}

func NewService() Service {
	return &service{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (s *service) GetAllStation() ([]StationResponse, error) {
	url := "https://www.jakartamrt.co.id/id/val/stasiuns"

	byteResponse, err := client.DoRequest(s.client, url)
	if err != nil {
		return nil, err
	}

	var stations []Station
	err = json.Unmarshal(byteResponse, &stations)
	if err != nil {
		return nil, errors.New("gagal mem-parsing response stasiun")
	}

	var response []StationResponse
	for _, item := range stations {
		response = append(response, StationResponse{
			Id:   item.Id,
			Name: item.Name,
		})
	}

	return response, nil
}
