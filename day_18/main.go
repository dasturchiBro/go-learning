package main

import ("fmt")

func main() {
	staff := map[string]int{
		"Bob": 10,
		"Sara": 4,
		"John": 3,
		"Peter": 5,
		"William": 4,
		"Thomas": 2,
		"Tom": 7,
		"Sherlock": 1,
		"Clara": 5,
		"Jonathan": 1,
	}

	names := []string{"Bob", "Sara", "John", "Peter", "William", "Tom", "Clara", "Sherlock", "Thomas", "Jonathan"}
	// Switch 
	for _, name := range names {
		switch role := staff[name]; {
		case role == 3:
			fmt.Printf("*%s is a junior software engineer.\n",name)
		case role == 4:
			fmt.Printf("***%s is a web designer.\n",name)
		case role == 5:
			fmt.Printf("**%s is a project manager.\n",name)
		case role == 7:
			fmt.Printf("-%s is a CTO.\n",name)
		case role == 10:
			fmt.Printf("--%s is a CEO.\n",name)
		default:
			fmt.Printf("---%s is an intern.\n",name)
		}
	}
}