import client from '../httpClient';

export const getReportsByTags = (filter) => {
  return client.get(`/reports?from=${filter.startDate}&until=${filter.endDate}&type=${filter.type}&tags=${filter.tags}`);
};