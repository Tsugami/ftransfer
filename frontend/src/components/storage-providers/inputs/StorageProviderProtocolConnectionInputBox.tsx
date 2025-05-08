import { useRecordContext, FormDataConsumer } from 'react-admin';
import { SFTPProtocolConnectionInputBox } from './protocol-connections/SFTPProtocolConnectionInputBox';
import { FTPProtocolConnectionInputBox } from './protocol-connections/FTPProtocolConnectionInputBox';
import { S3ProtocolConnectionInputBox } from './protocol-connections/S3ProtocolConnectionInputBox';

export const StorageProviderProtocolConnectionInputBox = () => {
  const record = useRecordContext();
  return <FormDataConsumer>
    {({ formData }) => {
      const protocol = formData?.protocol_connection?.protocol ?? record?.protocol_connection?.protocol;
      switch (protocol) {
        case 'SFTP':
          return <SFTPProtocolConnectionInputBox />;
        case 'FTP':
          return <FTPProtocolConnectionInputBox />;
        case 'S3':
          return <S3ProtocolConnectionInputBox />;
        default:
          return null;
      }
    }}
  </FormDataConsumer>
}; 