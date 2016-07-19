import * as service from './service';
import { push } from 'react-router-redux';
import jwtDecode from 'jwt-decode';

function showLoadingBar () {
  return {
    type: 'LOADING'
  };
}

export function authenticate (login, password, redirectTo = '/') {
  const type = 'AUTHENTICATE';
  return dispatch => {
    dispatch(showLoadingBar());
    return service.authenticate(login, password)
      .then(response => {
        const decoded = jwtDecode(response.data.token);
        localStorage.setItem('token', response.data.token);
        dispatch({
          type: `${type}_SUCCESS`,
          payload: {
            token: response.data.token,
            user: decoded
          }
        });
        dispatch(push(redirectTo));
      })
      .catch(err => {
        localStorage.removeItem('token');
        dispatch({ type: `${type}_FAIL`, error: err });
      });
  };
}
