import * as React from "react";
import { useState } from "react";
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
import { signUp } from "../api/api";

const defaultTheme = createTheme();

export default function SignUp() {
  const navigate = useNavigate();
  const [showAlert, setShowAlert] = useState({
    status: "",
    value: "",
  });
  const [buttonDisabled, setButtonDisabled] = useState(false);

  // handle sign up logic
  const handleSubmit = async (event) => {
    event.preventDefault();
    setButtonDisabled(true);
    const data = new FormData(event.currentTarget);
    const username = data.get("username");
    const password = data.get("password");
    try {
      const res = await signUp(username, password);
      console.log(res);
      setShowAlert({
        status: "success",
        value: res.data.message,
      });
      navigate("/signin");
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
            Sign up
          </Typography>
          <Box
            component="form"
            noValidate
            onSubmit={handleSubmit}
            sx={{ mt: 3 }}
          >
            <Grid container spacing={2}>
              <Grid item xs={12}>
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
              </Grid>
              <Grid item xs={12}>
                <TextField
                  required
                  fullWidth
                  name="password"
                  label="Password"
                  type="password"
                  id="password"
                  autoComplete="new-password"
                />
              </Grid>
            </Grid>
            <Button
              type="submit"
              fullWidth
              variant="contained"
              sx={{ mt: 3, mb: 2 }}
              disabled={buttonDisabled}
            >
              Sign Up
            </Button>
            <Grid
              container
              rowSpacing={1}
              justifyContent="center"
              direction="column"
            >
              <Grid item>
                <Link to="/signin">Already have an account? Sign in</Link>
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
