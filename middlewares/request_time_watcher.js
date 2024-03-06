const fs = require("fs");
const path = require("path");

const log_file_path = path.join(
  __dirname,
  "..",
  "logs",
  "req_time_responses.log"
);

const logs_directory = path.dirname(log_file_path);
if (!fs.existsSync(logs_directory)) {
  fs.mkdirSync(logs_directory, { recursive: true });
}

function request_time_watcher(req, res, next) {
  const start = process.hrtime();
  res.on("finish", () => {
    const duration = process.hrtime(start);
    const duration_in_ms = (duration[0] * 1000 + duration[1] / 1000000).toFixed(
      2
    );
    const size = res.getHeader("content-length");
    const sizeStr = size ? `${size} bytes` : "unknown size";
    const log_message = `${req.method} ${req.url} - ${duration_in_ms}ms - ${sizeStr}\n`;

    fs.appendFile(log_file_path, log_message, (err) => {
      if (err) {
        console.error(err.message);
      }
    });
  });
  next();
}

module.exports = request_time_watcher;
