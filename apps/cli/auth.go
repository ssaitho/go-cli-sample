package cli

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "github.com/spf13/cobra"
)

type Credentials struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

var authCmd = &cobra.Command{
    Use:   "auth",
    Short: "Authenticate a user",
    Run: func(cmd *cobra.Command, args []string) {
        name, _ := cmd.Flags().GetString("name")
        email, _ := cmd.Flags().GetString("email")

        creds := Credentials{
            Name:  name,
            Email: email,
        }

        jsonData, err := json.Marshal(creds)
        if err != nil {
            fmt.Println("Error encoding JSON:", err)
            return
        }

        response, err := http.Post("http://localhost:8080/auth", "application/json", bytes.NewBuffer(jsonData))
        if err != nil {
            fmt.Println("Error sending request:", err)
            return
        }
        defer response.Body.Close()

        if response.StatusCode == http.StatusOK {
            body, err := ioutil.ReadAll(response.Body)
            if err != nil {
                fmt.Println("Error reading response:", err)
                return
            }
            token := string(body)
            err = os.WriteFile("auth_token.txt", []byte(token), 0644)
            if err != nil {
                fmt.Println("Error saving token:", err)
                return
            }
            fmt.Println("Authentication successful")
        } else {
            fmt.Println("Authentication failed:", response.Status)
        }
    },
}

func init() {
    authCmd.Flags().StringP("name", "n", "", "Name of the user")
    authCmd.Flags().StringP("email", "e", "", "Email of the user")
    authCmd.MarkFlagRequired("name")
    authCmd.MarkFlagRequired("email")
    rootCmd.AddCommand(authCmd)
}
