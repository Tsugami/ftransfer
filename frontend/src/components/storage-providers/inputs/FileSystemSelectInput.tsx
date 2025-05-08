import { SelectInput } from 'react-admin';

export const FileSystemSelectInput = () => (
  <SelectInput source="file_system" choices={[
    { id: 'UNIX', name: 'UNIX' },
    { id: 'WINDOWS', name: 'WINDOWS' },
  ]} />
); 