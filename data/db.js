const sqlite3 = require("sqlite3").verbose();
const path = require("path");

const db_path = path.resolve(__dirname, "blog.db");

const db = new sqlite3.Database(db_path, (err) => {
  if (err) {
    return console.error(err.message);
  }
  console.log("Connected to the blog database.");
});

module.exports = db;
