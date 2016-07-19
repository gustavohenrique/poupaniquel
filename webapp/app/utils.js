import moment from 'moment';

export function deepSearch (id, data) {
  for (let item of data) {
    if (item.id === id) {
      return item;
    }
    else if (Array.isArray(item.children)) {
      const found = deepSearch(id, item.children);
      if (found) {
        return found;
      }
    };
  }
}

export function today () {
  return new Date();
}

export const date = moment;