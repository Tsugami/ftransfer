import { TextInput } from 'react-admin';

export const S3ProtocolConnectionInputBox = () => (
  <>
    <TextInput source="protocol_connection.access_key_id" label="Access Key ID" />
    <TextInput source="protocol_connection.secret_access_key" label="Secret Access Key" />
    <TextInput source="protocol_connection.region" label="Region" />
    <TextInput source="protocol_connection.bucket" label="Bucket" />
    <TextInput source="protocol_connection.endpoint" label="Endpoint" />
  </>
); 