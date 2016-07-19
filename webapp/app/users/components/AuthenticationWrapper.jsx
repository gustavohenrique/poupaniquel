import React from 'react';
import { connect } from 'react-redux';
import { push } from 'react-router-redux';

export function loginRequired (Component) {

  class AuthenticationWrapper extends React.Component {

    componentWillMount () {
      this.checkAuth(this.props.isAuthenticated);
    }

    componentWillReceiveProps (nextProps) {
      this.checkAuth(nextProps.isAuthenticated);
    }

    checkAuth (isAuthenticated) {
      if (! isAuthenticated) {
        let redirectAfterLogin = this.props.location.pathname;
        this.props.dispatch(push(`/login?next=${redirectAfterLogin}`));
      }
    }

    render () {
      return (
        <div>
          {this.props.isAuthenticated ? <Component {...this.props}/> : null}
        </div>
      );
    }
  }

  const mapStateToProps = state => {
    return {
      isAuthenticated: state.users.isAuthenticated
    };
  };

  return connect(mapStateToProps)(AuthenticationWrapper);

}