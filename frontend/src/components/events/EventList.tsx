import {
  List,
  Datagrid,
  TextField,
  DateField,
  ReferenceInput,
  SelectInput,
} from 'react-admin';
import { Box, Paper, Typography } from '@mui/material';

export const EventList = () => {
  
  return (
    <List 
    queryOptions={{
      refetchInterval: 1000,
    }}
    filters={[
      <ReferenceInput
        source="transfer_id"
        reference="transfers"
        label="Transferência"
        alwaysOn
      >
        <SelectInput
          optionText={(record) => record.source_dir + " -> " + record.destination_dir}
          sx={{
            color: '#00ff00',
            '& .MuiSelect-select': { color: '#00ff00' },
            '& .MuiOutlinedInput-notchedOutline': { borderColor: '#00ff00' },
            '&:hover .MuiOutlinedInput-notchedOutline': { borderColor: '#00ff00' },
            '&.Mui-focused .MuiOutlinedInput-notchedOutline': { borderColor: '#00ff00' },
          }}
        />
      </ReferenceInput>
    ]}>
      <Box sx={{ bgcolor: '#1e1e1e', minHeight: '100vh', p: 2 }}>
        <Paper
          sx={{
            bgcolor: '#000',
            color: '#00ff00',
            fontFamily: 'monospace',
            p: 2,
            borderRadius: 1,
            boxShadow: '0 0 10px rgba(0,255,0,0.2)',
          }}
        >
          <Typography variant="h6" sx={{ color: '#00ff00', mb: 2, borderBottom: '1px solid #00ff00', pb: 1 }}>
            Last Events Logs
          </Typography>
          <Datagrid bulkActionButtons={false}>
            <DateField
              source="created_at"
              label="Data"
              showTime
              options={{
                year: 'numeric',
                month: '2-digit',
                day: '2-digit',
                hour: '2-digit',
                minute: '2-digit',
                second: '2-digit',
              }}
              sx={{ color: '#00ff00' }}
            />
            <TextField
              source="level"
              label="Nível"
              sx={{ color: '#00ff00' }}
            />
            <TextField
              source="message"
              label="Mensagem"
              sx={{ color: '#00ff00' }}
            />
          </Datagrid>
        </Paper>
      </Box>
    </List>
  );
}; 