import {
  List,
  Datagrid,
  TextField,
  DateField,
  ReferenceField,
  EditButton,
  DeleteButton,
} from 'react-admin';

const TransferList = () => {
  return (
    <List>
      <Datagrid>
        <TextField source="id" />
        <ReferenceField source="source_storage_provider_id" reference="storage-providers" />
        <ReferenceField source="destination_storage_provider_id" reference="storage-providers" />
        <TextField source="source_dir" />
        <TextField source="destination_dir" />
        <TextField source="post_transfer_source_dir" />
        <DateField source="created_at" />
        <DateField source="updated_at" />
        <EditButton />
        <DeleteButton />
      </Datagrid>
    </List>
  );
};

export default TransferList; 