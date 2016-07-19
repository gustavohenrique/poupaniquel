import React from 'react';
import { connect } from 'react-redux';

export class Alert extends React.Component {

  render () {
    const css = {
      error: 'toast-danger',
      success: 'toast-success'
    };
    const { showIf, message, fail, className, type } = this.props;
    const text = message ? message : fail.error.data.message;
    const cls = ' toast margin-b-10 ' + (css[type] || '');
    const classes = className ? className + cls : cls;
    const component = showIf ? (<div className={classes}>{text}</div>) : null;

    return (<div>{component}</div>);
  }
}

const mapStateToProps = state => {
  return {
    fail: state.common.fail
  };
};

export default connect(mapStateToProps)(Alert);