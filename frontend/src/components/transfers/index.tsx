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
  DateField,
  BooleanField,
} from 'react-admin';

export const TransferList = () => (
  <List>
    <Datagrid rowClick="edit">
      <TextField source="id" />
      <TextField source="name" />
      <TextField source="description" />
      <TextField source="source_storageProvider_id" />
      <TextField source="destination_storageProvider_id" />
      <TextField source="source_path" />
      <TextField source="destination_path" />
      <DateField source="schedule.time" />
      <BooleanField source="status.active" />
    </Datagrid>
  </List>
);

export const TransferEdit = () => (
  <Edit>
    <SimpleForm>
      <TextInput source="name" />
      <TextInput source="description" multiline rows={3} />
      <ReferenceInput source="source_storageProvider_id" reference="storage-providers">
        <SelectInput optionText="name" />
      </ReferenceInput>
      <ReferenceInput source="destination_storageProvider_id" reference="storage-providers">
        <SelectInput optionText="name" />
      </ReferenceInput>
      <TextInput source="source_path" />
      <TextInput source="destination_path" />
      <TextInput source="schedule.type" />
      <TextInput source="schedule.time" />
      <TextInput source="schedule.timezone" />
    </SimpleForm>
  </Edit>
);

export const TransferCreate = () => (
  <Create>
    <SimpleForm>
      <TextInput source="name" />
      <TextInput source="description" multiline rows={3} />
      <ReferenceInput source="source_storageProvider_id" reference="storage-providers">
        <SelectInput optionText="name" />
      </ReferenceInput>
      <ReferenceInput source="destination_storageProvider_id" reference="storage-providers">
        <SelectInput optionText="name" />
      </ReferenceInput>
      <TextInput source="source_path" />
      <TextInput source="destination_path" />
      <TextInput source="schedule.type" />
      <TextInput source="schedule.time" />
      <TextInput source="schedule.timezone" />
    </SimpleForm>
  </Create>
); 