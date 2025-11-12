package earray

import (
	"reflect"
	"testing"
)

// 定义结构体 User 和 Order
type User struct {
	ID   string
	Name string
}

type Order struct {
	ID       string
	UserID   string
	ItemName string
}

func TestLeftJoin(t *testing.T) {
	users := []User{
		{ID: "1", Name: "Alice"},
		{ID: "2", Name: "Bob"},
		{ID: "3", Name: "Charlie"},
	}

	orders := []Order{
		{ID: "1", UserID: "1", ItemName: "Book"},
		{ID: "2", UserID: "1", ItemName: "Pen"},
		{ID: "3", UserID: "2", ItemName: "Notebook"},
	}

	expected := []LeftJoinResult[User, Order]{
		{Left: users[0], Right: []Order{orders[0], orders[1]}},
		{Left: users[1], Right: []Order{orders[2]}},
		{Left: users[2], Right: []Order{}},
	}

	result := LeftJoin(users, orders, func(u User) string { return u.ID }, func(o Order) string { return o.UserID })

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("LeftJoin() = %v; want %v", result, expected)
	}
}

func TestRightJoin(t *testing.T) {
	users := []User{
		{ID: "1", Name: "Alice"},
		{ID: "2", Name: "Bob"},
	}

	orders := []Order{
		{ID: "1", UserID: "1", ItemName: "Book"},
		{ID: "2", UserID: "3", ItemName: "Pen"},
	}

	expected := []RightJoinResult[User, Order]{
		{Left: []User{users[0]}, Right: orders[0]},
		{Left: []User{}, Right: orders[1]},
	}

	result := RightJoin(users, orders, func(u User) string { return u.ID }, func(o Order) string { return o.UserID })

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("RightJoin() = %v; want %v", result, expected)
	}
}

func TestDistinctSimple(t *testing.T) {
	input := []int{1, 2, 2, 3, 4, 4, 5}
	expected := []int{1, 2, 3, 4, 5}

	result := Distinct(input, func(i int) int { return i })

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Distinct() = %v; want %v", result, expected)
	}
}

func TestDistinctStruct(t *testing.T) {
	orders := []Order{
		{ID: "1", UserID: "1", ItemName: "Book"},
		{ID: "2", UserID: "1", ItemName: "Book"}, // Duplicate order by ID
		{ID: "3", UserID: "2", ItemName: "Pen"},
		{ID: "1", UserID: "1", ItemName: "Book"}, // Another duplicate by ID
	}

	expected := []Order{
		{ID: "1", UserID: "1", ItemName: "Book"},
		{ID: "2", UserID: "1", ItemName: "Book"},
		{ID: "3", UserID: "2", ItemName: "Pen"},
	}

	result := Distinct(orders, func(o Order) string { return o.ID })

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("DistinctStruct() = %v; want %v", result, expected)
	}
}
