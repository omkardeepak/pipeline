package main

import (
	"testing"
	"time"
)

func TestScript(t *testing.T) {
	t.Log("waiting for test script...will take 3 seconds")
	time.Sleep(3 * time.Second)
	t.Log("Test script executed")

}
