package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type void struct{}

var member void

func TestGenerateOpinion(t *testing.T) {

	cases := []struct {
		data     map[string][]string
		expected map[Opinion]void
	}{
		{
			data: map[string][]string{
				"adjective": {
					"foo",
				},
				"subject": {
					"bar",
				},
				"object": {
					"baz",
				},
			},
			expected: map[Opinion]void{
				{Text: "foo bar baz."}: member,
			},
		},
		{
			data: map[string][]string{
				"adjective": {
					"foo",
					"foo2",
				},
				"subject": {
					"bar",
					"bar2",
				},
				"object": {
					"baz",
					"baz2",
				},
			},
			expected: map[Opinion]void{
				{Text: "foo bar baz."}:    member,
				{Text: "foo bar baz2."}:   member,
				{Text: "foo bar2 baz."}:   member,
				{Text: "foo bar2 baz2."}:  member,
				{Text: "foo2 bar baz."}:   member,
				{Text: "foo2 bar baz2."}:  member,
				{Text: "foo2 bar2 baz."}:  member,
				{Text: "foo2 bar2 baz2."}: member,
			},
		},
	}

	for _, tc := range cases {
		got := generateOpinion(tc.data)
		if _, ok := tc.expected[got]; !ok {

			t.Errorf("generateOpinion(%s): Expected: One of %s, got %s", tc.data, tc.expected, got)
		}
	}
}

func TestOpinionHandler(t *testing.T) {
	tt := []struct {
		method string
		status int
		body   string
	}{
		{
			method: "GET",
			status: http.StatusOK,
			body:   `{"text":"foo bar baz."}`,
		},
		{
			method: "POST",
			status: http.StatusMethodNotAllowed,
			body:   "Method not allowed\n",
		},
		{
			method: "PUT",
			status: http.StatusMethodNotAllowed,
			body:   "Method not allowed\n",
		},
		{
			method: "HEAD",
			status: http.StatusMethodNotAllowed,
			body:   "Method not allowed\n",
		},
		{
			method: "OPTIONS",
			status: http.StatusMethodNotAllowed,
			body:   "Method not allowed\n",
		},
		{
			method: "DELETE",
			status: http.StatusMethodNotAllowed,
			body:   "Method not allowed\n",
		},
		{
			method: "PATCH",
			status: http.StatusMethodNotAllowed,
			body:   "Method not allowed\n",
		},
	}

	data := map[string][]string{
		"adjective": {
			"foo",
		},
		"subject": {
			"bar",
		},
		"object": {
			"baz",
		},
	}

	handler := OpinionHandler{Data: data}

	for _, tc := range tt {
		request := httptest.NewRequest(tc.method, "/", nil)
		responseRecorder := httptest.NewRecorder()

		handler.ServeHTTP(responseRecorder, request)

		if status := responseRecorder.Code; status != tc.status {
			t.Errorf("OpinionHandler.ServeHTTP: wrong status code: got %v want %v", status, tc.status)
		}

		if body := responseRecorder.Body.String(); body != tc.body {
			t.Errorf("OpinionHandler.ServeHTTP: wrong body: got %v want %v", body, tc.body)
		}
	}
}
