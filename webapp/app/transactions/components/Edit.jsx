import React from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import Form from './Form.jsx';
import Children from './Children.jsx';
import * as transactionsActions from '../actions';
import * as utils from '../../utils';

export class Edit extends React.Component {

  componentDidMount () {
    this._setSelectedTransactionById();
  }

  componentDidUpdate () {
    this._setSelectedTransactionById();
  }

  _setSelectedTransactionById () {
    const { transactions, selected } = this.props;
    const id = parseInt(this.props.params.transactionId);
    // const transaction = utils.deepSearch(id, transactions) || { id: 0 };
    // if (transaction.id !== this.props.selected.id) {
    //   this.props.actions.setSelected(transaction);
    // }
    if (id > 0 && id !== selected.id) {
      this.props.actions.fetchOne(id);
    }
  }

  showCreateChildForm () {
    this.props.actions.addChild();
  }

  saveChild (data) {
    const parent = this.props.selected;
    this.props.actions.save(data, parent);
  }

  render () {
    const { selected, actions, child, history } = this.props;
    const hasSelectedTransaction = selected && selected.id > 0;
    const hasSelectedButNotChild = hasSelectedTransaction && ! child;
    const EditComponent = hasSelectedButNotChild
      ? (<div>
          <div className="btn-fixed float-right" style={{bottom: "75px"}}>
            <button onClick={this.showCreateChildForm.bind(this)} className="btn btn-primary btn-circle">
              <i className="fa fa-plus"></i>
            </button>
          </div>
          <div className="btn-fixed float-right">
            <button onClick={history.goBack} type="button" className="btn btn-secondary btn-circle">
              <i className="fa fa-undo"></i>
            </button>
          </div>
          <Children transactions={selected.children} actions={actions} />
        </div>)
      : null;

    const FormComponent = hasSelectedButNotChild
      ? (<Form initialValues={selected} save={actions.save} cancel={history.goBack} />)
      : (<Form initialValues={child} save={this.saveChild.bind(this)} cancel={actions.cancelChild} />);

    const Component = hasSelectedTransaction
      ? (<div>
          <h2>{child ? `Add in ${selected.description}` : selected.description}</h2>
          {FormComponent}
          {EditComponent}
        </div>)
      : null; //(<h2>No result for this ID</h2>);

    return (
      <div>{Component}</div>
    );
  }
}

const mapStateToProps = state => {
  return {
    transactions: state.transactions.list,
    selected: state.transactions.selected,
    child: state.transactions.child
  };
};

const mapDispatchToProps = dispatch => ({
  actions : bindActionCreators(transactionsActions, dispatch)
});

export default connect(mapStateToProps, mapDispatchToProps)(Edit);
