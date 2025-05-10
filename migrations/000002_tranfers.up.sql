CREATE TABLE transfers (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  source_storage_provider_id uuid NOT NULL,
  destination_storage_provider_id uuid NOT NULL,
  source_dir VARCHAR(255) NOT NULL,
  destination_dir VARCHAR(255) NOT NULL,
  post_transfer_source_dir VARCHAR(255) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (source_storage_provider_id) REFERENCES storage_providers(id),
  FOREIGN KEY (destination_storage_provider_id) REFERENCES storage_providers(id)
);