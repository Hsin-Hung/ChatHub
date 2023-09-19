import * as React from "react";
import { useState, useEffect, useReducer } from "react";
import { Box, TextField, Button, Grid } from "@mui/material";
import SendIcon from "@mui/icons-material/Send";
import Message from "./Message";
import useWebSocket, { ReadyState } from "react-use-websocket";
import { createMessageTemplate } from "../utils/message";
import { CHAT_SERVER_WS_URL } from "../utils/constants";

function reducer(state, action) {
  let newState;
  let newMessage = action.payload;
  switch (action.type) {
    case "new_msg":
      newState = [
        state[0].concat(newMessage),
        { ...state[1], [newMessage.Id]: state[0].length },
      ];
      break;
    case "update_msg":
      let updatedMessageHistory = [...state[0]];
      updatedMessageHistory[state[1][newMessage.Id]] = newMessage;
      newState = [updatedMessageHistory, { ...state[1] }];
      break;
    default:
      throw new Error();
  }
  return newState;
}

export default function Chat() {
  const username = localStorage.getItem("username");
  const token = localStorage.getItem("token");

  const [socketUrl, setSocketUrl] = useState(
    `${CHAT_SERVER_WS_URL}?token=${token}&username=${username}`
  );
  const [chatState, dispatch] = useReducer(reducer, [[], {}]);
  const { sendJsonMessage, lastJsonMessage, readyState } = useWebSocket(
    socketUrl,
    {
      onOpen: () => console.log("opened"),
      //Will attempt to reconnect on all close events, such as server shutting down
      shouldReconnect: (closeEvent) => true,
    }
  );

  const _ = {
    [ReadyState.CONNECTING]: "Connecting",
    [ReadyState.OPEN]: "Open",
    [ReadyState.CLOSING]: "Closing",
    [ReadyState.CLOSED]: "Closed",
    [ReadyState.UNINSTANTIATED]: "Uninstantiated",
  }[readyState];

  useEffect(() => {
    if (lastJsonMessage !== null) {
      if (lastJsonMessage.Id in chatState[1]) {
        dispatch({ type: "update_msg", payload: lastJsonMessage });
        return;
      }
      dispatch({ type: "new_msg", payload: lastJsonMessage });
    }
  }, [lastJsonMessage]);

  const [input, setInput] = React.useState("");

  const handleSend = () => {
    if (input.trim() !== "") {
      setInput("");
      let jsonMessage = createMessageTemplate();
      jsonMessage.content = input;
      jsonMessage.sender = username;
      sendJsonMessage(jsonMessage);
    }
  };

  const handleInputChange = (event) => {
    setInput(event.target.value);
  };

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
        {chatState[0].map((message, idx) => (
          <Message
            key={idx}
            message={message}
            username={username}
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
