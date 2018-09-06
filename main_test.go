package main

import "testing"

func TestStartQuery(t *testing.T) {
	StartQuery()
}

func TestMain(t *testing.T) {
	go main()
}