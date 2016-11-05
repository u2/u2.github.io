package omdbapi

import (
    "encoding/json"
    "fmt"
    "net/http"
    "net/url"
    "strings"
)

const MoviesURL = "http://www.omdbapi.com"

type MoviesSearchResult struct {
    Items          []*Movie
}

type Movie struct {
    Title     string
    Poster    string
}

// SearchMovies queries the GitHub Movie tracker.
func SearchMovies(terms []string) (*MoviesSearchResult, error) {
    q := url.QueryEscape(strings.Join(terms, " "))
    fmt.Println(q)
    // resp, err := http.Get(MoviesURL + "?Frozen&y=&plot=short&r=json&t=" + q)
    resp, err := http.Get("http://www.omdbapi.com/?t=Frozen&y=&plot=short&r=json")
    if err != nil {
        return nil, err
    }

    // We must close resp.Body on all execution paths.
    // (Chapter 5 presents 'defer', which makes this simpler.)
    if resp.StatusCode != http.StatusOK {
        resp.Body.Close()
        return nil, fmt.Errorf("search query failed: %s", resp.Status)
    }

    var result MoviesSearchResult

    fmt.Println("#%s", resp)
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        resp.Body.Close()
        return nil, err
    }
    resp.Body.Close()
    return &result, nil
}
