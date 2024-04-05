/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/Marushio/go_viper_cobra/models"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"encoding/json"
	"bytes"
	"net/http"
)


// createUserCmd represents the createUser command
var createUserCmd = &cobra.Command{
	Use:   "createUser",
	Short: "Criar um usuário na API http://localhost:3000/api/v1/createUser",
	Long: `Realizar o post:
POST http://localhost:3000/api/v1/createUser
Content-Type: application/json

{
    "name":"usuario",
    "email":"usuario@teste.com"
}`,
	Run: func(cmd *cobra.Command, args []string) {
		var user models.User
		user.Name = viper.GetString("name")
		user.Email = viper.GetString("email")
		url := "http://localhost:3000/api/v1/createUser"
		// Codificando os dados do usuário para o formato JSON
		dadosJson, err := json.Marshal(user)
		if err != nil {
			fmt.Printf("Erro ao codificar dados JSON: %v\n", err)
			return
		}

		resp, err := http.Post(url, "application/json", bytes.NewBuffer(dadosJson))
		if err != nil {
			fmt.Printf("Erro ao fazer a requisição HTTP: %v\n", err)
			return
		}

		defer resp.Body.Close()

		// Verificando o status da resposta
		if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated {
			fmt.Printf("Novo usuario:\nNome: %s\nEmail: %s \nCadastrado com sucesso! \n", user.Name, user.Email)
		} else {
			fmt.Printf("Erro ao cadastrar usuário. Código de status: %d\n", resp.StatusCode)
		}
	},
}

func init() {
	
	rootCmd.PersistentFlags().String("name", "", "Nome a ser exibido")
	viper.BindPFlag("name", rootCmd.PersistentFlags().Lookup("name"))
	
	rootCmd.PersistentFlags().String("email", "", "Nome a ser exibido")
	viper.BindPFlag("email", rootCmd.PersistentFlags().Lookup("email"))

	rootCmd.AddCommand(createUserCmd)
}
