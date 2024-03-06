const article_model = require("../models/articles");

async function create_new_article(req, res, next) {
  try {
    const { article_title, article_content } = req.body;
    if (!article_title || !article_content) {
      return res.status(400).send("Article title and content are required");
    } else {
      await article_model.create_article(article_title, article_content);
      res.status(201).send("Article created successfully");
    }
  } catch (err) {
    return next(err);
  }
}
async function get_all_articles(req, res, next) {
  let articles;
  try {
    articles = await article_model.get_all_articles();
  } catch (err) {
    return next(err);
  }
  res.render("../views/pages/articles.ejs", { articles: articles });
}

async function get_article_by_id(req, res, next) {
  let article;
  try {
    const article_id = req.params.id;
    article = await article_model.get_article_by_id(article_id);
  } catch (err) {
    return next(err);
  }
  if (article) {
    res.render("../views/pages/index.ejs", { article: article });
  } else {
    res.render("../views/pages/404.ejs");
  }
}

module.exports = { create_new_article, get_article_by_id, get_all_articles };
