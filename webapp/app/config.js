export function get (key) {
  let config = {
    baseUrl: 'http://localhost:7000/api/v1'
  };

  if (process.env.NODE_ENV !== 'production') {
    config = {
      baseUrl: 'http://localhost:7000/api/v1'
    };
  }

  return config[key];
};
