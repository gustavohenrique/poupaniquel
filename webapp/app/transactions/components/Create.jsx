import React from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import Form from './Form.jsx';
import * as transactionsActions from '../actions';

export class Create extends React.Component {

  componentDidMount () {
    this.props.actions.setSelected();
  }

  showEditForm () {
    const { actions, lastInsertedId } = this.props;
    actions.goTo(`/transactions/edit/${lastInsertedId}`);
  }

  render () {
    const { selected, history, actions, lastInsertedId } = this.props;

    const AddChildComponent = lastInsertedId > 0
      ? (<div>
          <div className="btn-fixed float-right">
            <button onClick={this.showEditForm.bind(this)} className="btn btn-primary btn-circle">
              <i className="fa fa-th-list"></i>
            </button>
          </div>
        </div>)
      : null;

    return (
      <div>
        <Form initialValues={selected} save={actions.save} cancel={history.goBack} />
        {AddChildComponent}
      </div>
    );
  }
}

const mapStateToProps = state => {
  return {
    selected: state.transactions.selected,
    lastInsertedId: state.transactions.lastInsertedId
  };
};

const mapDispatchToProps = dispatch => ({
  actions : bindActionCreators(transactionsActions, dispatch)
});

export default connect(mapStateToProps, mapDispatchToProps)(Create);
