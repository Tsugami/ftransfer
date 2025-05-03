-- Enum para o protocolo
CREATE TYPE protocol AS ENUM ('SFTP', 'FTP', 'S3', 'LOCAL');

-- Enum para o tipo de sistema de arquivos
CREATE TYPE file_system_type AS ENUM ('POSIX', 'WINDOWS');

-- Tabela de conectores
CREATE TABLE storage_providers (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  protocol TEXT NOT NULL,
  protocol_connection JSONB NOT NULL,
  file_system TEXT NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL
);

-- Tabela de pastas
CREATE TABLE folders (
  id TEXT PRIMARY KEY,
  storage_provider_id TEXT NOT NULL REFERENCES storage_providers(id),
  directory_path TEXT NOT NULL,
  tags TEXT[] NOT NULL DEFAULT '{}',
  created_at TIMESTAMP WITH TIME ZONE NOT NULL,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL
);

-- Tabela de transferências
CREATE TABLE transfers (
  id TEXT PRIMARY KEY,
  source_folder_id TEXT NOT NULL REFERENCES folders(id),
  destination_folder_id TEXT NOT NULL REFERENCES folders(id),
  tags JSONB DEFAULT '[]',
  post_transfer_source_path TEXT NOT NULL, -- path where the file will be moved after the transfer
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);
