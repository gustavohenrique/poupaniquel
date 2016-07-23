import React from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { reduxForm /*reset*/ } from 'redux-form';
import Alert from '../../components/Alert.jsx';

export class Form extends React.Component {

  save (data) {
    this.props.save(data);
    this.props.resetForm('transactionForm');
  }

  render () {
    const {
      lastInsertedId,
      fail,
      cancel,
      handleSubmit,
      submitting,
      fields: { id, description, amount, dueDate, tags, type }
    } = this.props;

    return (
      <div>
        <form onSubmit={handleSubmit(this.save.bind(this))} className="margin-t-10">
          <Alert
            type="error"
            showIf={fail.type === 'TRANSACTION_SAVE_FAIL'} />
          <Alert
            type="success"
            showIf={fail.type !== 'TRANSACTION_SAVE_FAIL' && lastInsertedId > 0}
            message={`The transaction #${lastInsertedId} was successfully saved.`} />
          <input {...id} type="hidden" />
          <div className="form-group">
              <label className="form-label">Date</label>
              <input {...dueDate} type="date" maxlength="10" className="form-input input-lg" />
          </div>
          <div className="form-group">
            <label className="form-label">Type</label>
            <select {...type} className="form-select select-lg">
              <option value="expense">Expense</option>
              <option value="income">Income</option>
            </select>
          </div>
          <div className="form-group">
              <label className="form-label">Description</label>
              <input {...description} type="text" maxlength="250" placeholder="Describe in the maximum of 250 characters" className="form-input input-lg" />
          </div>
          <div className="form-group">
              <label className="form-label">Amount</label>
              <input {...amount} type="text" maxlength="7" className={amount.touched && amount.error ? "form-input input-lg is-danger" : "form-input input-lg"} />
              {amount.touched && amount.error && <div className="validation error">{amount.error}</div>}
          </div>
          <div className="form-group">
              <label className="form-label">Tags</label>
              <input {...tags} type="text" placeholder="Use comma (,) as a separator. Example: car,home,it services" className="form-input input-lg" />
          </div>
          <div className="form-group buttons btn-group">
            <button type="submit" className="btn btn-primary btn-lg">Save</button>
            <button type="button" onClick={cancel} className="btn btn-lg">Cancel</button>
          </div>
        </form>
      </div>
    );
  }
}

const mapStateToProps = state => {
  return {
    lastInsertedId: state.transactions.lastInsertedId,
    fail: state.common.fail
  };
};

const validate = values => {
  const errors = {};
  if (! values.amount) {
    errors.amount = 'Required.';
  }
  else if (! /^[0-9]+(\.[0-9]{1,2})?$/.test(values.amount)) {
    errors.amount = 'Use . as decimal separator. Example: 25.99';
  }
  else if (parseFloat(values.amount) < 0) {
    errors.amount = 'The amount must be > 0.';
  }
  return errors;
};

export default reduxForm({
  form: 'transactionForm',
  fields: ['id', 'dueDate', 'description', 'amount', 'tags', 'type'],
  validate
}, mapStateToProps)(Form);
