import * as utils from '../utils';
import * as service from './service';
import { push } from 'react-router-redux';
import { reset } from 'redux-form';

const MODULE_NAME = 'TRANSACTION';

function showLoadingBar () {
  return {
    type: 'LOADING'
  };
}

export function setSelected (transaction) {
  return dispatch => {
    dispatch({
      type: `${MODULE_NAME}_SET_SELECTED`,
      payload: transaction
    });
  };
}

export function addChild () {
  return dispatch => {
    dispatch({
      type: `${MODULE_NAME}_ADD_CHILD`
    });
  };
}

export function cancelChild () {
  return dispatch => {
    dispatch({
      type: `${MODULE_NAME}_CANCEL_CHILD`
    });
  };
}

export function unsetSelected () {
  return dispatch => {
    dispatch({
      type: `${MODULE_NAME}_CANCEL_SAVE`
    });
  };
}

export function goTo (url) {
  return dispatch => {
    dispatch(push(url));
  };
}

export function resetForm (form) {
  return dispatch => {
    dispatch(reset(form));
  };
}

export function fetchAll (options) {
  const type = `${MODULE_NAME}_FETCH`;
  return dispatch => {
    dispatch(showLoadingBar());
    return service.fetchAll(options)
      .then(response => {
        const nextPage = parseInt(response.pagination.next.page);
        dispatch({
          type: `${type}_SUCCESS`,
          payload: {
            data: response.data,
            sort: options.sort,
            pagination: {
              next: nextPage,
              previous: parseInt(response.pagination.previous.page),
              current: nextPage - 1,
              page: options.pagination.page,
              perPage: options.pagination.perPage
            }
          }
        });
      })
      .catch(err => {
        dispatch({ type: `${type}_FAIL`, error: err });
      });
  };
}

export function fetchOne (id) {
  const type = `${MODULE_NAME}_FETCH_ONE`;
  return dispatch => {
    dispatch(showLoadingBar());
    return service.fetchOne(id)
      .then(response => {
        dispatch({
          type: `${type}_SUCCESS`,
          payload: response.data
        });
      })
      .catch(err => {
        dispatch({ type: `${type}_FAIL`, error: err });
      });
  };
}

export function remove (index, item) {
  const type = `${MODULE_NAME}_REMOVE`;
  return dispatch => {
    return service.remove(item)
      .then(() => {
        dispatch({
          type: `${type}_SUCCESS`,
          payload: { index, item }
        });
      })
      .catch(err => {
        dispatch({ type: `${type}_FAIL`, error: err });
      });
  };
}

export function save (data, parent = null) {
  const type = `${MODULE_NAME}_SAVE`;
  return dispatch => {
    dispatch(showLoadingBar());
    const parentId = parent && parent.id > 0 ? parent.id : data.parentId;
    const item = Object.assign({}, data, {
      createdAt: utils.date(data.createdAt, utils.date.ISO_8601),
      amount: parseFloat(data.amount),
      tags: data.tags && ! data.tags.push ? data.tags.split(',') : data.tags,
      parentId: parentId
    });
    return service.save(item)
      .then(response => {
        const transaction = parent ? parent : item;
        const payload = data.id > 0 ? Object.assign({}, transaction, {
          id: response.data.id,
          createdAt: utils.date(transaction.createdAt).format('YYYY-MM-DD')
        }) : null;
        dispatch({
          type: `${type}_SUCCESS`,
          payload: payload,
          lastId: response.data.id
        });
      })
      .catch(err => {
        dispatch({ type: `${MODULE_NAME}_SAVE_FAIL`, error: err });
      });
  };
}
