import React from 'react';
import Navbar from './Navbar';
import Loading from './Loading';
import Menu from './Menu';

export default class App extends React.Component {
  render() {
    return (
      <div>
        <Loading />
        <Menu path={this.props.children.props.route.path} />
        <Navbar path={this.props.children.props.location} />
        <div className="container">
          {this.props.children}
        </div>
      </div>
    );
  }
}
