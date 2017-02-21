package gokit

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"generator/app/domain"
	"generator/app/inject"

	transport "github.com/go-kit/kit/transport/http"
)

func AddRequest() transport.RequestFunc {
	return func(ctx context.Context, req *http.Request) context.Context {
		return context.WithValue(ctx, inject.Request, req)
	}
}

func emptyDecoder() transport.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		closeRequest(r)

		return nil, nil
	}
}

func simpleEncoder() transport.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
		w.Header().Add("Content-type", "application/json")
		return json.NewEncoder(w).Encode(response)
	}
}

func campaignCollectionDecoder() transport.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		var request domain.CampaignCollection

		defer closeRequest(r)

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			return request, fmt.Errorf("Incorrect input. Error: %v", err)
		}

		if !request.Verify() {
			return request, fmt.Errorf("Incorrect input: %v", request)
		}

		return request, nil
	}
}

func userDecoder() transport.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		var request domain.User

		defer closeRequest(r)

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			return request, fmt.Errorf("Incorrect input. Error: %v", err)
		}

		if !request.Verify() {
			return request, fmt.Errorf("Incorrect input: %v", request)
		}

		return request, nil
	}
}

func closeRequest(r *http.Request) {
	if r != nil && r.Body != nil {
		_, err := io.Copy(ioutil.Discard, r.Body)
		r.Body.Close()

		if err != nil {
			log.Print("Closing request error: %v", err)
		}
	}
}
