
# User Authentication with JWT

Este é um template básico para implementar autenticação de usuários utilizando JWT (JSON Web Tokens) com o framework **Gin** em Go (Golang). Ele permite o registro e login de usuários, além de fornecer rotas privadas para acessar e atualizar informações do perfil do usuário de forma segura.

## Funcionalidades

- **Registro de Usuário**: Criação de novos usuários com validação de email único.
- **Login de Usuário**: Geração de token JWT para autenticação.
- **Rotas Privadas**: Rotas protegidas que exigem um token JWT válido para acessar informações do perfil e atualizar os dados do usuário.
- **Armazenamento de Dados**: Uso de banco de dados PostgreSQL para armazenar dados dos usuários.

## Tecnologias Utilizadas

- **Go (Golang)**: Linguagem de programação backend.
- **Gin**: Framework web para Go, utilizado para a criação de APIs RESTful.
- **JWT (JSON Web Token)**: Implementação de autenticação baseada em token.
- **PostgreSQL**: Banco de dados relacional para armazenar dados dos usuários.
- **Gorm**: ORM (Object Relational Mapping) para Go, utilizado para interagir com o banco de dados.

## Instalação

### 1. Clone o repositório

```bash
git clone https://github.com/seuusuario/user-auth-token.git
cd user-auth-token
```

### 2. Instale as dependências

Certifique-se de que o Go está instalado em sua máquina. Você pode instalar as dependências utilizando o seguinte comando:

```bash
go mod tidy
```

### 3. Configure o banco de dados

Altere a string de conexão no arquivo `database/database.go` para corresponder às configurações do seu banco de dados PostgreSQL.

```go
dsn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
```

### 4. Defina a chave secreta do JWT

No arquivo `.env`, defina a chave secreta utilizada para assinar os tokens JWT:

```env
SECRET_KEY=seu_segredo_aqui
```

### 5. Rode o banco de dados

Execute as migrações do banco de dados com o comando:

```bash
go run main.go
```

Isso criará as tabelas necessárias no banco de dados.

## Como Usar

### 1. Criar um Novo Usuário

Para criar um novo usuário, envie uma requisição **POST** para a rota `/register` com o corpo JSON:

```json
{
  "username": "seu_usuario",
  "email": "email@dominio.com",
  "password": "senha_secreta"
}
```

### 2. Login de Usuário

Para fazer login e obter o token JWT, envie uma requisição **POST** para a rota `/login` com o corpo JSON:

```json
{
  "email": "email@dominio.com",
  "password": "senha_secreta"
}
```

A resposta conterá o token JWT:

```json
{
  "token": "seu_token_aqui"
}
```

### 3. Acessar Rotas Privadas

Com o token JWT obtido no login, envie uma requisição **GET** ou **PUT** para as rotas privadas `/profile` ou `/profile` para obter ou atualizar o perfil do usuário. Não se esqueça de incluir o token JWT no cabeçalho da requisição:

```bash
Authorization: Bearer seu_token_aqui
```

Exemplo de requisição para obter o perfil do usuário:

```bash
GET /profile
```

Exemplo de requisição para atualizar o perfil do usuário:

```bash
PUT /profile
Content-Type: application/json
{
  "username": "novo_usuario",
  "avatar": "nova_imagem.jpg"
}
```

## Estrutura do Projeto

```
.
├── controllers/
│   ├── auth_controller.go     # Lógica de autenticação
│   ├── profile_controller.go  # Manipulação do perfil do usuário
├── database/
│   ├── db.go                  # Configuração do banco de dados
├── middlewares/
│   ├── auth.go                # Middleware para verificar a autenticação do token JWT
├── models/
│   ├── users.go               # Definição do modelo de usuário
├── routes/
│   ├── routes.go              # Definição das rotas da API
├── .env                       # Variáveis de ambiente (como a chave secreta do JWT)
├── main.go                    # Arquivo principal que inicializa o servidor
├── go.mod                     # Dependências do Go
├── go.sum                     # Dependências do Go
└── README.md                  # Este arquivo
```

## Contribuindo

Se você tiver alguma sugestão ou melhoria, fique à vontade para abrir uma **issue** ou enviar um **pull request**.

## Licença

Este projeto está licenciado sob a Licença MIT - veja o arquivo [LICENSE](LICENSE) para mais detalhes.