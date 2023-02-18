package context

import (
	"context"
	"fmt"
	"net/http"
)

// takes a store
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		data, err := store.Fetch(ctx)
		if err != nil {
			return //todo
		}

		fmt.Fprintf(w, data)

	}
}

// store contains the method to fetch the data
type Store interface {
	Fetch(ctx context.Context) (string, error)
}
