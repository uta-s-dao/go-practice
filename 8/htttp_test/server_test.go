package main

import (
	"encoding/json"
	. "gopkg.in/check.v1"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type PostTestSuite struct {
	mux    *http.ServeMux
	post   *FakePost
	writer *httptest.ResponseRecorder
}

func init() {
	Suite(&PostTestSuite{})
}

func Test(t *testing.T) { TestingT(t) }

// func (s *PostTestSuite) TestHandleGet(c *C) {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/post/", handleRequest(&FakePost{}))
// 	writer := httptest.NewRecorder()
// 	request, _ := http.NewRequest("GET", "/post/1", nil)
// 	mux.ServeHTTP(writer, request)

// 	c.Check(writer.Code, Equals, 200)
// 	var post Post
// 	json.Unmarshal(writer.Body.Bytes(), &post)
// 	c.Check(post.Id, Equals, 1)

// }

func (s *PostTestSuite) SetUpTest(c *C) {
	s.post = &FakePost{}
	s.mux = http.NewServeMux()
	s.mux.HandleFunc("/post/", handleRequest(s.post))
	s.writer = httptest.NewRecorder()
}

func (s *PostTestSuite) TestGetPost(c *C) {
	request, _ := http.NewRequest("GET", "/post/1", nil)
	s.mux.ServeHTTP(s.writer, request)

	c.Check(s.writer.Code, Equals, 200)
	var post Post
	json.Unmarshal(s.writer.Body.Bytes(), &post)
	c.Check(post.Id, Equals, 1)
}

func (s *PostTestSuite) TestPutPost(c *C) {
	json := strings.NewReader(`{"content":"Update post","author":"Yuta"}`)
	request, _ := http.NewRequest("PUT", "/post/5", json)
	s.mux.ServeHTTP(s.writer, request)

	c.Check(s.writer.Code, Equals, 200)
	c.Check(s.post.Id, Equals, 5)
	c.Check(s.post.Content, Equals, "Update post")
}

func (s *PostTestSuite) TearDownTest(c *C) {
	c.Log("Finiished test -", c.TestName())
}

func (s *PostTestSuite) SetUpSuite(c *C) {
	c.Log("Starting Post tests")
}

func (s *PostTestSuite) TearDownSuite(c *C) {
	c.Log("Finiished all tests")
}
