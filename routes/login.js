const express = require("express");
const router = express.Router();
const user_controller = require("../controllers/users");

router.get("/", (req, res) => {
  res.render("../views/pages/login.ejs");
});

router.post("/", user_controller.get_user_by_email);

router.post("/register", user_controller.create_new_user);

module.exports = router;
