import client from '../httpClient';

export const importDataFromNubank = data => {
  console.log('postar data', data);
  return client.post('/import/nubank', data, { useToken: true });
};
