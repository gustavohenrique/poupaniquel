import * as service from './service';
import * as utils from '../utils';

const MODULE_NAME = 'REPORTS';

function showLoadingBar () {
  return {
    type: 'LOADING'
  };
}

export function showReports (data) {
  const type = `${MODULE_NAME}_BY_TAGS`;
  return dispatch => {
    dispatch(showLoadingBar());
    const filter = Object.assign({}, data, {
      createdAt: utils.date(data.createdAt).format('YYYY-MM-DD'),
      type: data.type === 'income' ? data.type : 'expense',
      tags: data.tags ? data.tags.replace(', ', ',') : ''
    });
    return service.getReportsByTags(filter)
      .then(response => {
        if (response.data && response.data.length > 0) {
          const sum = response.data.reduce((pre, cur) => {
            return { amount: pre.amount + cur.amount };
          });
          const pie = {
            labels: [filter.tags, 'Total'],
            data: [sum.amount, response.data[0].total]
          };
          dispatch({
            type: `${type}_SUCCESS`,
            payload: {
              pie: pie,
              line: response.data || {},
              tags: filter.tags
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
