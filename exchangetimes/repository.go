package exchangetimes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	EXCHANGE_TIMES_ENDPOINT = "exchangetimes"
)

type ExchangeTime struct {
	Name     string
	Open     string
	Close    string
	Holidays []string
}

type ExchangeTimes struct {
	Results []ExchangeTime
}

type ExchangeTimesRepository interface {
	GetExchangeTimes(context.Context) (*ExchangeTimes, error)
}

type HttpExchangeTimesRepository struct {
	client  *http.Client
	baseURL string
}

func NewExchangeTimesRepository(baseURL string) ExchangeTimesRepository {
	client := &http.Client{
		Timeout: time.Duration(time.Second * 1),
	}
	return &HttpExchangeTimesRepository{client: client, baseURL: baseURL}
}

func (repo *HttpExchangeTimesRepository) GetExchangeTimes(ctx context.Context) (*ExchangeTimes, error) {
	url := fmt.Sprintf("%s/%s", repo.baseURL, EXCHANGE_TIMES_ENDPOINT)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Accept", "application/json")

	response, err := repo.client.Do(request)
	if err != nil {
		fmt.Printf("error occurred here = %s", err.Error())
		return nil, err
	}
	fmt.Println("response = ", response)

	if response.StatusCode != http.StatusOK {
		return nil, err
	}

	var v *ExchangeTimes

	if err = json.NewDecoder(response.Body).Decode(&v); err != nil {
		return nil, err
	}

	fmt.Printf("exchangeTimes = %+v\n", v)

	return v, nil
}
