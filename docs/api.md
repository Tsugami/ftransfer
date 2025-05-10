# API Documentation

This document provides detailed information about the FTransfer API endpoints.

## Base URL

All API endpoints are prefixed with `/api/v1`

## Authentication

Currently, the API does not require authentication.

## Endpoints

### Storage Providers

#### List Storage Providers
```http
GET /storage-providers
```

Returns a list of all storage providers.

**Response**
```json
[
  {
    "id": "string",
    "name": "string",
    "file_system": "string",
    "protocol_connection": {
      // Protocol specific connection details
    }
  }
]
```

#### Get Storage Provider
```http
GET /storage-providers/:id
```

Returns a specific storage provider by ID.

**Response**
```json
{
  "id": "string",
  "name": "string",
  "file_system": "string",
  "protocol_connection": {
    // Protocol specific connection details
  }
}
```

#### Create Storage Provider
```http
POST /storage-providers
```

Creates a new storage provider.

**Request Body**
```json
{
  "name": "string",
  "file_system": "string",
  "protocol_connection": {
    // Protocol specific connection details
  }
}
```

**Response**
```json
{
  "id": "string",
  "name": "string",
  "file_system": "string",
  "protocol_connection": {
    // Protocol specific connection details
  }
}
```

#### Update Storage Provider
```http
PUT /storage-providers/:id
```

Updates an existing storage provider.

**Request Body**
```json
{
  "name": "string",
  "file_system": "string",
  "protocol_connection": {
    // Protocol specific connection details
  }
}
```

#### Delete Storage Provider
```http
DELETE /storage-providers/:id
```

Deletes a storage provider.

### Transfers

#### List Transfers
```http
GET /transfers
```

Returns a list of all transfers.

**Response**
```json
[
  {
    "id": "string",
    "source_storage_provider_id": "string",
    "destination_storage_provider_id": "string",
    "source_dir": "string",
    "destination_dir": "string",
    "post_transfer_source_dir": "string"
  }
]
```

#### Get Transfer
```http
GET /transfers/:id
```

Returns a specific transfer by ID.

**Response**
```json
{
  "id": "string",
  "source_storage_provider_id": "string",
  "destination_storage_provider_id": "string",
  "source_dir": "string",
  "destination_dir": "string",
  "post_transfer_source_dir": "string"
}
```

#### Create Transfer
```http
POST /transfers
```

Creates a new transfer.

**Request Body**
```json
{
  "source_storage_provider_id": "string",
  "destination_storage_provider_id": "string",
  "source_dir": "string",
  "destination_dir": "string",
  "post_transfer_source_dir": "string"
}
```

**Response**
```json
{
  "id": "string",
  "source_storage_provider_id": "string",
  "destination_storage_provider_id": "string",
  "source_dir": "string",
  "destination_dir": "string",
  "post_transfer_source_dir": "string"
}
```

#### Update Transfer
```http
PUT /transfers/:id
```

Updates an existing transfer.

**Request Body**
```json
{
  "source_storage_provider_id": "string",
  "destination_storage_provider_id": "string",
  "source_dir": "string",
  "destination_dir": "string",
  "post_transfer_source_dir": "string"
}
```

#### Delete Transfer
```http
DELETE /transfers/:id
```

Deletes a transfer.

### Events

#### List Events
```http
GET /events?transfer_id=string
```

Returns a list of events for a specific transfer.

**Query Parameters**
- `transfer_id` (required): The ID of the transfer to get events for

**Response**
```json
[
  {
    "id": "string",
    "transfer_id": "string",
    "level": "string",
    "message": "string",
    "created_at": "string"
  }
]
```

## HTTP Status Codes

The API uses the following HTTP status codes:

- **200 OK**: The request was successful
- **201 Created**: The resource was successfully created
- **204 No Content**: The request was successful but there is no content to return
- **400 Bad Request**: The request was invalid or missing required fields
- **404 Not Found**: The requested resource was not found
- **409 Conflict**: The request conflicts with the current state of the resource
- **500 Internal Server Error**: An error occurred on the server
- **502 Bad Gateway**: Error communicating with storage providers

## Error Responses

When an error occurs, the API returns a JSON object with an error message:

```json
{
  "error": "Error message description"
}
``` 