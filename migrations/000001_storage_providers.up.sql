CREATE TYPE storage_provider_file_system AS ENUM ('UNIX', 'WINDOWS');

CREATE TABLE storage_providers (
  id uuid PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  file_system storage_provider_file_system NOT NULL,
  protocol_connection JSONB NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);