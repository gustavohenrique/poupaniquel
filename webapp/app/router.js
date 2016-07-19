import React from 'react';
import { Router, IndexRoute, Route } from 'react-router';
import { history } from './store';

import { loginRequired } from 'users/components/AuthenticationWrapper';
import App from 'components/App';
import Transactions from 'transactions/components/index';
import Reports from 'reports/components/index';
import Login from 'users/components/Login';

export default (
  <Router onUpdate={() => window.scrollTo(0, 0)} history={history}>
    <Route path='/' component={App}>
      <IndexRoute component={Transactions.List} />
      <Route path='/transactions' component={Transactions.List} />
      <Route path='/transactions/create' component={Transactions.Create} />
      <Route path='/transactions/edit/:transactionId' component={Transactions.Edit} />
      <Route path='/reports' component={Reports.Tags} />
      <Route path='/login' component={Login} />
    </Route>
  </Router>
);
