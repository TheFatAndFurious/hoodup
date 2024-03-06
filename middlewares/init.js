const express = require("express");
const cookieParser = require("cookie-parser");
const request_time_watcher = require("./request_time_watcher");
const path = require("path");
const body_parser = require("body-parser");

function initialize_middlewares(app) {
  app.use(body_parser.urlencoded({ extended: false }));
  app.use(body_parser.json());
  app.use(express.static(path.join(__dirname, "..", "public")));
  app.use(cookieParser());
  app.use(request_time_watcher);
}

module.exports = initialize_middlewares;
