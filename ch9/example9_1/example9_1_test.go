package main

import (
	"testing"
	"time"
)

func TestPrint1(t *testing.T) {
	print1()
}

func TestGoPrint1(t *testing.T) {
	goPrint1()
	time.Sleep(1 * time.Millisecond)
}

func TestPrint2(t *testing.T) {
	print2()
}

func TestGoPrint2(t *testing.T) {
	goPrint2()
	time.Sleep(20 * time.Millisecond)
}
