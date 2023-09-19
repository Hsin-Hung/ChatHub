import * as React from "react";
import {
  Box,
  Typography,
  Paper,
  ButtonGroup,
  IconButton,
  createTheme,
  ThemeProvider,
} from "@mui/material";
import ThumbUpIcon from "@mui/icons-material/ThumbUp";
import ThumbDownIcon from "@mui/icons-material/ThumbDown";

const theme = createTheme({
  typography: {
    fontFamily: ["Nunito Sans", "sans-serif"].join(","),
  },
});

export default function Message({ message, isSelf, onSendJsonMessage }) {
  const handleUpVote = () => {
    onSendJsonMessage({
      id: message.Id,
      content: "",
      sender: "",
      upvotes: 1,
      downvotes: 0,
      timestamp: 0,
    });
  };

  const handleDownVote = () => {
    onSendJsonMessage({
      id: message.Id,
      content: "",
      sender: "",
      upvotes: 0,
      downvotes: 1,
      timestamp: 0,
    });
  };

  return (
    <ThemeProvider theme={theme}>
      <Box
        sx={{
          display: "flex",
          justifyContent: isSelf ? "flex-start" : "flex-end",
          mb: 2,
        }}
      >
        <Box
          sx={{
            display: "flex",
            flexDirection: "column",
            alignItems: "left",
          }}
        >
          <Typography variant="body1">{message.Sender}</Typography>
          <Paper
            variant="outlined"
            sx={{
              p: 1,
              backgroundColor: isSelf ? "primary.light" : "secondary.light",
              borderRadius: isSelf
                ? "20px 20px 20px 5px"
                : "20px 20px 5px 20px",
              maxWidth: 400,
            }}
          >
            <Typography variant="body1">{message.Content}</Typography>
          </Paper>
          <ButtonGroup size="small" aria-label="small button group">
            <IconButton
              aria-label="thumb-up"
              size="small"
              color="success"
              onClick={handleUpVote}
            >
              <ThumbUpIcon fontSize="small" />{" "}
              <Typography variant="caption" style={{ marginLeft: 4 }}>
                {message.Upvotes}
              </Typography>
            </IconButton>
            <IconButton
              aria-label="thumb-down"
              size="small"
              color="error"
              onClick={handleDownVote}
            >
              <ThumbDownIcon fontSize="small" />
              <Typography variant="caption" style={{ marginLeft: 4 }}>
                {message.Downvotes}
              </Typography>
            </IconButton>
          </ButtonGroup>
        </Box>
      </Box>
    </ThemeProvider>
  );
}
