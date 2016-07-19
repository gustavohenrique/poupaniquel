import React from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import * as usersActions from '../actions';
import Alert from '../../components/Alert.jsx';

export class Login extends React.Component {

  handleOnClickAuthenticate () {
    const number = this.refs.number.value || '';
    const password = this.refs.password.value || '';
    const redirectTo = this.props.location.query.next || '/offers';
    this.props.actions.authenticate(number, password, redirectTo);
  }

  render () {
    const { fail, inProgress } = this.props;
    return (
      <div className="columns">
        <div className="column col-12">
          <form className="form-horizontal">
            <Alert type="error" showIf={fail.type === 'AUTHENTICATE_FAIL'} />
            <div className="form-group">
              <input ref="number" className="form-input" type="number" placeholder="Number" />
            </div>
            <div className="form-group">
              <input ref="password" className="form-input" type="password" placeholder="Password" />
            </div>
            <div className="form-group">
              <button onClick={this.handleOnClickAuthenticate.bind(this)} disabled={inProgress} className="btn btn-primary" type="button">Enter</button>
            </div>
          </form>
        </div>
      </div>
    );
  }
}

const mapStateToProps = state => {
  return {
    user: state.users.user,
    fail: state.common.fail,
    inProgress: state.common.inProgress
  };
};

const mapDispatchToProps = dispatch => ({
  actions : bindActionCreators(usersActions, dispatch)
});

export default connect(mapStateToProps, mapDispatchToProps)(Login);
