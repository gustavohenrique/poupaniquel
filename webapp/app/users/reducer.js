export const INITIAL_STATE = {
  user: {},
  token: '',
  isAuthenticated: false
};

export default function (state = INITIAL_STATE, action) {
  switch (action.type) {
  case 'AUTHENTICATE_SUCCESS':
    return Object.assign({}, state, {
      user: action.payload.user,
      token: action.payload.token,
      isAuthenticated: true
    });

  default:
    return state;
  }
}
