package main

import (
	"errors"
	"time"
)

type Club struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

func Fetch(id int) (Club, error) {
	// Simulate API call takes a second to complete.
	time.Sleep(1 * time.Second)

	switch id {
	case 1:
		return Club{
			ID:      1,
			Name:    "Arsenal",
			Country: "England",
		}, nil
	case 2:
		return Club{
			ID:      2,
			Name:    "Barcelona",
			Country: "Spain",
		}, nil
	case 3:
		return Club{
			ID:      3,
			Name:    "AC Milan",
			Country: "Italy",
		}, nil
	default:
		return Club{}, errors.New("club not found")
	}
}
