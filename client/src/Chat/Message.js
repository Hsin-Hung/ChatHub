import * as React from "react";
import {
  Box,
  Typography,
  Paper,
  Grid,
  Avatar,
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

export default function Message({ message }) {
  const isUser = message.sender === "Henry";
  return (
    <ThemeProvider theme={theme}>
      <Box
        sx={{
          display: "flex",
          justifyContent: isUser ? "flex-start" : "flex-end",
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
          <Typography variant="body1">{message.sender}</Typography>
          <Paper
            variant="outlined"
            sx={{
              p: 1,
              backgroundColor: isUser ? "primary.light" : "secondary.light",
              borderRadius: isUser
                ? "20px 20px 20px 5px"
                : "20px 20px 5px 20px",
              maxWidth: 400,
            }}
          >
            <Typography variant="body1">{message.text}</Typography>
          </Paper>
          <ButtonGroup size="small" aria-label="small button group">
            <IconButton aria-label="thumb-up" size="small" color="success">
              <ThumbUpIcon fontSize="small" />{" "}
              <Typography variant="caption" style={{ marginLeft: 4 }}>
                {message.upvotes}
              </Typography>
            </IconButton>
            <IconButton aria-label="thumb-down" size="small" color="error">
              <ThumbDownIcon fontSize="small" />
              <Typography variant="caption" style={{ marginLeft: 4 }}>
                {message.downvotes}
              </Typography>
            </IconButton>
          </ButtonGroup>
        </Box>
      </Box>
    </ThemeProvider>
  );
}
