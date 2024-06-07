package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMakeGet(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.String() {
		case "/ping":
			response := map[string]string{"message": "pong"}
			v, err := json.Marshal(response)
			if err != nil {
				t.Fatal("failed to marshal response!")
			}
			w.Write(v)
		default:
			http.NotFound(w, r)
		}

	}))
	defer mockServer.Close()
	ApiEndpoint = mockServer.URL
	defer func() {
		ApiEndpoint = "actualEndpoint"
	}()

	tests := []struct {
		name             string
		target           string
		expectedResponse string
	}{
		{
			name:             "hit the endpoint",
			target:           "/ping",
			expectedResponse: "pong",
		},
		{
			name:             "miss the endpoint",
			target:           "/test",
			expectedResponse: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := MakeGet(tt.target)
			r := require.New(t)
			r.NoError(err)
			r.Equal(tt.expectedResponse, result)
		})
	}
}
