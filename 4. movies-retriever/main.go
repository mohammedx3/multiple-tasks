package main

import (
	"bufio"
	"fmt"
	"movies/internal/api"
	"movies/internal/services"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		// Print menu at the start of each loop
		printMainMenu()

		fmt.Print("\nEnter your choice (1-3): ")
		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)

		switch choiceStr {
		case "1":
			fmt.Print("Enter search term: ")
			searchTerm, _ := reader.ReadString('\n')
			searchTerm = strings.TrimSpace(searchTerm)

			fmt.Println("Searching for movies containing:", searchTerm)
			titles := getMovieTitles(searchTerm) // Using the challenge function

			if len(titles) == 0 {
				fmt.Println("No movies found matching your search.")
			} else {
				fmt.Printf("Found %d movies:\n", len(titles))
				for i, title := range titles {
					fmt.Printf("%3d. %s\n", i+1, title)
				}
			}

			// Pause before returning to menu
			fmt.Print("\nPress Enter to continue...")
			reader.ReadString('\n')

		case "2":
			fmt.Println("Fetching all movies (this might take a moment)...")
			titles := getMovieTitles("") // Using the challenge function with empty string

			if len(titles) == 0 {
				fmt.Println("No movies found.")
			} else {
				showPaginatedResults(titles)
			}

		case "3":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please enter 1, 2, or 3.")
			fmt.Print("\nPress Enter to continue...")
			reader.ReadString('\n')
		}
	}
}

// Helper function to print the main menu
func printMainMenu() {
	fmt.Println("\nMovie Title Search")
	fmt.Println("=================")
	fmt.Println("1. Search for movies by title")
	fmt.Println("2. List all movies (this might return a lot of results)")
	fmt.Println("3. Exit")
}

// Helper function to show paginated results
func showPaginatedResults(titles []string) {
	const pageSize = 10
	totalPages := (len(titles) + pageSize - 1) / pageSize
	currentPage := 1

	reader := bufio.NewReader(os.Stdin)

	for {
		// Calculate start and end indices for the current page
		start := (currentPage - 1) * pageSize
		end := start + pageSize
		if end > len(titles) {
			end = len(titles)
		}

		// Display current page
		fmt.Printf("\nShowing page %d of %d (%d total movies)\n", currentPage, totalPages, len(titles))
		fmt.Println("------------------------------------------------")

		for i := start; i < end; i++ {
			fmt.Printf("%3d. %s\n", i+1, titles[i])
		}

		// Navigation options
		fmt.Println("\nNavigation:")
		fmt.Println("n - Next page")
		fmt.Println("p - Previous page")
		fmt.Println("g - Go to page number")
		fmt.Println("q - Return to main menu")

		fmt.Print("\nEnter choice: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "n":
			if currentPage < totalPages {
				currentPage++
			} else {
				fmt.Println("Already on the last page.")
				fmt.Print("\nPress Enter to continue...")
				reader.ReadString('\n')
			}

		case "p":
			if currentPage > 1 {
				currentPage--
			} else {
				fmt.Println("Already on the first page.")
				fmt.Print("\nPress Enter to continue...")
				reader.ReadString('\n')
			}

		case "g":
			fmt.Printf("Enter page number (1-%d): ", totalPages)
			pageStr, _ := reader.ReadString('\n')
			pageStr = strings.TrimSpace(pageStr)
			var pageNum int
			_, err := fmt.Sscanf(pageStr, "%d", &pageNum)
			if err == nil && pageNum >= 1 && pageNum <= totalPages {
				currentPage = pageNum
			} else {
				fmt.Println("Invalid page number.")
				fmt.Print("\nPress Enter to continue...")
				reader.ReadString('\n')
			}

		case "q":
			return

		default:
			fmt.Println("Invalid choice.")
			fmt.Print("\nPress Enter to continue...")
			reader.ReadString('\n')
		}
	}
}

// The function required by the challenge
func getMovieTitles(substr string) []string {
	client := api.NewClient("https://jsonmock.hackerrank.com")
	movieService := services.NewMovieService(client)
	titles, err := movieService.GetMovieTitles(substr)
	if err != nil {
		return []string{}
	}
	return titles
}
