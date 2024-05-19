const data = require('../../data');

module.exports = (req, res) => {
  let body = '';
  const id = parseInt(req.url.split('/')[2]);

  req.on('data', (chunk) => {
    body += chunk;
  });

  req.on('end', () => {
    const parsedBody = new URLSearchParams(body);
    const updatedData = {};
    parsedBody.forEach((value, key) => {
      updatedData[key] = key === 'age' ? parseInt(value) : value;
    });

    const user = data.updateUser(id, updatedData);

    if (user) {
      res.writeHead(200);
      res.end(JSON.stringify(user));
    } else {
      res.writeHead(404);
      res.end(JSON.stringify({ message: 'User not found' }));
    }
  });
};
