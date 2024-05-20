import React from "react";

import "./App.css";
import {
  Box,
  FormControl,
  FormControlLabel,
  Grid,
  Switch,
  TextField,
  ThemeProvider,
  createTheme,
} from "@mui/material";
import Button from "@mui/material/Button";
import { pink } from "@mui/material/colors";

// TODO: move out this
const theme = createTheme({
  palette: {
    mode: "dark",
    primary: {
      main: pink[200],
    },
  },
});

function App() {
  function focusNext(event: React.KeyboardEvent<HTMLDivElement>): void {
    if (event.key === "Enter") {
      const fields = Array.from(document.querySelectorAll("input")) || [];
      const position = fields.indexOf(event.target as HTMLInputElement);
      fields[position + 1]?.focus();
    }
  }

  function handleSubmit(event: React.FormEvent<HTMLFormElement>): void {
    event.preventDefault();
  }

  return (
    <div className="App">
      <ThemeProvider theme={theme}>
        <FormControl
          component="form"
          autoComplete="off"
          onSubmit={handleSubmit}
        >
          <Box sx={{ flexGrow: 1 }} margin={10}>
            <Grid container spacing={2}>
              {/* Database Name grid */}
              <Grid container item spacing={2} rowSpacing={2}>
                <Grid item xs={2}></Grid>
                <Grid item xs={5}>
                  <TextField
                    id="database-name"
                    label="Database Name"
                    helperText="Name of SQLite database file"
                    variant="outlined"
                    fullWidth
                    required
                    onKeyDown={focusNext}
                  />
                </Grid>
                <Grid item xs={5}></Grid>
              </Grid>

              {/* Session Name grid */}
              <Grid container item spacing={2}>
                <Grid item xs={2}></Grid>
                <Grid item xs={5}>
                  <TextField
                    id="session-name"
                    label="Session Name"
                    variant="outlined"
                    fullWidth
                    required
                    onKeyDown={focusNext}
                  />
                </Grid>
                <Grid item xs={5}></Grid>
              </Grid>

              {/* IP Range 1 grid */}
              <Grid container item spacing={2}>
                <Grid item xs={2}>
                  <FormControlLabel
                    labelPlacement="bottom"
                    control={<Switch id="enable-proxy" />}
                    label="Proxy"
                  />
                </Grid>
                <Grid item xs={5}>
                  <TextField
                    id="ip-range-1"
                    label="IP Range 1"
                    variant="outlined"
                    fullWidth
                    required
                    onKeyDown={focusNext}
                  />
                </Grid>
                <Grid item xs={5}>
                  <TextField
                    id="proxy-1"
                    label="Proxy 1"
                    variant="outlined"
                    fullWidth
                    required
                    hidden={true}
                    onKeyDown={focusNext}
                  />
                </Grid>
              </Grid>

              {/* IP Range 2 grid */}
              <Grid container item spacing={2}>
                <Grid item xs={2}></Grid>
                <Grid item xs={5}>
                  <TextField
                    id="ip-range-2"
                    label="IP Range 2"
                    variant="outlined"
                    fullWidth
                    onKeyDown={focusNext}
                  />
                </Grid>
                <Grid item xs={5}>
                  <TextField
                    id="proxy-2"
                    label="Proxy 2"
                    variant="outlined"
                    fullWidth
                    hidden={true}
                    onKeyDown={focusNext}
                  />
                </Grid>
              </Grid>

              {/* IP Range 3 grid */}
              <Grid container item spacing={2}>
                <Grid item xs={2}></Grid>
                <Grid item xs={5}>
                  <TextField
                    id="ip-range-3"
                    label="IP Range 3"
                    variant="outlined"
                    fullWidth
                  />
                </Grid>
                <Grid item xs={5}>
                  <TextField
                    id="proxy-3"
                    label="Proxy 3"
                    variant="outlined"
                    fullWidth
                    hidden={true}
                    onKeyDown={focusNext}
                  />
                </Grid>
              </Grid>

              {/* Button grid */}
              <Grid container item>
                <Grid item xs={1}></Grid>
                <Grid item xs={7}>
                  <Button variant="outlined" type="submit" fullWidth>
                    Submit
                  </Button>
                </Grid>
                <Grid item xs={4}></Grid>
              </Grid>
            </Grid>
          </Box>
        </FormControl>
      </ThemeProvider>
    </div>
  );
}

export default App;
