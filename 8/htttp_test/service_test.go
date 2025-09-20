package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleGet(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest)

	writer := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/post/11", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != http.StatusOK {
		t.Error("Was expecting status code 200 but got", writer.Code)
	}

	var post Post

	json.Unmarshal(writer.Body.Bytes(), &post)
	if post.Id != 11 {
		t.Error("Was expecting post id 1 but got", post.Id)
	}
}

func TestHandlePut(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest)

	writer := httptest.NewRecorder()
	json := strings.NewReader(`{"content":"Update post","author":"Joe"}`)

	request, _ := http.NewRequest("PUT", "/post/11", json)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Error("Was expecting status code 200 but got", writer.Code)
	}
}
