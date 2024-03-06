const db = require("../data/db");

function create_article(title, content) {
  return new Promise((resolve, reject) => {
    db.run(
      "INSERT INTO articles (title, content) VALUES (?,?)",
      [title, content],
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

function get_article_by_id(id) {
  return new Promise((resolve, reject) => {
    db.get("SELECT * FROM articles WHERE id = ?", [id], (err, result) => {
      if (err) {
        console.error(err.message);
        reject(err);
      } else {
        resolve(result);
      }
    });
  });
}

function get_all_articles() {
  return new Promise((resolve, reject) => {
    db.all("SELECT * FROM articles", [], (err, result) => {
      if (err) {
        console.error(err.message);
      } else {
        resolve(result);
      }
    });
  });
}

module.exports = { create_article, get_article_by_id, get_all_articles };
