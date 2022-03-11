import { Site } from "./classes.js";

export default function main() {
  const s = new Site({ url: "https://www.priteshtupe.com" });
  s.check();
}

main();
