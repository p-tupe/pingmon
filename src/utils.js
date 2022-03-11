import http from "http";
import https from "https";

export function isUP(url = "") {
  if (url.startsWith("https")) return _isUPhttps(url);
  else if (url.startsWith("http")) return _isUPhttp(url);
  else
    return Promise.reject(
      "Invalid url encountered : " + url + "\n url must start with http[s]://"
    );
}

function _isUPhttp(url) {
  return new Promise((resolve, reject) => {
    const req = http.get(
      url,
      ({ statusCode, statusMessage, headers: { location } }) => {
        if (statusCode === 200) return resolve(true);
        if (statusCode === 301 || statusCode === 302) return isUP(location);
        resolve(false);
      }
    );
    req.on("error", reject);
  });
}

function _isUPhttps(url) {
  return new Promise((resolve, reject) => {
    const req = https.get(
      url,
      ({ statusCode, statusMessage, headers: { location } }) => {
        if (statusCode === 200) return resolve(true);
        if (statusCode === 301 || statusCode === 302) return isUP(location);
        resolve(false);
      }
    );
    req.on("error", reject);
  });
}

export function notify(msg = "") {
  const data = JSON.stringify(msg);

  const options = {
    hostname: "hooks.slack.com",
    path: "/service/xxxx",
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      "Content-Length": data.length,
    },
  };

  return new Promise((resolve, reject) => {
    const req = https.request(options, ({ statusCode, statusMessage }) => {
      if (statusCode === 200) resolve();
      else reject(statusMessage);
    });
    req.on("error", reject);
    req.write(data);
    req.end();
  });
}
