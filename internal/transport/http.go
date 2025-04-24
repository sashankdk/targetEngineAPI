package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"targetApi/internal/delivery"

	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHTTPHandler(svc delivery.DeliveryService) http.Handler {
	getDeliveryHandler := httptransport.NewServer(
		delivery.GetdeliveryEndpoint(svc),
		decodeGetDeliveryRequest,
		encodeResponse,
	)

	mux := http.NewServeMux()
	mux.Handle("/v1/delivery", getDeliveryHandler)
	return mux
}

func decodeGetDeliveryRequest(_ context.Context, r *http.Request) (interface{}, error) {
	appID := r.URL.Query().Get("app")
	country := r.URL.Query().Get("country")
	os := r.URL.Query().Get("os")
	return delivery.GetdeliveryRequest{
		AppId:   appID,
		Country: country,
		Os:      os,
	}, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(delivery.GetdeliveryResponse)

	if res.Err != "" {
		w.WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(w).Encode(map[string]string{"error": res.Err})
	}

	if len(res.Campaigns) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}

	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(res.Campaigns)
}
