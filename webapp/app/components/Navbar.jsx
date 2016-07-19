import React from 'react';
import { connect } from 'react-redux';

export class Navbar extends React.Component {
  render () {
    const { path } = this.props;
    const Menu = (
      <section className="navbar-section">
        <span onClick={this.props.showMenu} className="hamburger">&#9776;</span>
      </section>
    );

    const titles = {
      '': 'Transactions',
      'transactions': 'Transactions',
      'reports': 'Reports'
    };

    const title = path ? titles[path.pathname.split('/')[1]] : '' || 'Poupaniquel';

    return (
      <header className="navbar bg-grey">
        <section className="navbar-section">
          <a href="/" className="navbar-brand">{title}</a>
        </section>
          {Menu}
      </header>
    );
  }
}

const mapStateToProps = state => {
  return {
    isAuthenticated: state.users.isAuthenticated
  };
};

const mapDispatchToProps = dispatch => ({
  showMenu: () => {
    dispatch({ type: 'TOGGLE_MENU' });
  }
});

export default connect(mapStateToProps, mapDispatchToProps)(Navbar);
