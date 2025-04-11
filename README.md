# Go JSONPlaceholder API

Este é um pequeno projeto em Go que cria uma API HTTP para buscar dados de posts usando o [JSONPlaceholder](https://jsonplaceholder.typicode.com/), um serviço de API fake para testes.

## 🚀 Funcionalidades

- Buscar todos os posts disponíveis
- Buscar um post específico por ID
- Servidor HTTP simples rodando na porta `5000`

## 📁 Estrutura do Projeto

- `go.mod`: Gerenciador de dependências do Go
- `main.go`: Ponto de entrada da aplicação
- `handlers/`
  - `posts.go`: Manipulador HTTP
- `services/`
  - `jsonplaceholder.go`: Integração com API externa
- `models/`
  - `post.go`: Modelo Post

## ▶️ Como rodar o projeto

1. Clone este repositório:
   
      ```bash
   git clone https://github.com/seunome/go-jsonplaceholder-api.git
   cd go-jsonplaceholder-api
3. Execute o projeto:
   
      ```bash
        go run main.go
