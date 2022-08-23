package signal

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		t.Fatalf("http.NewRequest() err = %s", err)
	}
	Handler(w, r)

	response := w.Result()
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		t.Fatalf("Handler() status = %d; want %q", response.StatusCode, "200-299")
	}

	contentType := response.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Handler() Content-Type = %q; want %q", contentType, "application/json")
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("ioutil.ReadAll(response.Body) err = %s", err)
	}
	t.Logf("data: %#v", data)

	var p Person
	err = json.Unmarshal(data, &p)
	if err != nil {
		t.Fatalf("json.Unmarshal(response.Body) err = %s", err)
	}
	t.Logf("Person: %#v", p)

	wantAge := 30
	if p.Age != wantAge {
		t.Errorf("Person.Age = %d; want %d", p.Age, wantAge)
	}
	wantName := "Bob Jones"
	if p.Name != wantName {
		t.Errorf("Person.Name = %q; want %q", p.Name, wantName)
	}
}
