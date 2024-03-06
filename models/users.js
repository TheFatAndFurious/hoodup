const db = require("../data/db");

function create_new_user(name, email, password) {
  return new Promise((resolve, reject) => {
    db.run(
      "INSERT INTO users (name, email, password) VALUES (?,?,?)",
      [name, email, password],
      function (err) {
        if (err) {
          reject(err);
        } else {
          resolve(this.lastID);
        }
      }
    );
  });
}

function get_user_by_email(email) {
  return new Promise((resolve, reject) => {
    db.all("SELECT * FROM users WHERE email =?", [email], (err, result) => {
      if (err) {
        reject(err);
      } else {
        resolve(result);
      }
    });
  });
}

module.exports = { create_new_user, get_user_by_email };
