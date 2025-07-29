const express = require('express');
const bodyParser = require('body-parser');
const mysql = require('mysql2');

const app = express();
app.use(bodyParser.urlencoded({ extended: true }));

// MySQL connection
const db = mysql.createConnection({
  host: 'mysql',
  user: 'root',
  password: 'rootpass',
  database: 'logindb'
});

db.connect(err => {
  if (err) throw err;
  console.log('Connected to MySQL');
});

// Show login form
app.get('/', (req, res) => {
  res.sendFile(__dirname + '/views/login.html');
});

// Handle login
app.post('/login', (req, res) => {
  const { username, password } = req.body;
  db.query('SELECT * FROM users WHERE username = ? AND password = ?', 
    [username, password], (err, results) => {
    if (err) throw err;
    if (results.length > 0) {
      res.send('Login successful!');
    } else {
      res.send('Invalid credentials');
    }
  });
});

app.listen(3000, () => console.log('App running on port 3000'));
