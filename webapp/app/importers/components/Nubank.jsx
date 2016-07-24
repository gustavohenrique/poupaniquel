import React from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { reduxForm /*reset*/ } from 'redux-form';
import Alert from '../../components/Alert.jsx';
import * as nubankActions from '../actions';

export class Form extends React.Component {

  importData (data) {
    this.props.actions.importData(data);
    // this.props.resetForm('nubankForm');
  }

  render () {
    const {
      success,
      loading,
      fail,
      cancel,
      handleSubmit,
      submitting,
      fields: { username, password }
    } = this.props;
    return (
      <div>
        <form onSubmit={handleSubmit(this.importData.bind(this))} className="margin-t-10">
          <Alert
            type="error"
            showIf={fail.type === 'NUBANK_IMPORT_FAIL'} />
          <Alert
            type="success"
            showIf={success}
            message={`Transactions from Nubank was imported.`} />
          <div className="form-group">
              <label className="form-label">Username</label>
              <input {...username} type="number" maxLength="11" placeholder="CPF" className="form-input input-lg" />
          </div>
          <div className="form-group">
            <label className="form-label">Password</label>
            <input {...password} type="password" maxLength="250" placeholder="The same password used in Nubank site" className="form-input input-lg" />
          </div>
          <div className="form-group buttons btn-group">
            <button type="submit" className={`btn btn-primary btn-lg ${loading}`}>Import data</button>
          </div>
        </form>
      </div>
    );
  }
}

const mapStateToProps = state => {
  return {
    success: state.importers.success,
    loading: state.importers.loading ? 'loading' : '',
    fail: state.common.fail
  };
};

const mapDispatchToProps = dispatch => ({
  actions : bindActionCreators(nubankActions, dispatch)
});

export default reduxForm({
  form: 'nubankForm',
  fields: ['username', 'password']
}, mapStateToProps, mapDispatchToProps)(Form);
