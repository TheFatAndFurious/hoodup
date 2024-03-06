const express = require("express");
const initialize_database = require("./data/init_db");
const articles_routes = require("./routes/articles");
const login_routes = require("./routes/login");
const admin_routes = require("./routes/admin");
const homepage_route = require("./routes/homepage");
const contact_route = require("./routes/contact");
const verify_token = require("./middlewares/auth_middleware");
const initialize_middlewares = require("./middlewares/init");

const app = express();

initialize_middlewares(app);
initialize_database();

app.set("view engine", "ejs");

app.use("/articles", articles_routes);
app.use("/admin", verify_token, admin_routes);
app.use("/login", login_routes);
app.use("/contact", contact_route);
app.use("/", homepage_route);

app.listen(3000);

console.log("Server is running on port 3000");
