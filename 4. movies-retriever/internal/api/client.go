package api

import (
	"encoding/json"
	"fmt"
	"movies/internal/models"
	"net/http"
	"net/url"
)

// Client handles API requests to the movie database
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewClient creates a new API client with the given base URL
func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL:    baseURL,
		HTTPClient: &http.Client{},
	}
}

// SearchMovies queries the API for movies matching the given title substring
func (c *Client) SearchMovies(titleSubstr string, page int) (*models.MovieResponse, error) {
	// Construct the query URL
	endpoint := fmt.Sprintf("%s/api/movies/search", c.BaseURL)

	// Create a request with query parameters
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Add query parameters
	q := url.Values{}
	q.Add("Title", titleSubstr)
	if page > 0 {
		q.Add("page", fmt.Sprintf("%d", page))
	}
	req.URL.RawQuery = q.Encode()

	// Make the request
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	// Check for successful response
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned non-200 status code: %d", resp.StatusCode)
	}

	// Decode the response
	var movieResp models.MovieResponse
	if err := json.NewDecoder(resp.Body).Decode(&movieResp); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &movieResp, nil
}
