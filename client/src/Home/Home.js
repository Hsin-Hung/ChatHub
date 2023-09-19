import * as React from "react";
import AppBar from "@mui/material/AppBar";
import Grid from "@mui/material/Grid";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import IconButton from "@mui/material/IconButton";
import MenuIcon from "@mui/icons-material/Menu";
import ChatInfo from "./ChatInfo";
import { useNavigate } from "react-router-dom";

export default function Home() {
  const navigate = useNavigate();
  const logout = () => {
    localStorage.removeItem("username");
    localStorage.removeItem("token");
    navigate("/");
  };

  return (
    <Grid container spacing={0} style={{ minHeight: "80vh" }}>
      <Grid item xs={12}>
        <AppBar position="static">
          <Toolbar>
            <IconButton
              size="large"
              edge="start"
              color="inherit"
              aria-label="menu"
              sx={{ mr: 2 }}
            >
              <MenuIcon />
            </IconButton>
            <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
              Simple Chat
            </Typography>
            <Button color="inherit" onClick={logout}>
              Sign Out
            </Button>
          </Toolbar>
        </AppBar>
      </Grid>
      <Grid
        item
        xs={12}
        style={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
          justifyContent: "center",
          flex: 1, // Allow the ChatInfo component to take up available vertical space
        }}
      >
        <ChatInfo />
      </Grid>
    </Grid>
  );
}
