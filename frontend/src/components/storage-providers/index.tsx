import {
  List,
  Datagrid,
  TextField,
  Edit,
  SimpleForm,
  TextInput,
  Create,
  ReferenceInput,
  SelectInput,
} from 'react-admin';

export const StorageProviderList = () => (
  <List>
    <Datagrid rowClick="edit">
      <TextField source="id" />
      <TextField source="name" />
      <TextField source="description" />
      <TextField source="protocol" />
      <TextField source="file_system" />
    </Datagrid>
  </List>
);

export const StorageProviderEdit = () => (
  <Edit>
    <SimpleForm>
      <TextInput source="name" />
      <TextInput source="description" multiline rows={3} />
      <TextInput source="protocol" />
      <TextInput source="file_system" />
      <TextInput source="protocol_connection.path" label="Storage Path" />
    </SimpleForm>
  </Edit>
);

export const StorageProviderCreate = () => (
  <Create>
    <SimpleForm>
      <TextInput source="name" />
      <TextInput source="description" multiline rows={3} />
      <TextInput source="protocol" />
      <TextInput source="file_system" />
      <TextInput source="protocol_connection.path" label="Storage Path" />
    </SimpleForm>
  </Create>
); 