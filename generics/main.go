package main

import "fmt"

func FilterInts(arr []int, predicate func(int) bool) []int {
	result := make([]int, 0)
	for _, n := range arr {
		if predicate(n) {
			result = append(result, n)
		}
	}
	return result
}

func FilterUsers(users []User, predicate func(User) bool) []User {
	result := make([]User, 0)
	for _, u := range users {
		if predicate(u) {
			result = append(result, u)
		}
	}
	return result
}

func Filter[T any](items []T, predicate func(T) bool) []T {
	result := make([]T, 0)
	for _, i := range items {
		if predicate(i) {
			result = append(result, i)
		}
	}
	return result
}

type User struct {
	Name string
	Age int
}

func main() {
	users := []User{
		User{"John", 25},
		User{"Cleven", 21},
		User{"Clara", 15},
	}
	fmt.Println(FilterInts([]int{13,123,12,65,32,124,43}, func(n int) bool {
		if n%2 == 0 {
			return true
		} else {
			return false
		}
	}))

	fmt.Println(FilterUsers(users, func(u User) bool {
		if u.Age >= 18 {
			return true
		} else {
			return false
		}
	}))

	fmt.Println("USING GENERICS: \n", Filter(users, func(u User) bool {
		if u.Age >= 18 {
			return true
		} else {
			return false
		}
	}))

	fmt.Println("USING GENERICS: \n", Filter([]int{13,123,12,65,32,124,43}, func(n int) bool {
		if n%2 == 0 {
			return true
		} else {
			return false
		}
	}))
}