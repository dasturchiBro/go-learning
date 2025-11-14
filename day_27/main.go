package main

import (
	"fmt"
	"errors"
)

func Div(numF, numS int) (int, error) {
	if numS == 0 {
		return 0, errors.New("can't be divided by 0")
	}
	return numF/numS, nil
}

type User struct {
	Name string
	Age int
}



func findUser(users map[int]User, user_id int) (User, error) {
	user, ok := users[user_id]
	if !ok {
		return User{}, errors.New("user not found")
	}
	return user, nil
}

func main() {
	// Errors
	result, err := Div(100,0)
	if err == nil {
		fmt.Println("Everything works well: ", result)
	} else {
		fmt.Println("Something is going wrong: ", err)
	}

	users := map[int]User {
		1: User{"John", 17},
		2: User{"Smith", 45},
		3: User{"Hurry", 34},
	}
	fmt.Println("\n")

	user, err := findUser(users, 3)
	if err != nil {
		fmt.Println("Something went wrong: ", err)
	} else {
		fmt.Println("User found: ", user)
	}

	user, err = findUser(users, 300)
	if err != nil {
		fmt.Println("Something went wrong: ", err)
	} else {
		fmt.Println("User found: ", user)
	}
}