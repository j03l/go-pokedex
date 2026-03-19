package main

type PokeApi struct {
	Count     int    `json:"count"`
	Next      string `json:"next"`
	Previous  string `json:"previous"`
	Locations `json:"results"`
}

type Locations []struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
