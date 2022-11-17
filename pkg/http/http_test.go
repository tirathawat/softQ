package http

import (
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestMakeCall(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{"id": 1, "name": "test", "info": "test info"}`)
	}))

	defer server.Close()

	want := &Response{
		ID:   1,
		Name: "test",
		Info: "test info",
	}

	t.Run("should return correctly response", func(t *testing.T) {
		got, err := MakeCall(server.URL)
		if err != nil {
			t.Errorf("got error %v, want nil", err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

	})
}
