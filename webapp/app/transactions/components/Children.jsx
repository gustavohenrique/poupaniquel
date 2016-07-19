import React from 'react';
import { connect } from 'react-redux';
import Card from './Card.jsx';

export class Children extends React.Component {

  render () {
    const { transactions, actions, removed } = this.props;
    const removedIndex = removed.index;
    const Component = transactions && transactions.length > 0
      ? (<div>
          <h3>Content</h3>
          <div className="columns cards flex">
            {transactions.map((item, index) => {
              return index !== removedIndex ? (<Card item={item} index={index} actions={actions} key={`child_${index}`} />) : null;
            })}
          </div>
        </div>)
      : null;

    return (
      <div>
        {Component}
      </div>
    );
  }
}

const mapStateToProps = state => {
  return {
    removed: state.transactions.removed
  };
};

export default connect(mapStateToProps)(Children);