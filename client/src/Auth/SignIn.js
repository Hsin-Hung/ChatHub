import * as React from "react";
import { useState, useEffect } from "react";
import {
  Button,
  Alert,
  CssBaseline,
  TextField,
  Grid,
  Box,
  Typography,
  Container,
  createTheme,
  ThemeProvider,
} from "@mui/material";
import { Link, useNavigate } from "react-router-dom";
import { signIn, authUser } from "../api/api";

const defaultTheme = createTheme();

export default function SignIn() {
  const navigate = useNavigate();

  // automatically signin if user has already signed in with valid token
  useEffect(() => {
    const authUserStatus = async () => {
      const token = localStorage.getItem("token");
      const username = localStorage.getItem("username");
      if (token != null && username != null) {
        try {
          const res = await authUser(username, token);
          if (res.status === 200) {
            navigate("/home");
          }
        } catch (err) {
          console.log(err);
        }
      }
    };

    authUserStatus().catch(console.error);
  }, []);

  const [showAlert, setShowAlert] = useState({
    status: "",
    value: "",
  });
  const [buttonDisabled, setButtonDisabled] = useState(false);

  // handle sign in logic
  const handleSubmit = async (event) => {
    event.preventDefault();
    setButtonDisabled(true);
    const data = new FormData(event.currentTarget);
    const username = data.get("username");
    const password = data.get("password");

    try {
      const res = await signIn(username, password);
      localStorage.setItem("username", username);
      localStorage.setItem("token", res.data.token);
      navigate("/home");
    } catch (err) {
      let errResponse = err.message;
      if (err.hasOwnProperty("response")) {
        errResponse = err.response.data.error;
      }
      setShowAlert({
        status: "error",
        value: `Failed to sign in: ${errResponse}`,
      });
    } finally {
      setButtonDisabled(false);
    }
  };

  return (
    <ThemeProvider theme={defaultTheme}>
      <Container component="main" maxWidth="xs">
        <CssBaseline />
        <Box
          sx={{
            marginTop: 8,
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
          }}
        >
          <Typography component="h1" variant="h5">
            Sign in
          </Typography>
          <Box
            component="form"
            onSubmit={handleSubmit}
            noValidate
            sx={{ mt: 1 }}
          >
            <TextField
              margin="normal"
              required
              fullWidth
              id="username"
              label="Username"
              name="username"
              autoComplete="username"
              autoFocus
            />
            <TextField
              margin="normal"
              required
              fullWidth
              name="password"
              label="Password"
              type="password"
              id="password"
              autoComplete="current-password"
            />
            <Button
              type="submit"
              fullWidth
              variant="contained"
              sx={{ mt: 3, mb: 2 }}
              disabled={buttonDisabled}
            >
              Sign In
            </Button>
            <Grid
              container
              rowSpacing={1}
              justifyContent="center"
              direction="column"
            >
              <Grid item>
                <Link to="/signup">Don't have an account? Sign Up"</Link>
              </Grid>
              {showAlert.status && (
                <Grid item>
                  <Alert severity={showAlert.status}>{showAlert.value}</Alert>
                </Grid>
              )}
            </Grid>
          </Box>
        </Box>
      </Container>
    </ThemeProvider>
  );
}
