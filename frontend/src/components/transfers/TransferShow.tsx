import {
  Show,
  SimpleShowLayout,
  TextField,
  DateField,
  ReferenceField,
} from 'react-admin';

const TransferShow = () => {
  return (
    <Show>
      <SimpleShowLayout>
        <TextField source="id" />
        <ReferenceField source="source_storage_provider_id" reference="storage-providers" />
        <ReferenceField source="destination_storage_provider_id" reference="storage-providers" />
        <TextField source="source_dir" />
        <TextField source="destination_dir" />
        <TextField source="post_transfer_source_dir" />
        <DateField source="created_at" />
        <DateField source="updated_at" />
      </SimpleShowLayout>
    </Show>
  );
};

export default TransferShow; 