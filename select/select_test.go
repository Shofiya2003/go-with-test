package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("test the racer function", func(t *testing.T) {
		slowServer := createTestServer(20 * time.Millisecond)

		fastServer := createTestServer(0 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()
		slowURL := slowServer.URL
		fastURL := fastServer.URL

		got, err := Racer(slowURL, fastURL)
		want := fastURL

		if err != nil {
			t.Errorf("got an unexpected err %v", err)
		}
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("error if server does'nt respond within 10 seconds", func(t *testing.T) {
		slowServer := createTestServer(12 * time.Millisecond)

		defer slowServer.Close()

		slowURL := slowServer.URL

		_, err := ConfigurableRacer(slowURL, slowURL, 10*time.Millisecond)

		if err == nil {
			t.Errorf("wanted an error but did not get one")
		}
	})

}

func createTestServer(duration time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(duration)
		w.WriteHeader(http.StatusOK)
	}))
}
