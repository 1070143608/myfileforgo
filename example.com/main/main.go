package main

import (
	"example.com/house"
	"example.com/study"
	"fmt"
)

func main() {

	// example 1
	name := "Tom"
	s, err := study.New(name)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s.Listen("english"))
	fmt.Println(s.Speak("english"))
	fmt.Println(s.Read("english"))
	fmt.Println(s.Write("english"))

	// example 2
	h := house.NewHouse(
		house.WithConcrete(),
		house.WithoutFireplace(),
		house.WithFloors(3),
	)
	fmt.Println(h.Material)

	// example 3
	friend := house.FindFriend(
		house.WithWhere("shanghai"),
		house.WithSex(1),
		house.WithAge(20),
		)
	fmt.Println(friend.Where)
}