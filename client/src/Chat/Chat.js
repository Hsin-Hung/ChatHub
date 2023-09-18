import * as React from "react";
import { useState, useEffect } from "react";
import { Box, TextField, Button, Grid } from "@mui/material";
import SendIcon from "@mui/icons-material/Send";
import Message from "./Message";
import { useLocation } from "react-router-dom";
import useWebSocket, { ReadyState } from "react-use-websocket";

export default function Chat() {
  const location = useLocation();
  const username = location.state.username;
  console.log(location);

  const [socketUrl, setSocketUrl] = useState(
    `ws://localhost:8081/ws?token=${location.state.token}`
  );
  console.log(socketUrl);
  const [messageHistory, setMessageHistory] = useState([]);
  const [messageIndex, setMessageIndex] = useState({});
  const { sendJsonMessage, lastJsonMessage, readyState } = useWebSocket(
    socketUrl,
    {
      onOpen: () => console.log("opened"),
      //Will attempt to reconnect on all close events, such as server shutting down
      shouldReconnect: (closeEvent) => true,
    }
  );

  useEffect(() => {
    if (lastJsonMessage !== null) {
      console.log(messageIndex);
      console.log(lastJsonMessage.Id);
      if (lastJsonMessage.Id in messageIndex) {
        setMessageHistory((prev) => {
          const updatedMessageHistory = [...prev];
          updatedMessageHistory[
            messageIndex[lastJsonMessage.Id]
          ] = lastJsonMessage;
          return updatedMessageHistory;
        });
        return;
      }
      const messageHistoryLength = messageHistory.length;
      setMessageHistory((prev) => prev.concat(lastJsonMessage));
      setMessageIndex((prev) => ({
        ...prev,
        [lastJsonMessage.Id]: messageHistoryLength,
      }));
    }
  }, [lastJsonMessage, setMessageHistory]);
  console.log(messageHistory);
  const [input, setInput] = React.useState("");

  const handleSend = () => {
    if (input.trim() !== "") {
      setInput("");
      sendJsonMessage({
        id: "",
        content: input,
        sender: username,
        upvotes: 0,
        downvotes: 0,
        timestamp: -1,
      });
    }
  };

  const handleInputChange = (event) => {
    setInput(event.target.value);
  };

  const connectionStatus = {
    [ReadyState.CONNECTING]: "Connecting",
    [ReadyState.OPEN]: "Open",
    [ReadyState.CLOSING]: "Closing",
    [ReadyState.CLOSED]: "Closed",
    [ReadyState.UNINSTANTIATED]: "Uninstantiated",
  }[readyState];

  return (
    <Box
      sx={{
        height: "100vh",
        display: "flex",
        flexDirection: "column",
        bgcolor: "grey.200",
      }}
    >
      <Box sx={{ flexGrow: 1, overflow: "auto", p: 2 }}>
        {messageHistory.map((message, idx) => (
          <Message
            key={idx}
            message={message}
            isSelf={message.Sender === username}
            onSendJsonMessage={sendJsonMessage}
          />
        ))}
      </Box>
      <Box sx={{ p: 2, backgroundColor: "background.default" }}>
        <Grid container spacing={2}>
          <Grid item xs={10}>
            <TextField
              fullWidth
              placeholder="Type a message"
              value={input}
              onChange={handleInputChange}
            />
          </Grid>
          <Grid item xs={2}>
            <Button
              fullWidth
              size="large"
              color="primary"
              variant="contained"
              endIcon={<SendIcon />}
              onClick={handleSend}
              disabled={readyState !== ReadyState.OPEN}
            >
              Send
            </Button>
          </Grid>
        </Grid>
      </Box>
    </Box>
  );
}
