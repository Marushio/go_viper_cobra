/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Marushio/go_viper_cobra/models"
	"github.com/spf13/cobra"
)

type UsersResponse struct {
	Users []models.User `json:"users"`
}

// listUsersCmd represents the listUsers command
var listUsersCmd = &cobra.Command{
	Use:   "listUsers",
	Short: "List all users from the http://localhost:3000/api/v1/users",
	Long:  `List all users from the http://localhost:3000/api/v1/users`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("listUsers called")
		url := "http://localhost:3000/api/v1/users"

		response, err := http.Get(url)
		if err != nil {
			fmt.Printf("Erro ao fazer a requisição HTTP: %v\n", err)
			return
		}

		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("Erro ao ler o corpo da resposta: %v\n", err)
			return
		}

		// Decodificando a resposta JSON para a estrutura intermediária
		var usersResponse UsersResponse
		err = json.Unmarshal(body, &usersResponse)
		if err != nil {
			fmt.Printf("Erro ao decodificar JSON: %v\n", err)
			return
		}

		fmt.Println("Lista de usuários:")
		for _, user := range usersResponse.Users {
			fmt.Println("ID:", user.ID)
			fmt.Println("Nome:", user.Name)
			fmt.Println("Email:", user.Email)
			fmt.Println()
		}
	},
}

func init() {
	rootCmd.AddCommand(listUsersCmd)
}
