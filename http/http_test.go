package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	players map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) string {
	score, status := s.players[name]
	if !status {
		return ""
	}
	return fmt.Sprint(score)
}

func TestGetPlayer(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"pepper": 20,
			"floyd":  10,
		},
	}
	server := &PlayerServer{store: &store}
	t.Run("get Pepper's score", func(t *testing.T) {
		want := "20"
		//request
		request, _ := http.NewRequest(http.MethodGet, "/players/pepper", nil)
		//mock of the response
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		assertStatus(http.StatusOK, response.Code, t)
		assertResponse(want, got, t)

	})

	t.Run("get Flyod's score", func(t *testing.T) {
		want := "10"
		//request
		request := getNewRequest("floyd")
		//mock of the response
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()

		assertStatus(http.StatusOK, response.Code, t)
		assertResponse(want, got, t)

	})

	t.Run("returns 404 on missing Player", func(t *testing.T) {

		//request
		request := getNewRequest("apollo")
		//mock of the response
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(http.StatusNotFound, response.Code, t)
	})
}

func getNewRequest(player string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
	return request
}

func assertResponse(want string, got string, t testing.TB) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertStatus(want int, got int, t testing.TB) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}
