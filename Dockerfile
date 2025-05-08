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

# Build do frontend
RUN pnpm run build

# Build Backend
FROM golang:1.24-alpine AS backend-builder

WORKDIR /app

# Instalar dependências necessárias
RUN apk add --no-cache git

# Copiar arquivos de dependências
COPY go.mod go.sum ./
RUN go mod download

# Copiar código fonte
COPY . .

# Build do backend
RUN CGO_ENABLED=0 GOOS=linux go build -o ftransfer ./cmd/server

# Imagem Final
FROM alpine:3.19

WORKDIR /app

# Instalar dependências necessárias
RUN apk add --no-cache ca-certificates tzdata

# Copiar binário do backend
COPY --from=backend-builder /app/ftransfer .

# Copiar build do frontend para o diretório public
COPY --from=frontend-builder /app/frontend/dist ./public

ENV PUBLIC_DIR=/app/public
# Expor porta
EXPOSE 8080

# Comando para executar a aplicação
CMD ["./ftransfer"]
