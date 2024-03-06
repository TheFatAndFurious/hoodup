const jwt = require("jsonwebtoken");

function verify_token(req, res, next) {
  const token = req.cookies.auth_token;
  if (!token) return res.status(401).send("fuck you");
  try {
    const decoded = jwt.verify(token, "secretkey");
    req.user_id = decoded.id;
    next();
  } catch (error) {
    res.status(401).send("fuck you");
  }
}

module.exports = verify_token;
