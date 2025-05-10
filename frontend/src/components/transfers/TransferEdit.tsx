import {
  Edit,
  SimpleForm,
  TextInput,
  ReferenceInput,
  SelectInput,
  required,
} from 'react-admin';

const TransferEdit = () => {
  return (
    <Edit>
      <SimpleForm>
        <ReferenceInput source="source_storage_provider_id" reference="storage-providers">
          <SelectInput optionText="name" validate={required()} />
        </ReferenceInput>
        <ReferenceInput source="destination_storage_provider_id" reference="storage-providers">
          <SelectInput optionText="name" validate={required()} />
        </ReferenceInput>
        <TextInput source="source_dir" validate={required()} />
        <TextInput source="destination_dir" validate={required()} />
        <TextInput source="post_transfer_source_dir" validate={required()} />
      </SimpleForm>
    </Edit>
  );
};

export default TransferEdit; 