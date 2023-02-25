package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	players  map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) string {
	score, status := s.players[name]
	if !status {
		return ""
	}
	return fmt.Sprint(score)
}

func (s *StubPlayerStore) RecordWins(name string) {
	s.winCalls = append(s.winCalls, name)
}
func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
	}

	server := PlayerServer{&store}

	t.Run("returns accepted on POST", func(t *testing.T) {
		player := "bob"
		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(http.StatusAccepted, response.Code, t)
		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], player)
		}

	})
}
func TestGetPlayer(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"pepper": 20,
			"floyd":  10,
		},
		nil,
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

func newPostWinRequest(player string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", player), nil)
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
