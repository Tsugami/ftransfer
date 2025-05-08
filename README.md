# FTransfer

Sistema de transferência de arquivos entre diversos protocolos.

## Tecnologias Utilizadas

- Backend:
  - Go (Golang)
  - PostgreSQL
  - Docker

- Frontend:
  - React
  - React Admin
  - TypeScript
  - Vite

## Requisitos

- Docker e Docker Compose
- Go 1.21 ou superior
- Node.js 18 ou superior

## Comandos Makefile

O projeto utiliza Makefile para facilitar o desenvolvimento. Aqui estão os principais comandos:

### Desenvolvimento
```bash
make dev              # Inicia o servidor com hot-reload
make docker_dev      # Inicia os containers de desenvolvimento
make docker_dev_down # Para os containers de desenvolvimento
make docker_dev_logs # Mostra os logs dos containers de desenvolvimento
make psql_dev        # Conecta ao PostgreSQL de desenvolvimento
```

### Banco de Dados
```bash
make migrate-up      # Executa as migrações para cima
make migrate-down    # Reverte as migrações
make migrate-create  # Cria um novo arquivo de migração (use: make migrate-create name=nome_da_migracao)
```

### Docker
```bash
make test_dockerfile_build    # Constrói a imagem Docker de produção
make test_dockerfile_up       # Inicia os containers de produção
make test_dockerfile_down     # Para os containers de produção
make test_dockerfile_psql     # Conecta ao PostgreSQL de produção
make clean_docker_volumes     # Remove todos os volumes Docker de produção
```

### Outros
```bash
make build           # Compila o projeto
make run            # Executa o projeto
make test           # Executa os testes
make lint           # Executa o linter
make deps           # Instala as dependências
make dev-tools      # Instala as ferramentas de desenvolvimento (migrate, air, etc)
make clean          # Remove arquivos de build
```

## Executando o Projeto

### Desenvolvimento

1. Clone o repositório:
```bash
git clone https://github.com/seu-usuario/ftransfer.git
cd ftransfer
```

2. Inicie o banco de dados:
```bash
make docker_dev
```

4. Inicie o backend:
```bash
make dev
```

5. Em outro terminal, inicie o frontend:
```bash
cd frontend
npm install
npm run dev
```

### Produção

Para executar em ambiente de produção:

```bash
make test_dockerfile_up
```

## Estrutura do Projeto

```
.
├── cmd/               # Ponto de entrada da aplicação
├── internal/         # Regras de negócio
├── pkg/             # Código que pode ser usado por aplicações externas
├── repositories/    # Implementações dos repositórios
├── migrations/      # Arquivos de migração do banco de dados
├── frontend/        # Aplicação React
├── docs/           # Documentação do projeto
└── public/         # Arquivos públicos
```

## Licença

Este projeto está licenciado sob a licença MIT - veja o arquivo [LICENSE](LICENSE) para mais detalhes. 