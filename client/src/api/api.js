import axios from "axios";

const api_instance = axios.create({
  baseURL: "http://localhost:8080",
  timeout: 1000,
});

const chat_instance = axios.create({
  baseURL: "http://localhost:8081",
  timeout: 1000,
});

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

export const authUser = async (token) => {
  return api_instance.get("/user", { params: { token: token } });
};

export const getChatRoomInfo = async (jwt) => {
  return chat_instance.get("/info", { params: { token: jwt } });
};
