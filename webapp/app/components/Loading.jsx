import React from 'react';
import { connect } from 'react-redux';

export class Loading extends React.Component {

  render () {
    const cls = this.props.inProgress ? 'active' : '';

    return (
      <div className={`ui tiny ${cls} progress`}>
        <div className="bar"></div>
      </div>
    );
  }
}

const mapStateToProps = state => {
  return {
    inProgress: state.common.inProgress
  };
};

export default connect(mapStateToProps)(Loading);



