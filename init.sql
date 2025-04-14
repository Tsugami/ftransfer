-- Enum para o protocolo
CREATE TYPE protocol AS ENUM ('SFTP', 'FTP', 'S3', 'LOCAL');

-- Enum para o tipo de sistema de arquivos
CREATE TYPE file_system_type AS ENUM ('POSIX', 'WINDOWS');

-- Tabela de conectores
CREATE TABLE connectors (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL UNIQUE,
  description TEXT,
  protocol protocol NOT NULL,
  protocol_connection JSONB NOT NULL,
  file_system file_system_type NOT NULL, -- define comportamento de paths/permissões
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  tags JSONB DEFAULT '[]'
);

-- Tabela de pastas
CREATE TABLE folders (
  id TEXT PRIMARY KEY,
  connector_id TEXT NOT NULL REFERENCES connectors(id),
  directory_path TEXT NOT NULL,
  matches JSONB DEFAULT '[]',
  tags JSONB DEFAULT '[]',
  created_at TIMESTAMP,
  updated_at TIMESTAMP
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
