import React from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import Pagination from '../../components/Pagination.jsx';
import Alert from '../../components/Alert.jsx';
import Card from './Card.jsx';
import * as transactionsActions from '../actions';

export class List extends React.Component {

  redirectToCreatePage () {
    this.props.actions.goTo('/transactions/create');
  }

  render () {
    const { actions, transactions, removed, sort, pagination, fail } = this.props;
    const removedIndex = removed.index;
    const paginationOptions = {
      sort,
      pagination
    };
    const hasError = fail.type === 'TRANSACTION_FETCH_FAIL';
    return (
      <div>
        <div className="margin-t-10">
          <div className={hasError ? "hidden" : "inline-flex"}>
            <Pagination actions={actions} options={paginationOptions} defaultSort={sort.selected} />
          </div>
          <div className="btn-fixed float-right">
            <button onClick={this.redirectToCreatePage.bind(this)} className="btn btn-primary btn-circle">
              <i className="fa fa-pencil"></i>
            </button>
          </div>
        </div>
        <Alert type="error" showIf={hasError} />
        <Alert showIf={transactions.length <= 0 && ! hasError} message="No items found." className="margin-t-10" />
        <div className="columns cards flex">
          {transactions.map((item, index) => {
            return index !== removedIndex ? (<Card item={item} index={index} actions={actions} key={`transaction${index}`} />) : null;
          })}
        </div>
      </div>
    );
  }
}

const mapStateToProps = state => {
  return {
    transactions: state.transactions.list,
    removed: state.transactions.removed,
    sort: state.transactions.sort,
    pagination: state.transactions.pagination,
    fail: state.common.fail
  };
};

const mapDispatchToProps = dispatch => ({
  actions : bindActionCreators(transactionsActions, dispatch)
});

export default connect(mapStateToProps, mapDispatchToProps)(List);
