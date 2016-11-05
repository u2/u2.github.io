// Issues prints a table of GitHub issues matching the search terms.
package main

import (
    "fmt"
    "log"
    "os"

    "./github"
    "time"
)

func main() {
    result, err := github.SearchIssues(os.Args[1:])
    if err != nil {
        log.Fatal(err)
    }

    issues := make(map[int]int)

    fmt.Printf("%d issues:\n", result.TotalCount)
    for _, item := range result.Items {
        if item.CreatedAt.After(time.Now().AddDate(0,0,-30)) {
            fmt.Printf("In One Month  %4d-%2d-%2d #%-5d %9.9s %.55s\n",
                item.CreatedAt.Year(), item.CreatedAt.Month(), item.CreatedAt.Day(), item.Number, item.User.Login, item.Title)
            issues[1]++
        } else if item.CreatedAt.After(time.Now().AddDate(-1,0,0)) {
            fmt.Printf("In One Year   %4d-%2d-%2d #%-5d %9.9s %.55s\n",
                item.CreatedAt.Year(), item.CreatedAt.Month(), item.CreatedAt.Day(), item.Number, item.User.Login, item.Title)
            issues[2]++
        } else {
            fmt.Printf("More one Year %4d-%2d-%2d #%-5d %9.9s %.55s\n",
                item.CreatedAt.Year(), item.CreatedAt.Month(), item.CreatedAt.Day(), item.Number, item.User.Login, item.Title)
            issues[3]++
        }
    }

    fmt.Printf("%d %d %d", issues[1], issues[2], issues[3])
}
