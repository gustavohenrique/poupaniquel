import * as utils from '../utils';

export const INITIAL_STATE = {
  filter: {
    startDate: utils.date().subtract(3, 'months').format('YYYY-MM-DD'),
    endDate: utils.date().format('YYYY-MM-DD'),
    type: 'expense',
    tags: ''
  },
  pie: {
    labels: [],
    datasets: [{
      data: [],
      backgroundColor: [
        '#FF6384',
        '#36A2EB'
      ],
      hoverBackgroundColor: [
        '#FF6384',
        '#36A2EB'
      ]
    }]
  },
  line: {
    labels: [],
    datasets: [{
      label: '',
      data: [],
      fill: false,
      lineTension: 0.1,
      backgroundColor: 'rgba(75,192,192,0.4)',
      borderColor: 'rgba(75,192,192,1)',
      borderCapStyle: 'butt',
      borderDash: [],
      borderDashOffset: 0.0,
      borderJoinStyle: 'miter',
      pointBorderColor: 'rgba(75,192,192,1)',
      pointBackgroundColor: '#fff',
      pointBorderWidth: 1,
      pointHoverRadius: 5,
      pointHoverBackgroundColor: 'rgba(75,192,192,1)',
      pointHoverBorderColor: 'rgba(220,220,220,1)',
      pointHoverBorderWidth: 2,
      pointRadius: 1,
      pointHitRadius: 10
    }]
  }
};

const MODULE_NAME = 'REPORTS';

export default function (state = INITIAL_STATE, action) {
  switch (action.type) {

  case `${MODULE_NAME}_BY_TAGS_SUCCESS`: {
    const pie = Object.assign({}, INITIAL_STATE.pie);
    pie.labels = action.payload.pie.labels;
    pie.datasets[0].data = action.payload.pie.data;

    const line = Object.assign({}, INITIAL_STATE.line);
    line.labels = action.payload.line.map(item => {
      return item.month;
    });
    line.datasets[0].label = action.payload.tags;
    line.datasets[0].data = action.payload.line.map(item => {
      return item.amount;
    });

    return Object.assign({}, state, {
      pie,
      line
    });
  }

  default:
    return state;
  }
}
