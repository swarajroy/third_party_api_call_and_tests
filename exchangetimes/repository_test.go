package exchangetimes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetExchangeTimes(t *testing.T) {

	exchangeTimes := &ExchangeTimes{
		Results: []ExchangeTime{
			{
				Name:     "XNAS",
				Open:     "openTs",
				Close:    "closeTs",
				Holidays: []string{"1", "2"},
			},
			{
				Name:     "XNAS",
				Open:     "openTs",
				Close:    "closeTs",
				Holidays: []string{"1", "2"},
			},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		var enpoint = fmt.Sprintf("/%s", EXCHANGE_TIMES_ENDPOINT)

		if r.URL.Path != enpoint {
			fmt.Print(r.URL.Path)
			t.Errorf("Expected to request '%s', got: %s", enpoint, r.URL.Path)
		}
		if r.Header.Get("Accept") != "application/json" {
			t.Errorf("Expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
		}

		w.WriteHeader(http.StatusOK)
		v, err := json.Marshal(exchangeTimes)
		if err != nil {
			t.Log("got an error while Marshalling")
			t.Errorf("error =%+v\n", err)
			t.Fail()
		}
		w.Write(v)
	}))
	defer server.Close()

	repo := NewLogMiddleware(NewExchangeTimesRepository(server.URL))

	results, err := repo.GetExchangeTimes(context.TODO())
	if err != nil {
		t.Errorf("error =%+v\n", err)
		t.Failed()
	}

	assert.Equal(t, exchangeTimes, results)
}
