import axios from 'axios';
import * as config from '../config';

export const authenticate = (login, password) => {
  const url = `${config.get('baseUrl')}/auth/authenticate`;
  return axios.post(url, { number: login, password: password });
};