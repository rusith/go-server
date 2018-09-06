package main

import "testing"

func TestStartQuery(t *testing.T) {
	StartQuery()
}

func TestMainFunc(t *testing.T) {
	go main()
}