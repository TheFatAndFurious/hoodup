const user_model = require("../models/users");
const bcrypt = require("bcrypt");
const jwt = require("jsonwebtoken");

async function create_new_user(req, res) {
  const { name, email, password } = req.body;
  if (name || email || password) {
    const salt = bcrypt.genSaltSync(10);
    const hashed_password = bcrypt.hashSync(password, salt);
    try {
      await user_model.create_new_user(name, email, hashed_password);
    } catch (err) {
      return next(err);
    }
    return res.status(201).send("User created successfully");
  } else {
    return res.status(400).send("User name, email and password are required");
  }
}

async function get_user_by_email(req, res) {
  const { email, password } = req.body;
  if (!email || !password) {
    return res.status(400).send("Email and password are required");
  }
  const user = await user_model.get_user_by_email(email);
  if (user.length == 0) {
    return res.status(400).send("Invalid email or password");
  }
  const is_password_valid = bcrypt.compareSync(password, user[0].password);
  if (!is_password_valid) {
    return res.status(400).send("Invalid email or password");
  }
  const token = jwt.sign({ id: user[0].id }, "secretkey", { expiresIn: "1h" });
  res.cookie("auth_token", token, {
    maxAge: 1000 * 60 * 60 * 2,
    httpOnly: true,
  });
  return res.status(200).send("thats waddup");
}

module.exports = { create_new_user, get_user_by_email };
