import * as service from './service';
import { push } from 'react-router-redux';
import { reset } from 'redux-form';

const MODULE_NAME = 'NUBANK_IMPORT';


export function resetForm (form) {
  return dispatch => {
    dispatch(reset(form));
  };
}

export function importData (data) {
  return dispatch => {
    dispatch({ type: `${MODULE_NAME}_LOADING` });
    
    const credentials = {
      username: data.username.toString(),
      password: data.password
    };
    return service.importDataFromNubank(credentials)
      .then(response => {
        dispatch({ type: `${MODULE_NAME}_SUCCESS` });
      })
      .catch(err => {
        dispatch({ type: `${MODULE_NAME}_FAIL`, error: err });
      });
  };
}
