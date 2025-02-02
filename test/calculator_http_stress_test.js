import { check, sleep } from "k6";
import http from "k6/http";

export const options = {
  stages: [
    { duration: "1m", target: 200 },
    { duration: "3m", target: 200 },
    { duration: "1m", target: 0 },
  ],
};

export default function () {
  const params = {
    headers: { "Content-Type": "application/json" },
  };

  const body = JSON.stringify({
    a: Math.floor(Math.random() * 100),
    b: Math.floor(Math.random() * 100),
  });

  const res = http.post("http://localhost:8000/v1/calculator", body, params);

  check(res, { "status was 200": (r) => r.status == 200 });

  sleep(0.1);
}
