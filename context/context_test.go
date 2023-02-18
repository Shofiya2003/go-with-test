package context

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	t        *testing.T
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled 1")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}

		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		log.Println("spy store got cancelled")
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}

}

func TestServer(t *testing.T) {
	t.Run("return data from the store", func(t *testing.T) {
		data := "hello world"
		store := SpyStore{response: data, t: t}
		svr := Server(&store)
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)
		log.Printf("got %s ", response.Body.String())
		if response.Body.String() != data {
			t.Errorf("got %s expected %s ", response.Body.String(), data)
		}

	})
	t.Run("tells store to cancel the work when request is cancelled", func(t *testing.T) {

		data := "hello world"
		store := SpyStore{response: data, t: t}
		svr := Server(&store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingContext, cancel := context.WithCancel(request.Context())
		//wait for 5 millisecond and then cancel the request
		time.AfterFunc(1*time.Millisecond, cancel)

		request = request.WithContext(cancellingContext)
		response := &SpyResponseWriter{}
		svr.ServeHTTP(response, request)

		if response.written {
			t.Errorf("a response should not have been written")
		}

	})
}
