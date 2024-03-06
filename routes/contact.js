const express = require("express");
const router = express.Router();

router.get("/", function (req, res) {
  res.render("../views/pages/contact.ejs");
});

module.exports = router;
