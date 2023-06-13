const express = require("express");
const path = require("path");
const app = express();

const PORT = process.env.PORT ?? 3000;

app.use(express.static(path.join(__dirname, "public")));

app.listen(PORT, (err) => {
  if (err) throw new Error(err);
  else console.info("Pingmon listening on port", PORT);
});
