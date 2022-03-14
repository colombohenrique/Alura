package main

import (
	"gin-api-rest/database"
	"gin-api-rest/models"
	"gin-api-rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	models.Alunos = []models.Aluno{
		{Nome: "José", CPF: "11111111111", RG: "111111111"},
		{Nome: "João", CPF: "11111111111", RG: "111111111"},
		{Nome: "Luis", CPF: "11111111111", RG: "111111111"},
	}
	routes.HadnleRequest()
}
