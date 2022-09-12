package helper

import (
	"fmt"
	"testing"
)

func TestHelloWord(t *testing.T) {
	result := HelloWord("Golang")
	if result != "Hello Golang" {
		t.Fail()
	}
	fmt.Println("TestHelloWord Done")
}

func TestHelloWordDoom(t *testing.T) {
	result := HelloWord("Golang")
	if result != "Hello Golang" {
		panic("Result must be 'Hello Golang'")
	}
}
