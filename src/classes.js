import { isUP } from "./utils.js";

export class Site {
  constructor({ url }) {
    this.url = url;
  }

  async check() {
    try {
      this.isDown = !(await isUP(this.url));
      if (this.isDown) {
        console.log(this.url, "is DOWN!");
      } else {
        console.log(this.url, "is UP!");
      }
    } catch (error) {
      console.error("Error which pinging url,", error);
    }
  }
}
