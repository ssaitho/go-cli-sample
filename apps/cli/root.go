package cli

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "go-cli-sample",
    Short: "A simple CLI tool",
    Run: func(cmd *cobra.Command, args []string) {
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
