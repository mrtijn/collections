import axios from "axios";
const baseURL = "http://localhost:3000/v1";

const token = window.localStorage.getItem("token");
const expiresIn = window.localStorage.getItem("expiresIn");

const instance = axios.create({
  baseURL,
  timeout: 1000,
  headers: { "x-access-token": token }
});

instance.interceptors.response.use(res => {
  return res.data;
});

export default instance;
