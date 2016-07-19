// const today = (new Date).toISOString().substring(0, 10);
import * as utils from '../utils';

export const INITIAL_STATE = {
  list: [],
  sort: {
    options: [{
      text: 'ID (desc)',
      value: '-id'
    }, {
      text: 'Date',
      value: '-createdAt'
    }],
    selected: '-createdAt'
  },
  pagination: {
    page: 1,
    perPage: 12,
    current: 0,
    next: 1,
    previous: 0
  },
  selected: {
    id: 0,
    createdAt: utils.date().format('YYYY-MM-DD'),
    description: '',
    amount: 0.00,
    type: 'expense',
    tags: [],
    children: []
  },
  child: null,
  lastInsertedId: 0,
  removed: {}
};

const MODULE_NAME = 'TRANSACTION';

export default function (state = INITIAL_STATE, action) {
  switch (action.type) {
  case `${MODULE_NAME}_FETCH_SUCCESS`:
    return Object.assign({}, INITIAL_STATE, {
      list: action.payload.data,
      pagination: action.payload.pagination,
      sort: Object.assign({}, state.sort, { selected: action.payload.sort }),
      parents: []
    });

  case `${MODULE_NAME}_FETCH_ONE_SUCCESS`: {
    let newState = Object.assign({}, state, {
      selected: action.payload,
      lastInsertedId: 0
    });
    newState.selected.createdAt = utils.date(newState.selected.createdAt).format('YYYY-MM-DD');
    return newState;
  }

  case `${MODULE_NAME}_CANCEL_SAVE`:
    return Object.assign({}, state, {
      selected: INITIAL_STATE.selected
    });

  case `${MODULE_NAME}_SAVE_SUCCESS`:
    return Object.assign({}, state, {
      lastInsertedId: action.lastId,
      selected: action.payload ? action.payload : INITIAL_STATE.selected,
      child: null
    });

  case `${MODULE_NAME}_REMOVE_SUCCESS`: {
    return Object.assign({}, state, {
      removed: action.payload
    });
  }

  case `${MODULE_NAME}_SET_SELECTED`: {
    let newState = Object.assign({}, state);
    newState.selected = action.payload ? action.payload : INITIAL_STATE.selected;
    newState.selected.createdAt = utils.date(newState.selected.createdAt).format('YYYY-MM-DD');
    return newState;
  }

  case `${MODULE_NAME}_ADD_CHILD`: {
    return Object.assign({}, state, {
      child: INITIAL_STATE.selected,
      lastInsertedId: INITIAL_STATE.lastInsertedId
    });
  }

  case `${MODULE_NAME}_CANCEL_CHILD`: {
    return Object.assign({}, state, {
      child: null
    });
  }

  default:
    return state;
  }
}
