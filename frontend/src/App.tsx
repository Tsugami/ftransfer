import { Admin, Resource } from 'react-admin';
import { StorageProviderList, StorageProviderCreate, StorageProviderEdit } from './components/storage-providers';
import { dataProvider } from './dataProvider';
import { Dashboard } from './components/dashboard';
import StorageIcon from '@mui/icons-material/Storage';

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
    </Admin>
  );
}

export default App;
