import * as service from './service';
import * as utils from '../utils';

const MODULE_NAME = 'REPORTS';

function showLoadingBar () {
  return {
    type: 'LOADING'
  };
}

export function showReports (data) {
  const type = `${MODULE_NAME}_BY_TAG`;
  return dispatch => {
    dispatch(showLoadingBar());
    const filter = Object.assign({}, data, {
      createdAt: utils.date(data.createdAt).format('YYYY-MM-DD'),
      type: data.type === 'income' ? data.type : 'expense',
      tag: data.tag
    });
    return service.getReportsByTag(filter)
      .then(response => {
        if (response.data && response.data.length > 0) {
          const sum = response.data.reduce((pre, cur) => {
            return { amount: pre.amount + cur.amount };
          });
          const pie = {
            labels: [filter.tag, 'Total'],
            data: [parseFloat(sum.amount.toFixed(2)), parseFloat(response.data[0].total.toFixed(2))]
          };
          dispatch({
            type: `${type}_SUCCESS`,
            payload: {
              pie: pie,
              line: response.data || {},
              tag: filter.tag
            }
          });
        }
        else {
          throw new Error('No result found.');
        }
      })
      .catch(err => {
        dispatch({ type: `${type}_FAIL`, error: err });
      });
  };
}
