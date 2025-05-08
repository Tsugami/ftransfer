import { Create, SimpleForm, TextInput } from 'react-admin';
import { FileSystemSelectInput } from './inputs/FileSystemSelectInput';
import { ProtocolSelectInput } from './inputs/ProtocolSelectInput';
import { StorageProviderProtocolConnectionInputBox } from './inputs/StorageProviderProtocolConnectionInputBox';

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