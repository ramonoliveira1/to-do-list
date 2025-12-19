# 📝 To-Do List API

Uma API RESTful para gerenciamento de tarefas (To-Do List) construída em Go, seguindo os princípios de Clean Architecture com separação clara de responsabilidades.

## 🚀 Tecnologias

- **Go 1.21+** - Linguagem de programação
- **MySQL 8.0+** - Banco de dados
- **go-sql-driver/mysql** - Driver MySQL para Go
- **godotenv** - Gerenciamento de variáveis de ambiente

## 📋 Funcionalidades

- ✅ Criar novas tarefas
- ✅ Listar todas as tarefas
- ✅ Buscar tarefa por ID
- ✅ Graceful shutdown
- ✅ Connection pooling otimizado
- ✅ Tratamento robusto de erros
- ✅ Validação de dados

## 🏗️ Arquitetura

O projeto segue os princípios de Clean Architecture com a seguinte estrutura:

```
to-do-list/
├── cmd/
│   └── api/
│       └── main.go              # Entry point da aplicação
├── internal/
│   ├── entity/
│   │   └── task.go              # Entidade de domínio
│   ├── handler/
│   │   └── task_handler.go      # Handlers HTTP (controllers)
│   ├── service/
│   │   └── task_service.go      # Lógica de negócio
│   ├── repository/
│   │   └── task_repo.go         # Camada de acesso a dados
│   └── infra/
│       └── db/
│           └── mysql.go         # Configuração do banco de dados
├── .env                         # Variáveis de ambiente
├── go.mod                       # Dependências do projeto
└── go.sum                       # Checksums das dependências
```

### Camadas:

- **Handler**: Recebe requisições HTTP, valida entrada e retorna respostas
- **Service**: Contém a lógica de negócio da aplicação
- **Repository**: Abstração para acesso ao banco de dados
- **Entity**: Modelos de domínio
- **Infrastructure**: Configurações externas (banco de dados, etc)

## 🔧 Pré-requisitos

- Go 1.21 ou superior
- MySQL 8.0 ou superior
- Git

## ⚙️ Instalação

### 1. Clone o repositório

```bash
git clone https://github.com/seu-usuario/to-do-list.git
cd to-do-list
```

### 2. Instale as dependências

```bash
go mod download
```

### 3. Configure o banco de dados

Crie um banco de dados MySQL:

```sql
CREATE DATABASE todo;

USE todo;

CREATE TABLE tasks (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    completed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

### 4. Configure as variáveis de ambiente

Crie um arquivo `.env` na raiz do projeto:

```env
DB_USER=root
DB_PASSWORD=sua_senha
DB_HOST=localhost
DB_PORT=3306
DB_NAME=todo
```

### 5. Execute a aplicação

```bash
go run cmd/api/main.go
```

Ou compile e execute:

```bash
go build -o api cmd/api/main.go
./api
```

O servidor estará rodando em `http://localhost:8080`

## 📡 Endpoints da API

### Criar uma tarefa

```http
POST /tasks
Content-Type: application/json

{
    "title": "Minha tarefa",
    "description": "Descrição da tarefa"
}
```

**Resposta (201 Created):**
```json
Status: 201 Created
```

**Erros possíveis:**
- `400 Bad Request` - JSON inválido ou título vazio
- `500 Internal Server Error` - Erro ao criar tarefa

---

### Listar todas as tarefas

```http
GET /tasks
```

**Resposta (200 OK):**
```json
[
    {
        "id": 1,
        "title": "Minha tarefa",
        "description": "Descrição da tarefa",
        "completed": false
    },
    {
        "id": 2,
        "title": "Outra tarefa",
        "description": "Outra descrição",
        "completed": true
    }
]
```

**Erros possíveis:**
- `500 Internal Server Error` - Erro ao buscar tarefas

---

### Buscar tarefa por ID

```http
GET /tasks/{id}
```

**Resposta (200 OK):**
```json
{
    "id": 1,
    "title": "Minha tarefa",
    "description": "Descrição da tarefa",
    "completed": false
}
```

**Erros possíveis:**
- `400 Bad Request` - ID inválido
- `404 Not Found` - Tarefa não encontrada
- `500 Internal Server Error` - Erro ao buscar tarefa

## 📮 API Collection

Uma coleção de requisições está disponível em `api-collection.json` para facilitar testes com Postman ou Insomnia.

