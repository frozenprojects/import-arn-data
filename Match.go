package main

// Match ...
type Match struct {
	ID            int     `json:"id"`
	ProviderID    int     `json:"providerId"`
	Title         string  `json:"title"`
	ProviderTitle string  `json:"providerTitle"`
	Similarity    float64 `json:"similarity"`
	Edited        string  `json:"edited"`
	EditedBy      string  `json:"editedBy"`
}

// Match2 ...
type Match2 struct {
	ID            int     `json:"id"`
	ProviderID    string  `json:"providerId"`
	Title         string  `json:"title"`
	ProviderTitle string  `json:"providerTitle"`
	Similarity    float64 `json:"similarity"`
	Edited        string  `json:"edited"`
	EditedBy      string  `json:"editedBy"`
}

// MatchNyaa ...
type MatchNyaa struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Edited   string `json:"edited"`
	EditedBy string `json:"editedBy"`
}
