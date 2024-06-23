package cli

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "go-cli-sample",
    Short: "A simple CLI tool",
    Run: func(cmd *cobra.Command, args []string) {
        token, err := ioutil.ReadFile("auth_token.txt")
        if err != nil || len(token) == 0 {
            fmt.Println("Error: You must authenticate first. Run 'go-cli-sample auth' command.")
            return
        }

        response, err := http.Get("http://localhost:8080/hello")
        if err != nil {
            fmt.Println("Error:", err)
            return
        }
        defer response.Body.Close()
        body, err := ioutil.ReadAll(response.Body)
        if err != nil {
            fmt.Println("Error:", err)
            return
        }
        fmt.Println(string(body))
    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        return
    }
}
