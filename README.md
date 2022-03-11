# PINGMON (under construction)

A small, succinct, self-hosted website and job monitoring service, that sends downtime notifications to a slack channel!

## Working

```js
/**
 * Logic for notifications for websites:
 *
 * After every interval, check: is site OK?
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
```

## Credits

Inspired by the excellent [pingzy](https://github.com/Flolagale/pingzy)
