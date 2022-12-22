package rick

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    pokedex, err := UnmarshalPokedex(bytes)
//    bytes, err = pokedex.Marshal()

import "encoding/json"

func UnmarshalRMCharacter(data []byte) (RMCharacter, error) {
	var r RMCharacter
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RMCharacter) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ApiResponse struct {
	Info    Info     `json:"info"`   
	Results []RMCharacter `json:"results"`
}

type Info struct {
	Count int64       `json:"count"`
	Pages int64       `json:"pages"`
	Next  string      `json:"next"` 
	Prev  interface{} `json:"prev"` 
}

type RMCharacter struct {
	ID       int64    `json:"id"`      
	Name     string   `json:"name"`    
	Status   Status   `json:"status"`  
	Species  Species  `json:"species"` 
	Type     Type     `json:"type"`    
	Gender   Gender   `json:"gender"`  
	Origin   Location `json:"origin"`  
	Location Location `json:"location"`
	Image    string   `json:"image"`   
	Episode  []string `json:"episode"` 
	URL      string   `json:"url"`     
	Created  string   `json:"created"` 
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"` 
}

type Gender string
const (
	Male Gender = "Male"
)

type Species string
const (
	Alien Species = "Alien"
	Cronenberg Species = "Cronenberg"
	Human Species = "Human"
	Humanoid Species = "Humanoid"
)

type Status string
const (
	Alive Status = "Alive"
	Dead Status = "Dead"
	Unknown Status = "unknown"
)

type Type string
const (
	Empty Type = ""
	FishPerson Type = "Fish-Person"
	HumanWithAntennae Type = "Human with antennae"
	Robot Type = "Robot"
)
