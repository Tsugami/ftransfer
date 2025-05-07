import { Title } from 'react-admin';
import { Card, CardContent, Typography, Box } from '@mui/material';
import { useGetList } from 'react-admin';

export const Dashboard = () => {
  const { total: storageProvidersTotal } = useGetList('storage-providers');
  const { total: transfersTotal } = useGetList('transfers');

  return (
    <>
      <Title title="Dashboard" />
      <Box sx={{ display: 'flex', gap: 2, flexWrap: 'wrap' }}>
        <Box sx={{ flex: '1 1 300px' }}>
          <Card>
            <CardContent>
              <Typography variant="h5" component="div">
                Storage Providers
              </Typography>
              <Typography variant="h3" component="div">
                {storageProvidersTotal || 0}
              </Typography>
            </CardContent>
          </Card>
        </Box>
        <Box sx={{ flex: '1 1 300px' }}>
          <Card>
            <CardContent>
              <Typography variant="h5" component="div">
                Transfers
              </Typography>
              <Typography variant="h3" component="div">
                {transfersTotal || 0}
              </Typography>
            </CardContent>
          </Card>
        </Box>
      </Box>
    </>
  );
}; 