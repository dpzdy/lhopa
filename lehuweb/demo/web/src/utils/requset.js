import axios from "axios";

const qs = require("qs");
const request = {
  async get(url) {
    let res = await axios.get(url);
    return res;
  },
  async post(url, data) {
    try {
      let res = await axios.post(url, qs.stringify(data));
      res = res.data;
      return new Promise((resolve, reject) => {
        if (res.code === 0) {
          resolve(res);
        } else {
          reject(res);
        }
      });
    } catch (err) {
      // return (e.message)
      alert("服务器出错");
      console.log(err);
    }
  },
};
export { request };
