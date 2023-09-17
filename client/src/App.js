import * as React from "react";
import "./App.css";
import SignIn from "./Auth/SignIn";
import SignUp from "./Auth/SignUp";
import Chat from "./Chat/Chat";
import { Routes, Route } from "react-router-dom";

function App() {
  return (
    <Routes>
      <Route path="/signin" element={<SignIn />} />
      <Route path="/signup" element={<SignUp />} />
      <Route path="/chat" element={<Chat />} />
    </Routes>
  );
}

export default App;