**Importar no Postman:**
1. Abra o Postman
2. Clique em "Import"
3. Selecione o arquivo `api-collection.json`

**Importar no Insomnia:**
1. Abra o Insomnia
2. Clique em "Import/Export" > "Import Data"
3. Selecione o arquivo `api-collection.json`

## 🧪 Exemplos de Uso

### Usando cURL

```bash
# Criar uma tarefa
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"Estudar Go","description":"Aprender sobre goroutines"}'

# Listar todas as tarefas
curl http://localhost:8080/tasks

# Buscar tarefa por ID
curl http://localhost:8080/tasks/1
```

### Usando HTTPie

```bash
# Criar uma tarefa
http POST localhost:8080/tasks title="Estudar Go" description="Aprender sobre goroutines"

# Listar todas as tarefas
http GET localhost:8080/tasks

# Buscar tarefa por ID
http GET localhost:8080/tasks/1
```

## 🔒 Boas Práticas Implementadas

- ✅ **Clean Architecture**: Separação clara de responsabilidades
- ✅ **Dependency Injection**: Facilita testes e manutenção
- ✅ **Error Handling**: Tratamento robusto de erros em todas as camadas
- ✅ **Graceful Shutdown**: Encerramento seguro do servidor
- ✅ **Connection Pooling**: Gerenciamento eficiente de conexões com o banco
- ✅ **Input Validation**: Validação de dados de entrada
- ✅ **HTTP Status Codes**: Uso adequado de códigos de status HTTP
- ✅ **Environment Variables**: Configuração via variáveis de ambiente
- ✅ **Timeouts**: Proteção contra operações longas

## 🛠️ Configurações do Servidor

O servidor possui as seguintes configurações de timeout:

- **ReadTimeout**: 15 segundos
- **WriteTimeout**: 15 segundos
- **IdleTimeout**: 60 segundos

### Connection Pool:

- **MaxOpenConns**: 10 conexões
- **MaxIdleConns**: 5 conexões
- **ConnMaxLifetime**: 1 hora
- **ConnMaxIdleTime**: 10 minutos

## 🧪 Testes

Para executar os testes (quando implementados):

```bash
# Executar todos os testes
go test ./...

# Executar com verbose
go test -v ./...

# Executar com cobertura
go test -cover ./...

# Gerar relatório de cobertura
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## 📦 Build

### Build simples

```bash
go build -o api cmd/api/main.go
```

### Build otimizado para produção

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o api cmd/api/main.go
```

### Usando Docker (exemplo)

```dockerfile
# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o api cmd/api/main.go

# Run stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/api .
COPY --from=builder /app/.env .
EXPOSE 8080
CMD ["./api"]
```

## 🚀 Deploy

### Variáveis de Ambiente para Produção

```env
DB_USER=prod_user
DB_PASSWORD=strong_password_here
DB_HOST=db.example.com
DB_PORT=3306
DB_NAME=todo_prod
```

## 🤝 Contribuindo

Contribuições são bem-vindas! Sinta-se à vontade para:

1. Fazer um fork do projeto
2. Criar uma branch para sua feature (`git checkout -b feature/MinhaFeature`)
3. Commit suas mudanças (`git commit -m 'Adiciona MinhaFeature'`)
4. Push para a branch (`git push origin feature/MinhaFeature`)
5. Abrir um Pull Request

## 📝 Roadmap

### Funcionalidades Futuras

- [ ] Atualizar tarefa (PUT/PATCH)
- [ ] Deletar tarefa (DELETE)
- [ ] Marcar tarefa como completa
- [ ] Filtros e paginação
- [ ] Autenticação JWT
- [ ] Testes unitários e de integração
- [ ] Logging estruturado
- [ ] Métricas (Prometheus)
- [ ] Documentação OpenAPI/Swagger
- [ ] Rate limiting
- [ ] CORS middleware
- [ ] Migrations automáticas

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## 👨‍💻 Autor

Feito com ❤️ por [Seu Nome]

## 📞 Contato

- GitHub: [@seu-usuario](https://github.com/seu-usuario)
- LinkedIn: [Seu Nome](https://linkedin.com/in/seu-perfil)
- Email: seu.email@example.com

## ⭐ Agradecimentos

Agradecimentos especiais à comunidade Go pela excelente documentação e recursos educacionais.

---

**Nota**: Este projeto foi desenvolvido para fins educacionais e demonstração de boas práticas em Go.

