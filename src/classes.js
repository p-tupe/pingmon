import { isUP, notify } from "./utils.js";

/**
 * Logic for notifications:
 *
 * After every interval, check: is site up?
 *
 * No?
 *   Already Notified?
 *     No? -> Notify immediately that site is down
 *     Yes?
 *       Has it been 15 mins since last notification?
 *         Yes? -> Notify again that site is still down
 *         No? -> noop
 *
 * Yes?
 *   Was it previously down?
 *     Yes? -> Notify that site is back up
 *     No? -> noop
 *
 */

const RETRY_NOTIF_INTERVAL = 15 * 60 * 1000;

export class Site {
  constructor({ url }) {
    this.url = url;

    this.lastCheckTs = 0; // Timestamp of last check
    this.totalDownCounter = 0; // Total count of uptime checks failed
    this.totalDownTs = 0; // Total time of site being down

    this.isDown = false; // Is site down?
    this.firstDownTs = 0; // Timestamp of the first time site went down after being up
    this.downCounter = 0; // Count of successive uptime check failed
    this.lastNotifTs = 0; // Timestamp of last sent notification
  }

  _resetDownVals() {
    this.totalDownCounter += this.downCounter;
    this.totalDownTs += Date.now() - this.firstDownTs;

    this.isDown = false;
    this.firstDownTs = 0;
    this.downCounter = 0;
    this.lastNotifTs = 0;
  }

  async check() {
    try {
      this.isDown = !(await isUP(this.url));
      this.lastCheckTs = Date.now();

      if (this.isDown) {
        if (this.downCounter === 0) {
          console.error(this.url, "is DOWN!");
          notify({ text: `Pingmon Alert! ${this.url} is DOWN!` });
          this.lastNotifTs = this.firstDownTs = Date.now();
        } else {
          console.warn("Site still DOWN!");
        }

        if (Date.now() - this.lastNotifTs >= RETRY_NOTIF_INTERVAL) {
          notify({ text: `Pingmon Alert! ${this.url} is *still* DOWN!` });
          this.lastNotifTs = Date.now();
        }

        this.downCounter++;
      } else {
        if (this.isDown) {
          notify({ text: `Pingmon Alert! ${this.url} is now back UP!` });
          this._resetDownVals();
        } else {
          console.info(this.url, "is UP!");
        }
      }
    } catch (error) {
      console.error("Unexpected Error: ", error);
    }
  }
}
