import React from 'react';
import * as utils from '../../utils';

export default class Card extends React.Component {

  edit (item) {
    this.props.actions.goTo(`/transactions/edit/${item.id}`);
  }

  remove (index, item, event) {
    this.props.actions.remove(index, item);
    event.stopPropagation();
  }

  render () {
    const { item, index } = this.props;
    const ChildrenIconComponent = item.children && item.children.length > 0
      ? (<div className="float-right">
          <i className="fa fa-th" style={{fontSize: "2em"}}></i>
        </div>)
      : null;
    const dueDate = utils.date(item.dueDate).format('YYYY-MM-DD');
    const tags = item.tags ? item.tags : [];

    return (
      <div className="column col-xs-12 col-sm-6 col-md-3">
        <div onClick={this.edit.bind(this, item)} className="card material" >
          <div className="card-header">
            <div className="card-meta">
              {dueDate}
            </div>
            <h4 className="card-title">{item.description}</h4>
            <h5 className="transaction amount">R$ {item.amount}</h5>
          </div>
          <div className="card-body">
            {tags.map((tag, index) => {
              return (<span className="label" key={`tag_${item.id}_${index}`} style={{marginRight: "5px", backgroundColor: tag === "" ? "white" : "#efefef"}}>{tag}</span>);
            })}
          </div>
          <div className="card-footer">
            <button onClick={this.remove.bind(this, index, item)} className="btn">Remove</button>
            {ChildrenIconComponent}
          </div>
        </div>
      </div>
    );
  }
}

