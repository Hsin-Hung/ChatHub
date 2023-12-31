import axios from "axios";
import { API_SERVER_URL, AXIOS_TIMEOUT } from "../utils/constants";

const api_instance = axios.create({
  baseURL: API_SERVER_URL,
  timeout: AXIOS_TIMEOUT,
});

var chat_instance = null;

export const createWSInstance = async () => {
  const ws_uri = localStorage.getItem("ws_uri");
  chat_instance = axios.create({
    baseURL: `http://${ws_uri}`,
    timeout: AXIOS_TIMEOUT,
  });
};

export const signUp = async (username, password) => {
  return api_instance.post(
    "/signup",
    { username: username, password: password },
    { headers: { "Content-Type": "application/json" } }
  );
};

export const signIn = async (username, password) => {
  return api_instance.post(
    "/signin",
    { username: username, password: password },
    { headers: { "Content-Type": "application/json" } }
  );
};

export const authUser = async (username, token) => {
  return api_instance.get("/user", {
    params: { token: token, username, username },
  });
};

export const getChatRoomInfo = async (username, token) => {
  return chat_instance.get("/info", {
    params: { token: token, username, username },
  });
};
