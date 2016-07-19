import React from 'react';
import { hashHistory } from 'react-router';
import { createStore, applyMiddleware } from 'redux';
import thunk from 'redux-thunk';
import { syncHistoryWithStore, routerMiddleware } from 'react-router-redux'
import reducers from './reducers';

const middlewares = [thunk, routerMiddleware(hashHistory)];

if (process.env.NODE_ENV === 'development') {
  const createLogger = require('redux-logger');
  const logger = createLogger();
  middlewares.push(logger);
}

export const store = createStore(
  reducers,
  applyMiddleware(...middlewares)
);

export const history = syncHistoryWithStore(hashHistory, store)

if (module.hot) {
  module.hot.accept('./reducers', () => {
    store.replaceReducer(require('./reducers'));
  });
  module.hot.accept();

  module.hot.dispose((data) => {
    data.state = store.getState();
    [].slice.apply(document.querySelector('#app').children).forEach(function(c) { c.remove() });
  });
}

export default store;