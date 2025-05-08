import {
  List,
  Datagrid,
  TextField,
} from 'react-admin';

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