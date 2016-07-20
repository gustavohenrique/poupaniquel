import React from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { reduxForm } from 'redux-form';
import Chart from 'chart.js';
import * as reportsActions from '../actions';
import Alert from '../../components/Alert.jsx';

export class Tags extends React.Component {

  constructor (props) {
    super(props);
    this.state = {};
  }

  componentDidMount () {
    let pieCanvas = this.refs.pieChart;
    let pieChart = new Chart(pieCanvas, {
      type: 'pie',
      data: this.props.pie
    });

    let lineCanvas = this.refs.lineChart;
    let lineChart = new Chart(lineCanvas, {
      type: 'line',
      data: this.props.line
    });

    this.setState({ pieChart: pieChart, lineChart: lineChart });
  }

  componentDidUpdate () {
    const { pie, line } = this.props;

    if (pie.labels.length > 0) {
      const pieChart = this.state.pieChart;
      pie.datasets.forEach((dataset, i) => pieChart.data.datasets[i].data = dataset.data);
      pieChart.data.labels = pie.labels;
      pieChart.update();
    }

    if (line.labels.length > 0) {
      const lineChart = this.state.lineChart;
      line.datasets.forEach((dataset, i) => lineChart.data.datasets[i].data = dataset.data);
      lineChart.data.labels = line.labels;
      lineChart.update();
    }
  }

  showReports (data) {
    this.props.actions.showReports(data)
  }

  render () {
    const {
      pie,
      line,
      fail,
      handleSubmit,
      submitting,
      fields: { startDate, endDate, type, tag }
    } = this.props;

    const hasData = chartType => {
      return chartType.labels.length > 0;
    };

    return (
      <div>
        <h2>By tag</h2>
        <Alert
            type="error"
            showIf={fail.type === 'REPORTS_BY_TAG_FAIL'} />
        <form onSubmit={handleSubmit(this.showReports.bind(this))} className="margin-t-10">
          <div className="form-group">
            <label className="form-label">From</label>
            <input {...startDate} type="date" className="form-input input-lg" />
          </div>
          <div className="form-group">
            <label className="form-label">Until</label>
            <input {...endDate} type="date" className="form-input input-lg" />
          </div>
          <div className="form-group">
            <label className="form-label">Type</label>
            <select {...type} className="form-select select-lg">
              <option value="expense">Expense</option>
              <option value="income">Income</option>
            </select>
          </div>
          <div className="form-group">
            <label className="form-label">Tag</label>
            <input {...tag} type="text" placeholder="" className="form-input input-lg" />
          </div>
          <div className="form-group buttons btn-group">
            <button type="submit" className="btn btn-primary btn-lg">Show reports</button>
          </div>
        </form>
        <div>
          <div className={hasData(pie) && hasData(line) ? '' : 'hidden'}>
            <canvas  ref={'pieChart'} height={'200'} width={'400'}></canvas>
            <canvas className="margin-t-50" ref={'lineChart'} height={'100'} width={'300'}></canvas>
          </div>
        </div>
      </div>
    );
  }
}

const mapStateToProps = state => {
  return {
    initialValues: state.reports.filter,
    pie: state.reports.pie,
    line: state.reports.line,
    fail: state.common.fail
  };
};

const mapDispatchToProps = dispatch => ({
  actions : bindActionCreators(reportsActions, dispatch)
});


export default reduxForm({
  form: 'reportForm',
  fields: ['startDate', 'endDate', 'type', 'tag']
}, mapStateToProps, mapDispatchToProps)(Tags);