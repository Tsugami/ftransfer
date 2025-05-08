import { TextInput } from 'react-admin';
import { PortInput } from '../PortInput';

export const SFTPProtocolConnectionInputBox = () => (
  <>
    <TextInput source="protocol_connection.host" label="Host" />
    <PortInput />
    <TextInput source="protocol_connection.username" label="Username" />
    <TextInput source="protocol_connection.password" label="Password" />
    <TextInput source="protocol_connection.private_key" label="Private Key" />
    <TextInput source="protocol_connection.key_passphrase" label="Key Passphrase" />
  </>
); 