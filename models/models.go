package models

// Cases model for cases
type Cases struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Cases     string `json:"cases"`
	Recovered string `json:"recovered"`
	Deaths    string `json:"deaths"`
}
