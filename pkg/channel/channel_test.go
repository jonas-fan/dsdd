package channel

import (
	"sort"
	"testing"
)

func push(to chan int, data ...int) {
	defer close(to)

	for _, each := range data {
		to <- each
	}
}

func pull(from chan int) []int {
	out := make([]int, 0)

	for each := range from {
		out = append(out, each)
	}

	return out
}

func makeStreams(size int) []chan int {
	streams := make([]chan int, size)

	for i := 0; i < size; i++ {
		streams[i] = make(chan int)
	}

	return streams
}

func concat(slice [][]int) []int {
	out := make([]int, 0)

	for _, each := range slice {
		out = append(out, each...)
	}

	return out
}

func hasSameElement(lhs []int, rhs []int) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	sort.Ints(lhs)
	sort.Ints(rhs)

	for i := range lhs {
		if lhs[i] != rhs[i] {
			return false
		}
	}

	return true
}

func testMultiplex(data [][]int) bool {
	in := makeStreams(len(data))
	out := make(chan int)

	for i := range data {
		go push(in[i], data[i]...)
	}

	Multiplex(out, in)

	return hasSameElement(pull(out), concat(data))
}

func TestMultiplex1To1(t *testing.T) {
	data := [][]int{
		{0x00, 0x01, 0x02, 0x03},
	}

	if testMultiplex(data) {
		t.Log("success")
	} else {
		t.Error("failure")
	}
}

func TestMultiplex2To1(t *testing.T) {
	data := [][]int{
		{0x00, 0x01, 0x02, 0x03},
		{0x04, 0x05, 0x06, 0x07},
	}

	if testMultiplex(data) {
		t.Log("success")
	} else {
		t.Error("failure")
	}
}

func TestMultiplex4To1(t *testing.T) {
	data := [][]int{
		{0x00, 0x01, 0x02, 0x03},
		{0x04, 0x05, 0x06, 0x07},
		{0x08, 0x09, 0x0a, 0x0b},
		{0x0c, 0x0d, 0x0e, 0x0f},
		{0x10, 0x11, 0x12, 0x13},
	}

	if testMultiplex(data) {
		t.Log("success")
	} else {
		t.Error("failure")
	}
}

func TestMultiplex8To1(t *testing.T) {
	data := [][]int{
		{0x00, 0x01, 0x02, 0x03},
		{0x04, 0x05, 0x06, 0x07},
		{0x08, 0x09, 0x0a, 0x0b},
		{0x0c, 0x0d, 0x0e, 0x0f},
		{0x10, 0x11, 0x12, 0x13},
		{0x14, 0x15, 0x16, 0x17},
		{0x18, 0x19, 0x1a, 0x1b},
		{0x1c, 0x1d, 0x1e, 0x1f},
	}

	if testMultiplex(data) {
		t.Log("success")
	} else {
		t.Error("failure")
	}
}

func TestMultiplex16To1(t *testing.T) {
	data := [][]int{
		{0x00, 0x01, 0x02, 0x03},
		{0x04, 0x05, 0x06, 0x07},
		{0x08, 0x09, 0x0a, 0x0b},
		{0x0c, 0x0d, 0x0e, 0x0f},
		{0x10, 0x11, 0x12, 0x13},
		{0x14, 0x15, 0x16, 0x17},
		{0x18, 0x19, 0x1a, 0x1b},
		{0x1c, 0x1d, 0x1e, 0x1f},
		{0x20, 0x21, 0x22, 0x23},
		{0x24, 0x25, 0x26, 0x27},
		{0x28, 0x29, 0x2a, 0x2b},
		{0x2c, 0x2d, 0x2e, 0x2f},
		{0x30, 0x31, 0x32, 0x33},
		{0x34, 0x35, 0x36, 0x37},
		{0x38, 0x39, 0x3a, 0x3b},
		{0x3c, 0x3d, 0x3e, 0x3f},
	}

	if testMultiplex(data) {
		t.Log("success")
	} else {
		t.Error("failure")
	}
}
