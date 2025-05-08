import { SelectInput } from 'react-admin';

export const ProtocolSelectInput = () => (
  <SelectInput source="protocol_connection.protocol"
    label="Protocol"
    choices={[
      { id: 'FTP', name: 'FTP' },
      { id: 'SFTP', name: 'SFTP' },
      { id: 'S3', name: 'S3' },
    ]} />
); 