const data = require('../../data');

module.exports = (req, res) => {
  let id = parseInt(req.url.split('/')[2]);

  if (data.deleteUser(id)) {
    res.writeHead(204);
    res.end();
  } else {
    res.writeHead(404);
    res.end(JSON.stringify({ message: 'User not found' }));
  }
};
