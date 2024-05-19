const url = require('url');
const getAllUsers = require('./getAllUsers');
const createUser = require('./createUser');
const getUserByID = require('./getUserByID');
const updateUser = require('./updateUser');
const deleteUser = require('./deleteUser');

const userRoutes = (req,res) => {
  const parsedURL = url.parse(req.url, true);
  const method = req.method;
  const path = parsedURL.pathname;

  if (path === '/users' && method === 'GET') {
    getAllUsers(req, res);
  } else if (path === '/users' && method === 'POST') {
    createUser(req, res);
  } else if (path.startsWith('/users/') && method === 'GET') {
    getUserByID(req, res);
  } else if (path.startsWith('/users/') && method === 'PUT') {
    updateUser(req, res);
  } else if (path.startsWith('/users/') && method === 'DELETE') {
    deleteUser(req, res);
  } else {
    res.writeHead(404);
    res.end(JSON.stringify({ message: 'Route not found in users' }));
  }
};

module.exports = userRoutes;
