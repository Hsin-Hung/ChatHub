import axios from "axios";

const instance = axios.create({
  baseURL: "http://localhost:8080",
  timeout: 1000,
});

export const signUp = async (username, password) => {
  return instance.post(
    "/signup",
    { username: username, password: password },
    { headers: { "Content-Type": "application/json" } }
  );
};

export const signIn = async (username, password) => {
  return instance.post(
    "/signin",
    { username: username, password: password },
    { headers: { "Content-Type": "application/json" } }
  );
};

export const connectChat = async (jwt) => {
  return instance.get("/ws", { params: { token: jwt } });
};
