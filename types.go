package main

// SberCloud API request data
type GPT3Data struct {
	Text string `json:"text"`
}

// SberCloud API error
type GPT3Error struct {
	Loc     []string `json:"loc"`
	Message string   `json:"msg"`
	Type    string   `json:"type"`
}

// SberCloud API answer
type GPT3Answer struct {
	Predictions string      `json:"predictions"`
	Detail      []GPT3Error `json:"detail"`
}
