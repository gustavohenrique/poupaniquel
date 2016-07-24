export const INITIAL_STATE = {
  success: false,
  loading: false
};

const MODULE_NAME = 'NUBANK_IMPORT';

export default function (state = INITIAL_STATE, action) {
  switch (action.type) {
  case `${MODULE_NAME}_SUCCESS`:
    return Object.assign({}, INITIAL_STATE, {
      success: true
    });

  case `${MODULE_NAME}_FAIL`: {
    return Object.assign({}, state, {
      success: false
    });
  }

  case `${MODULE_NAME}_LOADING`:
    return Object.assign({}, INITIAL_STATE, {
      loading: true
    });

  default:
    return state;
  }
}
