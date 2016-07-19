import axios from 'axios';
import parse from 'parse-link-header';
import * as config from './config';

export default {
  get: (url, config) => {
    return getInstance().get(getUrl(url), putTokenIn(config));
  },
  post: (url, data, config) => {
    return getInstance().post(getUrl(url), data, putTokenIn(config));
  },
  put: (url, data, config) => {
    return getInstance().put(getUrl(url), data, putTokenIn(config));
  },
  delete: (url, config) => {
    return getInstance().delete(getUrl(url), putTokenIn(config));
  }
};

function getInstance () {
  axios.interceptors.request.use(function (config) {
    config.headers.common
    return config;
  }, function (error) {
    // Do something with request error
    return Promise.reject(error);
  });
  axios.interceptors.response.use(function (response) {
    response.pagination = {
      next: 1,
      previous: 1
    };
    if (response.headers.link) {
      response.pagination = parse(response.headers.link);
    }
    return response;
  }, function (error) {
    return Promise.reject(error);
  });
  return axios;
}

function getUrl (url) {
  return `${config.get('baseUrl')}${url}`;
}

function putTokenIn (config) {
  if (config && config.useToken) {
    const TOKEN = localStorage.getItem('token');
    config.headers = TOKEN ? {
      common: {
        Authorization: `Bearer ${TOKEN}`,
        'Content-Type': 'application/json'
      }
    } : {
      common: {
        'Content-Type': 'application/json'
      }
    }
    delete config.useToken;
  }
  return config;
}