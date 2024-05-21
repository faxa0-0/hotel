const url = require('url');
const userRoutes = require('./userRoutes/userRoutes');

const routeHandler = (req, res) => {
  const parsedURL = url.parse(req.url, true);
  const path = parsedURL.pathname;

  res.setHeader('Content-Type', 'application/json');

  if (path === '/users' || path.startsWith('/users/')) {
    userRoutes(req, res);
  } else {
    res.writeHead(404);
    res.end(JSON.stringify({ message: 'Route not found' }));
  }
};
module.exports = routeHandler;
