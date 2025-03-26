package services

import (
	"movies/internal/api"
	"sort"
)

// MovieService provides operations related to movies
type MovieService struct {
	Client *api.Client
}

// NewMovieService creates a new movie service with the given client
func NewMovieService(client *api.Client) *MovieService {
	return &MovieService{
		Client: client,
	}
}

// GetMovieTitles retrieves and sorts movie titles containing the given substring
func (s *MovieService) GetMovieTitles(substr string) ([]string, error) {
	titles := []string{}
	page := 1

	// Fetch results from all pages
	for {
		// Get one page of results
		resp, err := s.Client.SearchMovies(substr, page)
		if err != nil {
			return nil, err
		}

		// Extract titles from this page
		for _, movie := range resp.Data {
			titles = append(titles, movie.Title)
		}

		// Check if we've reached the last page
		if page >= resp.TotalPages {
			break
		}

		// Go to the next page
		page++
	}

	// Sort the titles in ascending order
	sort.Strings(titles)

	return titles, nil
}
