import client from '../httpClient';

export const getReportsByTag = (filter) => {
  return client.get(`/reports?from=${filter.startDate}&until=${filter.endDate}&type=${filter.type}&tag=${filter.tag}`);
};