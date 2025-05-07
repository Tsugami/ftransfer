import {
  List,
  Datagrid,
  TextField,
  Edit,
  SimpleForm,
  TextInput,
  Create,
  SelectInput,
  useRecordContext,
  FormDataConsumer,
  NumberInput
} from 'react-admin';

const PortInput = () => (
  <NumberInput source="protocol_connection.port" label="Port" />
);

const FileSystemSelectInput = () => (
  <SelectInput source="file_system" choices={[
    { id: 'UNIX', name: 'UNIX' },
    { id: 'WINDOWS', name: 'WINDOWS' },
  ]} />
);

const ProtocolSelectInput = () => (
  <SelectInput source="protocol_connection.protocol"
    label="Protocol"
    choices={[
      { id: 'FTP', name: 'FTP' },
      { id: 'SFTP', name: 'SFTP' },
      { id: 'S3', name: 'S3' },
    ]} />
);

export const StorageProviderList = () => (
  <List>
    <Datagrid rowClick="edit">
      <TextField source="id" />
      <TextField source="name" />
      <TextField source="protocol_connection.protocol"
        label="Protocol"
      />
      <TextField source="file_system" />
    </Datagrid>
  </List>
);


const SFTPProtocolConnectionInputBox = () => (
  <>
    <TextInput source="protocol_connection.host" label="Host" />
    <PortInput />
    <TextInput source="protocol_connection.username" label="Username" />
    <TextInput source="protocol_connection.password" label="Password" />
    <TextInput source="protocol_connection.private_key" label="Private Key" />
    <TextInput source="protocol_connection.key_passphrase" label="Key Passphrase" />
  </>
);

const FTPProtocolConnectionInputBox = () => (
  <>
    <TextInput source="protocol_connection.host" label="Host" />
    <PortInput />
    <TextInput source="protocol_connection.username" label="Username" />
    <TextInput source="protocol_connection.password" label="Password" />
  </>
);

const S3ProtocolConnectionInputBox = () => (
  <>
    <TextInput source="protocol_connection.access_key_id" label="Access Key ID" />
    <TextInput source="protocol_connection.secret_access_key" label="Secret Access Key" />
    <TextInput source="protocol_connection.region" label="Region" />
    <TextInput source="protocol_connection.bucket" label="Bucket" />
    <TextInput source="protocol_connection.endpoint" label="Endpoint" />
  </>
);


const StorageProviderProtocolConnectionInputBox = () => {
  const record = useRecordContext();
  return <FormDataConsumer>
    {({ formData }) => {
      const protocol = formData?.protocol_connection?.protocol ?? record?.protocol_connection?.protocol;
      switch (protocol) {
        case 'SFTP':
          return <SFTPProtocolConnectionInputBox />;
        case 'FTP':
          return <FTPProtocolConnectionInputBox />;
        case 'S3':
          return <S3ProtocolConnectionInputBox />;
        default:
          return null;
      }
    }}
  </FormDataConsumer>
};

export const StorageProviderEdit = () => (
  <Edit>
    <SimpleForm>
      <TextInput source="name" />
      <FileSystemSelectInput />
      <ProtocolSelectInput />
      <StorageProviderProtocolConnectionInputBox />
    </SimpleForm>
  </Edit>
);

export const StorageProviderCreate = () => (
  <Create>
    <SimpleForm>
      <TextInput source="name" />
      <FileSystemSelectInput />
      <ProtocolSelectInput />
      <StorageProviderProtocolConnectionInputBox />
    </SimpleForm>
  </Create>
); 