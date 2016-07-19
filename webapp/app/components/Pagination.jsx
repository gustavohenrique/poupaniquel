import React from 'react';

export default class Pagination extends React.Component {
  constructor (props) {
    super(props);
    this.previousPage = this.previousPage.bind(this);
    this.nextPage = this.nextPage.bind(this);
    this.refreshFirstPage = this.refreshFirstPage.bind(this);
  }

  componentWillMount () {
    this._fetchData(this.props.options);
  }

  previousPage () {
    const options = Object.assign({}, this.props.options);
    const previous = options.pagination.previous > 0 ? options.pagination.previous : 0;
    options.pagination.page = previous;
    this._fetchData(options);
  }

  nextPage () {
    const options = Object.assign({}, this.props.options);
    const next = options.pagination.next;
    options.pagination.page = next;
    this._fetchData(options);
  }

  refreshFirstPage () {
    const options = Object.assign({}, this.props.options);
    options.pagination.page = 1;
    this._fetchData(options);
  }

  _fetchData (options) {
    const sort = this.refs.sort ? this.refs.sort.value : this.props.defaultSort;
    const { pagination, userId } = options;
    this.props.actions.fetchAll({ pagination, sort, userId });
  }

  render () {
    const { sort, pagination } = this.props.options;

    return (
      <div>
        <div className="btn-group">
          <a href="javascript:" onClick={this.previousPage} title="Previous" className="btn btn-lg btn-link">
            <i className="fa fa-arrow-left"></i>
          </a>
          <a title="Current page" onClick={this.refreshFirstPage} className="btn btn-lg btn-primary">
            {pagination.current}
          </a>
          <a href="javascript:" onClick={this.nextPage} title="Next" className="btn btn-lg btn-link">
            <i className="fa fa-arrow-right"></i>
          </a>
        </div>
        <div className="inline-flex">
          <div className="text-left">
            <div className="form-group">
              <select defaultValue={sort.selected} ref="sort" className="form-select select-lg">
                <option value="">Sort by</option>
                {sort.options.map((option, index) => {
                  return (<option key={`sort${index}`} value={option.value}>{option.text}</option>);
                })}
              </select>
            </div>
          </div>
        </div>
      </div>
    );
  }
}
