# myclockspace

## Como subir o projeto (3 formas)

### Via Makefile (comandos docker abstraídos)
1. Suba o conteiner do docker
```bash
make up
```

2. Isso deve bastar para subir o projeto que estará disponível em `http://localhost:8080`

3. Deixamos mais alguns comandos úteis que pode conferir em Makefile
```bash
make up #comando para subir conteiner
make down #comando para derrubar o conteiner
make prune #comando para fazer limpeza de todos os conteiners docker (necessário novo build)
make test path=./internal/healthcheck #comando para rodar os testes automatizados (necessário passar o caminho de pastas)
...
make help #comando para ajudar a listar comandos disponíveis com uma breve descrição
```

### Via docker
1. Suba o conteiner do docker
```bash
docker compose up -d
```

2. Isso deve bastar para subir o projeto que estará disponível em `http://localhost:8080`

### Local

1. Instale as dependências (se necessário):
```bash
go mod tidy
```

2. Rode o servidor:
```bash
go run ./cmd/app/main.go
```

O servidor estará disponível em `http://localhost:8080`.

## Como registrar rotas

Para registrar novas rotas, edite o arquivo `internal/server/routes.go`.

Exemplo de como adicionar uma rota:

```go
var register = []Routes{
    func(r *app.Router) { routes.NewHealthcheck().Register(r) },
    func(r *app.Router) { routes.NewMinhaRota().Register(r) }, // Nova rota
}
```

Crie o handler em `internal/routes/` e implemente o método `Register(*app.Router)` para adicionar a rota ao router.

Exemplo de handler:
```go
package routes

import "myclockspace.kevinmartiniano.com/internal/app"

type MinhaRota struct{}

func NewMinhaRota() *MinhaRota { return &MinhaRota{} }

func (h *MinhaRota) Register(r *app.Router) {
    r.Get("/minha-rota", h.handler)
}

func (h *MinhaRota) handler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("ok"))
}
```
