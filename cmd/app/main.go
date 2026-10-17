package main

import (
	"myclockspace.kevinmartiniano.com/internal/app"
	"myclockspace.kevinmartiniano.com/internal/server"
)

func main() {
	// 1. Cria a única instância do roteador que vai viver na aplicação
	router := app.NewRouter()

	server := server.NewServer(router)

	server.Up()

	//server := server.NewServer(router)

	//server.Up()

	// 2. Passa essa instância para os arquivos de rota registrarem seus caminhos nela
	// server.RegisterRoutes(router)

	// 3. Inicia o servidor escutando este roteador específico
	// server.StartServer(router)
}
