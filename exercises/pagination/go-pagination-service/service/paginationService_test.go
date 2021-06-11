package service

import (
	"service/pagination/model"
	"testing"
)

func ExamplePaginationservice() {
	Paginationservice(4, 5, 5, 1, 0)
	// Output: {"metadata":{"page":4,"per_page":5,"total":5},"pagination":"1 ... 4 5 ","items":[{"id":5,"name":"Buzz"},{"id":6,"name":"Fizz"},{"id":7,"name":""},{"id":8,"name":""},{"id":9,"name":"Fizz"}]}
}

func ExamplePaginationservice_two() {
	Paginationservice(4, 5, 10, 2, 2)
	// Output: {"metadata":{"page":4,"per_page":5,"total":10},"pagination":"1 2 3 4 5 6 ... 9 10 ","items":[{"id":5,"name":"Buzz"},{"id":6,"name":"Fizz"},{"id":7,"name":""},{"id":8,"name":""},{"id":9,"name":"Fizz"}]}
}

func TestPagination(t *testing.T) {
	expected := "1 ... 4 5 "
	if observed := pagination(4, 5, 1, 0); observed != expected {
		t.Fatalf("Got %v, want %v", observed, expected)
	}

	expected = "1 2 3 4 5 6 ... 9 10 "
	if observed := pagination(4, 10, 2, 2); observed != expected {
		t.Fatalf("Got %v, want %v", observed, expected)
	}

	expected = "1 2 ... 10 11 12 13 14 ... 19 20 "
	if observed := pagination(12, 20, 2, 2); observed != expected {
		t.Fatalf("Got %v, want %v", observed, expected)
	}
}

func TestProductList(t *testing.T) {
	expected := []model.Item{{Id: 5, Name: "Buzz"}, {Id: 6, Name: "Fizz"}, {Id: 7, Name: ""}, {Id: 8, Name: ""}, {Id: 9, Name: "Fizz"}}
	observed := productList(4, 5)
	for i := range expected {
		if observed[i] != expected[i] {
			t.Fatalf("Got %v, want %v", observed[i], expected[i])
		}
	}
}

func TestGenerateName(t *testing.T) {
	expected := ""
	if observed := generateName(2); observed != expected {
		t.Fatalf("generateName(2)) got %v, want %v", observed, expected)
	}

	expected = "Fizz"
	if observed := generateName(3); observed != expected {
		t.Fatalf("generateName(3)) got %v, want %v", observed, expected)
	}

	expected = "Buzz"
	if observed := generateName(5); observed != expected {
		t.Fatalf("generateName(5)) got %v, want %v", observed, expected)
	}

	expected = "BuzzZazz"
	if observed := generateName(15); observed != expected {
		t.Fatalf("generateName(15)) got %v, want %v", observed, expected)
	}
}
