const db = require("./db.js");
function setup_database() {
  db.run(
    `CREATE TABLE IF NOT EXISTS articles (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    published DATE DEFAULT CURRENT_TIMESTAMP,
    is_visible BOOLEAN DEFAULT FALSE
) `,
    (err) => {
      if (err) {
        console.error(err.message);
      } else {
        console.log("Table articles created successfully");
      }
    }
  );

  db.run(
    `CREATE TABLE IF NOT EXISTS users (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      name TEXT NOT NULL,
      email TEXT NOT NULL,
      password TEXT NOT NULL
    )`,
    (err) => {
      if (err) {
        console.error("Error creating users table:", err.message);
      } else {
        console.log("Table users created successfully");
      }
    }
  );
}

function initialize_database() {
  console.log("Initializing database...");
  setup_database();
}

module.exports = initialize_database;
