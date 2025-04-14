```mermaid
flowchart TD
    A[Usuário cadastra um conector (FTP, S3...)] --> B[Folders são registradas nesse conector]
    B --> C[Usuário cria uma Transfer entre pasta A e B]
    C --> D[Motor de execução inicia a transferência]
    D --> E[Arquivo é transferido para o destino]
    E --> F[Arquivo é movido para post_transfer_source_path]
```