package main

import (
	"fmt"
	"errors"
	"os"
)

type Status int  

const (
	InvalidStatus Status = 422
	NotFound Status = 404
)


type StatusErr struct {
	Status Status
	Message string
}

func (s StatusErr) Error() string {
	return s.Message
}

type UserType map[string]string

func login(key string, pass string, users UserType) (string, error) {
	value, ok := users[key]
	if !ok || value != pass {
		return "", errors.New(fmt.Sprintf("invalid credentials for user with id %s", key))
	}
	return key, nil
} 

func getInfoById(key string, infoById UserType) (string, error) {
	value, ok := infoById[key]
	if !ok {
		return "", errors.New("couldn't get information about this user")
	}
	return value, nil
} 


func LoginAndGetData(users UserType, key string, pass string, infoById UserType) (string, error) {
	isOkay, err := login(key, pass, users)
	if err != nil {
		return "", StatusErr{
			Status: InvalidStatus,
			Message: fmt.Sprintf("invalid credentials for user with id %s", key),
		}
	}

	info, err := getInfoById(isOkay, infoById)
	if err != nil {
		return "", StatusErr{
			Status: NotFound,
			Message: fmt.Sprintf("couldn't get information about this user"),
		}
	}
	return info, nil
}

func test(users UserType, key string, pass string, infoById UserType) {
	info, err := LoginAndGetData(users, key, pass, infoById)
	if err != nil {
		fmt.Println("Something went wrong: ", err, "\n")	
	} else {
		fmt.Println("\t", info, "\n")
	}
}

func fileChecker(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("in fileChecker: %w", err)
	}
	f.Close()
	return nil
}



func main() {
	users := UserType{
		"$1":"123456",
		"$2":"asdfghjkl;12356",
		"$3":"john_12346",
		"$4":"helper_2009002",
		"$5":"myhome_12345",
	}
	information := UserType{
		"$1":"Sherlock Holmes",
		"$2":"Peter Johnson",
		"$3":"John Doe",
		"$4":"Mr Smith",
	}
	test(users, "$3", "john_12346", information)
	test(users, "$3", "12345-password err", information)
	test(users, "$5", "myhome_12345", information)

	
	if err := fileChecker("not_a_file.txt"); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("That file doesn't exist!")
		}
	}
}