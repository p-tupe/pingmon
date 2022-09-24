const dummyData = [
  {
    id: 1,
    site: "priteshtupe.com",
    latest_status: "up",
    lastest_ts: Date.now(),
    recentStatuses: [
      { ts: Date.now() - 1000, status: "up" },
      { ts: Date.now() - 2000, status: "up" },
      { ts: Date.now() - 3000, status: "down" },
      { ts: Date.now() - 4000, status: "down" },
      { ts: Date.now() - 5000, status: "up" },
      { ts: Date.now() - 6000, status: "up" },
      { ts: Date.now() - 7000, status: "up" },
    ],
  },
  {
    id: 2,
    site: "priteshtupe.com",
    latest_status: "up",
    lastest_ts: Date.now(),
    recentStatuses: [
      { ts: Date.now() - 1000, status: "up" },
      { ts: Date.now() - 2000, status: "up" },
      { ts: Date.now() - 3000, status: "down" },
      { ts: Date.now() - 4000, status: "down" },
      { ts: Date.now() - 5000, status: "up" },
      { ts: Date.now() - 6000, status: "up" },
      { ts: Date.now() - 7000, status: "up" },
    ],
  },
  {
    id: 3,
    site: "priteshtupe.com",
    latest_status: "up",
    lastest_ts: Date.now(),
    recentStatuses: [
      { ts: Date.now() - 1000, status: "up" },
      { ts: Date.now() - 2000, status: "up" },
      { ts: Date.now() - 3000, status: "down" },
      { ts: Date.now() - 4000, status: "down" },
      { ts: Date.now() - 5000, status: "up" },
      { ts: Date.now() - 6000, status: "up" },
      { ts: Date.now() - 7000, status: "up" },
    ],
  },
  {
    id: 4,
    site: "priteshtupe.com",
    latest_status: "up",
    lastest_ts: Date.now(),
    recentStatuses: [
      { ts: Date.now() - 1000, status: "up" },
      { ts: Date.now() - 2000, status: "up" },
      { ts: Date.now() - 3000, status: "down" },
      { ts: Date.now() - 4000, status: "down" },
      { ts: Date.now() - 5000, status: "up" },
      { ts: Date.now() - 6000, status: "up" },
      { ts: Date.now() - 7000, status: "up" },
    ],
  },
];

window.addEventListener("DOMContentLoaded", (event) => {
  const liveStatusView = dummyData
    .map((obj) => {
      return ["<div><span>"]
        .concat(
          currentStatus(obj),
          obj.site,
          recentHistory(obj),
          "</span></div>"
        )
        .join("");
    })
    .join("");

  // document.getElementById("live-status").innerHTML = liveStatusView;
});

function currentStatus(o) {
  return "()";
}

function recentHistory(o) {
  return "[] [] [] [] [] []";
}
