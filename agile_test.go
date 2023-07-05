package main

import (
	"testing"
)

func TestPutToken01(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.put(1, 1, "o")

	if b.get(1, 1) != "o" {
		t.Errorf("......")
	}
}

func TestPutToken02(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.put(0, 0, "x")

	if b.get(0, 0) != "x" {
		t.Errorf("Error")
	}
}

func TestPutToken03(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.put(0, 0, "x")
	b.put(0, 0, "o")

	if b.get(0, 0) != "x" {
		t.Errorf("Error")
	}
}

func TestPutToken04(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.put(0, 0, "o")
	b.put(1, 0, "o")
	b.put(2, 0, "o")

	if b.judge() != "o win" {
		t.Errorf("Error")
	}
}

func TestPutToken05(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.put(0, 0, "x")
	b.put(1, 0, "x")
	b.put(2, 0, "x")

	if b.judge() != "x win" {
		t.Errorf("Error")
	}
}

func TestPutToken06(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.put(3, 0, "x")
	b.put(4, 0, "x")
	b.put(5, 0, "x")

	if b.judge() != "x win" {
		t.Errorf("Error")
	}
}
func TestPutToken07(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.put(0, 0, "x")
	b.put(0, 1, "x")
	b.put(0, 2, "x")

	if b.judge() != "x win" {
		t.Errorf("Error")
	}
}
