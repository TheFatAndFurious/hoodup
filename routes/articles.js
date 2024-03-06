const express = require("express");
const router = express.Router();
const articles_controller = require("../controllers/articles");

router.post("/", articles_controller.create_new_article);

router.get("/:id", articles_controller.get_article_by_id);

router.get("/", articles_controller.get_all_articles);

module.exports = router;
