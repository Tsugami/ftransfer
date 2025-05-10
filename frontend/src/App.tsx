import { Admin, Resource } from 'react-admin';
import { TransferList, TransferEdit, TransferCreate, TransferShow } from './components/transfers';
import { StorageProviderList, StorageProviderEdit, StorageProviderCreate } from './components/storage-providers';
import { EventList } from './components/events';
import { dataProvider } from './dataProvider';
import StorageIcon from '@mui/icons-material/Storage';
import EventIcon from '@mui/icons-material/Event';

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
      <Resource
        name="events"
        list={EventList}
        icon={EventIcon}
        options={{ label: 'Events' }}
      />
    </Admin>
  );
};

export default App;
