import http from "k6/http";
import { check, sleep } from "k6";

// Configuration for spike test
export let options = {
  stages: [
    { duration: "10s", target: 10 }, // Gradually ramp-up to 10 users
    { duration: "10s", target: 100 }, // Spike to 100 users
    { duration: "10s", target: 10 }, // Ramp-down back to 10 users
  ],
};

export default function () {
  let res = http.get("http://localhost:8080");
  check(res, {
    "is status 200": (r) => r.status === 200,
  });
  sleep(1);
}
