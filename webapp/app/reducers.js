import { combineReducers } from 'redux';
import { routerReducer } from 'react-router-redux';
import { reducer as formReducer } from 'redux-form';

import users from './users/reducer';
import transactions from './transactions/reducer';
import reports from './reports/reducer';


const INITIAL_STATE = {
  fail: {
    type: '',
    error: {
      data: {
        message: ''
      }
    }
  },
  inProgress: false,
  isMenuVisible: false
};

function commonReducer (state = INITIAL_STATE, action) {
  const { type, error } = action;

  if (error) {
    const err = error.data ? error : { data: { message: error.message } };
    return Object.assign({}, state, {
      fail: {
        type: type,
        error: err
      },
      inProgress: false
    });
  }
  else if (type === 'LOADING') {
    return Object.assign({}, state, {
      inProgress: true
    });
  }
  else if (type === 'TOGGLE_MENU') {
    const isVisbile = ! state.isMenuVisible;
    return Object.assign({}, state, {
      isMenuVisible: isVisbile
    });
  }

  return Object.assign({}, INITIAL_STATE, {
    inProgress: false
  });
}

export default combineReducers({
  common: commonReducer,
  routing: routerReducer,
  form: formReducer,
  users,
  transactions,
  reports
});
