package main

import (
	"testing"
	"time"
)

func TestDecode(t *testing.T) {
	post, err := decode("post.json")
	if err != nil {
		t.Error(err)
	}
	if post.Id != 1 {
		t.Error("Wrong id, was expecting 1 but got", post.Id)
	}
	if post.Content != "Hello World" {
		t.Error("Wrong content,was expecting 'Hello World'but got", post.Content)
	}
	if post.Author.Id != 2 {
		t.Error("Wrong author id, was expecting 2 but got", post.Author.Id)
	}

}

func TestEncode(t *testing.T) {
	t.Skip("Skipping TestEncode")
}

func TestLongRunning(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	time.Sleep(10 * time.Second)
}
