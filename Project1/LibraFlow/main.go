package main

import "fmt"

type Book struct {
	ID     int
	Title  string
	Author string
	Copies int
}

type User struct {
	ID          int
	Name        string
	IssuedBooks []int // // stores Book IDs the user has borrowed
}

type Library struct {
	Books map[int]*Book
	Users map[int]*User
}

func main() {
	fmt.Println("jai shree ram")

	lib := Library{
		Books: make(map[int]*Book),
		Users: make(map[int]*User),
	}


	lib.

}
