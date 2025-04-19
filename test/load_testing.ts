import http from "k6/http";
import { sleep, check } from "k6";

const file = open(`${__ENV.FILE_UPLOAD}`, "b");

export const options = {
  stages: [
    { duration: "5m", target: 25 },
    { duration: "30m", target: 25 },
    { duration: "5m", target: 0 },
  ],
};

export default function () {
  const data: http.RequestBody = {
    field: "file",
    file: http.file(file, `${__ENV.FILE_NAME}`),
  };

  const res = http.post(`${__ENV.SERVER_ENDPOINT}/upload`, data);

  if (res.status !== 200) {
    console.log(res.body);
  }

  check(res, {
    "Res status is 200": (r) => res.status === 200,
    "Res Content-Type header": (r) =>
      res.headers["Content-Type"] === "application/json",
  });

  sleep(3);
}
