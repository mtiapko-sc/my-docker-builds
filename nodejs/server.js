const express = require('express');
const app = express();

app.get('/healthz', (req, res) => res.send('ok'));
app.get('/', (req, res) => res.send('Hello from Node.js!'));

app.listen(8080, () => console.log('Server running on port 8080'));
