import { Admin, Resource } from 'react-admin';
import { StorageProviderList, StorageProviderCreate, StorageProviderEdit } from './components/storage-providers';
import { TransferList, TransferCreate, TransferEdit } from './components/transfers';
import { dataProvider } from './dataProvider';
import { Dashboard } from './components/dashboard';
import StorageIcon from '@mui/icons-material/Storage';
import SwapHorizIcon from '@mui/icons-material/SwapHoriz';

function App() {
  return (
    <Admin
      dataProvider={dataProvider}
      dashboard={Dashboard}
    >
      <Resource
        name="storage-providers"
        list={StorageProviderList}
        create={StorageProviderCreate}
        edit={StorageProviderEdit}
        icon={StorageIcon}
      />
      <Resource
        name="transfers"
        list={TransferList}
        create={TransferCreate}
        edit={TransferEdit}
        icon={SwapHorizIcon}
      />
    </Admin>
  );
}

export default App;
