package models

// Movie represents a movie from the API response
type Movie struct {
	Title  string `json:"Title"`
	Year   int    `json:"Year"`
	ImdbID string `json:"imdbID"`
}

// MovieResponse represents the JSON response from the API
type MovieResponse struct {
	Page       int     `json:"page"`
	PerPage    int     `json:"per_page"`
	Total      int     `json:"total"`
	TotalPages int     `json:"total_pages"`
	Data       []Movie `json:"data"`
}
