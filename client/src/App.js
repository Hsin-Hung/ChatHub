import * as React from "react";
import SignIn from "./Auth/SignIn";
import SignUp from "./Auth/SignUp";
import ChatRoom from "./Chat/ChatRoom";
import Home from "./Home/Home";
import { Routes, Route } from "react-router-dom";

function App() {
  return (
    <Routes>
      <Route path="/" element={<SignIn />} />
      <Route path="/signin" element={<SignIn />} />
      <Route path="/signup" element={<SignUp />} />
      <Route path="/home" element={<Home />} />
      <Route path="/chatroom" element={<ChatRoom />} />
    </Routes>
  );
}

export default App;
