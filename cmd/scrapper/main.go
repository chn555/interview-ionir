package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	luke, err := getPerson("1")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	filename, err := writeToCSV(luke)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Person info was saved to " + filename)
}

const swapiURL = "https://swapi.dev/api"

func writeToCSV(person Person) (string, error) {
	f, err := os.CreateTemp("", "luke")
	if err != nil {
		return "", fmt.Errorf("failed to create temp file: %w", err)
	}

	_, err = f.WriteString(person.Height)
	if err != nil {
		return "", fmt.Errorf("failed to write to file: %w", err)
	}

	return f.Name(), nil
}

func getPerson(id string) (Person, error) {
	url := fmt.Sprintf("%v/people/%v", swapiURL, id)

	resp, err := http.Get(url)
	if err != nil {
		return Person{}, fmt.Errorf("failed to get person info: %w", err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Person{}, fmt.Errorf("failed to read info: %w", err)
	}

	var person Person
	err = json.Unmarshal(body, &person)
	if err != nil {
		return Person{}, fmt.Errorf("failed to unmarshel person info: %w", err)
	}

	return person, nil
}

type Person struct {
	Name      string        `json:"name"`
	Height    string        `json:"height"`
	Mass      string        `json:"mass"`
	HairColor string        `json:"hair_color"`
	SkinColor string        `json:"skin_color"`
	EyeColor  string        `json:"eye_color"`
	BirthYear string        `json:"birth_year"`
	Gender    string        `json:"gender"`
	Homeworld string        `json:"homeworld"`
	Films     []string      `json:"films"`
	Species   []interface{} `json:"species"`
	Vehicles  []string      `json:"vehicles"`
	Starships []string      `json:"starships"`
	Created   time.Time     `json:"created"`
	Edited    time.Time     `json:"edited"`
	URL       string        `json:"url"`
}
