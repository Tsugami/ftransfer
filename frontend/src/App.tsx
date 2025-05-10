import { Admin, Resource } from 'react-admin';
import { TransferList, TransferEdit, TransferCreate, TransferShow } from './components/transfers';
import { StorageProviderList, StorageProviderEdit, StorageProviderCreate } from './components/storage-providers';
import { dataProvider } from './dataProvider';
import StorageIcon from '@mui/icons-material/Storage';

const App = () => {
  return (
    <Admin dataProvider={dataProvider}>
      <Resource
        name="transfers"
        list={TransferList}
        edit={TransferEdit}
        create={TransferCreate}
        show={TransferShow}
      />
      <Resource
        name="storage-providers"
        list={StorageProviderList}
        edit={StorageProviderEdit}
        create={StorageProviderCreate}
        icon={StorageIcon}
      />
    </Admin>
  );
};

export default App;
