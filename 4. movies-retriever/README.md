# Movies Retriever

A Go-based CLI application that allows users to search for movies by title or list all available movies using an external movie database API.

---

## How It Works

1. **Input**:
   - Users interact with the CLI-based menu to search for movies by title or list all movies.

2. **API Integration**:
   - The application uses a movie database API (`https://jsonmock.hackerrank.com/api/movies/search`) to fetch movie data.
   - Titles are retrieved across all pages of the API response.

3. **Movie Search**:
   - Users can search for movies containing a specific substring in the title.
   - The results are sorted alphabetically.

4. **Pagination**:
   - For large datasets, the application supports paginated results, allowing users to navigate through pages.

5. **Output**:
   - Displays movie titles in a user-friendly format, either as a complete list or paginated.

---

## Usage

1. **Run the Application**:
```bash
go run main.go
```
2. **Menu Options**:
```bash
1: Search for movies by title
2: List all movies (this might return a lot of results)
3: Exit
```
3. **Navigation**:
For paginated results:
```bash
n: Next page.
p: Previous page.
g: Go to a specific page.
q: Return to the main menu.
```

## Key Components
### Main Application (main.go)
Provides a menu-based interface for the user.
Handles user input, calls the movie service, and displays results.
Supports pagination for large datasets.
### Movie Service (internal/services/movie.go)
Fetches movie titles from the movie database API.
Retrieves and sorts titles containing the user-provided substring.
Handles pagination to fetch all pages of results.
### API Client (internal/api/client.go)
Handles HTTP requests to the movie database API.
Constructs query parameters and processes API responses.
### Data Models (internal/models/movie.go)
Defines data structures for API responses (Movie and MovieResponse).

## Output
### Main Menu

```bash
Movie Title Search
=================
1. Search for movies by title
2. List all movies (this might return a lot of results)
3. Exit

Enter your choice (1-3): 1
```

### Searching for Movies
```bash
Enter search term: spiderman
Searching for movies containing: spiderman
Found 13 movies:
  1. Amazing Spiderman Syndrome
  2. Fighting, Flying and Driving: The Stunts of Spiderman 3
  3. Hollywood's Master Storytellers: Spiderman Live
  4. Italian Spiderman
  5. Spiderman
  6. Spiderman
  7. Spiderman 5
  8. Spiderman and Grandma
  9. Spiderman in Cannes
 10. Superman, Spiderman or Batman
 11. The Amazing Spiderman T4 Premiere Special
 12. The Death of Spiderman
 13. They Call Me Spiderman
```

### Paginated Results
```bash
Showing page 20 of 277 (2770 total movies)
------------------------------------------------
191. A World of Sin
192. A Writer and Three Script Editors Walk Into a Bar
193. A walk through Hoogvliet
194. AAA, la película: Sin límite en el tiempo
195. ACT Honour Walk
196. Acacia Walk
197. Ach du lieber Harry
198. Achter de schermen bij 'Harry Potter en de orde van de feniks'
199. Achtung Harry! Augen auf!
200. Action Filmmaking: The Making of Wages of Sin

Navigation:
n - Next page
p - Previous page
g - Go to page number
q - Return to main menu
```
