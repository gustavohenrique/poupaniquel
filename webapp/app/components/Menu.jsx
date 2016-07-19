import React from 'react';
import { connect } from 'react-redux';
import { Link } from 'react-router';

export class Menu extends React.Component {

  render () {
    const { isMenuVisible } = this.props;
    const width = isMenuVisible ? '250px' : '0';
    return (
      <div className="sidenav" style={{width: width}}>
        <a href="javascript:" onClick={this.props.hideMenu} className="close">&times;</a>
        <Link to="/">Home</Link>
        <Link to="transactions">Transactions</Link>
        <Link to="reports">Reports</Link>
      </div>
    );
  }
}

const mapStateToProps = state => {
  return {
    isAuthenticated: state.users.isAuthenticated,
    isMenuVisible: state.common.isMenuVisible
  };
};

const mapDispatchToProps = dispatch => ({
  hideMenu: () => {
    dispatch({ type: 'TOGGLE_MENU' });
  }
});

export default connect(mapStateToProps, mapDispatchToProps)(Menu);



