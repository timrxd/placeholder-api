package testing

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/timrxd/placeholder-api/api"
)

func TestAPI(t *testing.T) {
	cases := []struct {
		method string
		path   string
		body   api.Item

		outCode int
		outBody string
	}{
		// Add first item to database
		{"POST", "/item", api.Item{UserID: 1, ID: 1, Title: "Lorum Ipsum", Body: "dolor sit amet"}, 201,
			`{"id":1,"userId":1,"title":"Lorum Ipsum","body":"dolor sit amet"}`},

		// Add second item to database
		{"POST", "/item", api.Item{UserID: 1, ID: 2, Title: "Lorum Ipsum 2", Body: "dolor sit amet 2"}, 201,
			`{"id":2,"userId":1,"title":"Lorum Ipsum 2","body":"dolor sit amet 2"}`},

		// Add third item to database
		{"POST", "/item", api.Item{UserID: 1, ID: 3, Title: "Lorum Ipsum 3", Body: "dolor sit amet 3"}, 201,
			`{"id":3,"userId":1,"title":"Lorum Ipsum 3","body":"dolor sit amet 3"}`},

		// Get first item
		{"GET", "/item/1", api.Item{}, 200,
			`{"id":1,"userId":1,"title":"Lorum Ipsum","body":"dolor sit amet"}`},

		// Remove first item
		{"DELETE", "/item/1", api.Item{}, 200, `"Item deleted"`},

		// Update second item
		{"PUT", "/item/2", api.Item{UserID: 1, ID: 2, Title: "Lorum Ipsum 4", Body: "dolor sit amet 4"}, 200,
			`{"id":2,"userId":1,"title":"Lorum Ipsum 4","body":"dolor sit amet 4"}`},

		// Get all items
		{"GET", "/items", api.Item{}, 200,
			`[{"id":2,"userId":1,"title":"Lorum Ipsum 4","body":"dolor sit amet 4"},{"id":3,"userId":1,"title":"Lorum Ipsum 3","body":"dolor sit amet 3"}]`},
	}

	server := api.CreateServer()
	for _, c := range cases {
		response := httptest.NewRecorder()
		body, _ := json.Marshal(c.body)
		request, _ := http.NewRequest(c.method, c.path, bytes.NewBuffer(body))
		server.ServeHTTP(response, request)

		b := response.Body.String()
		if b[:len(b)-1] != c.outBody {
			t.Errorf("Body Mismatch: %s %s\nResponse\t%s\nExpected\t%s", c.method, c.path, b[:len(b)-1], c.outBody)
		} else if response.Code != c.outCode {
			t.Errorf("Response Code Mismatch: %s %s\nResponse\t%d\nExpected\t%d", c.method, c.path, response.Code, c.outCode)
		}
	}
}
