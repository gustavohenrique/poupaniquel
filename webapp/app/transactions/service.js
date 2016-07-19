import client from '../httpClient';

export const fetchAll = (options = {}) => {
  let params = {
    page: options.pagination.page,
    perPage: options.pagination.perPage
  };

  if (options.sort.length > 0) {
    params.sort = options.sort;
  }

  return client.get('/transactions', { params: params });
};

export const fetchOne = id => {
  return client.get(`/transactions/${id}`);
};

export const save = item => {
  if (parseInt(item.id) > 0) {
    return client.put(`/transactions/${item.id}`, item, { useToken: true });
  }
  return client.post('/transactions', item, { useToken: true });
};

export const remove = item => {
  return client.delete(`/transactions/${item.id}`, { useToken: true });
};
