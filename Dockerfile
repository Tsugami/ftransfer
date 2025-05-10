# Build Frontend
FROM node:20-alpine AS frontend-builder

WORKDIR /app/frontend

# Instalar pnpm
RUN corepack enable && corepack prepare pnpm@latest --activate

# Instalar dependências
COPY frontend/package*.json frontend/pnpm-lock.yaml ./
RUN pnpm install --frozen-lockfile

# Copiar código fonte
COPY frontend/ ./

ENV VITE_API_URL=/api/v1

# Build do frontend
RUN pnpm run build

# Build Backend
FROM golang:1.23-alpine AS backend-builder

WORKDIR /app

# Instalar dependências necessárias
RUN apk add --no-cache git

# Copiar arquivos de dependências
COPY go.mod go.sum ./
RUN go mod download

# Copiar código fonte
COPY . .

# Build do backend
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o ftransfer ./cmd/server

# Imagem Final
FROM alpine:3.19

WORKDIR /app

# Instalar dependências necessárias
RUN apk add --no-cache ca-certificates tzdata

# Copiar binário do backend
COPY --from=backend-builder /app/ftransfer .

# Copiar build do frontend para o diretório public
COPY --from=frontend-builder /app/frontend/dist ./public
COPY ./migrations ./migrations

ENV GIN_MODE=release
ENV PUBLIC_DIR=/app/public
ENV MIGRATION_DIR=/app/migrations
# Expor porta
EXPOSE 8080

# Comando para executar a aplicação
CMD ["./ftransfer"]
