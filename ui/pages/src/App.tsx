import React from "react";

import "./App.css";
import {
  FormControl,
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
    const form = event.target as HTMLFormElement;
    const formData = new FormData(form);
    const data = Object.fromEntries(form.entries());
    console.log(data);
  }

  return (
    <div className="App">
      <ThemeProvider theme={theme}>
        <FormControl
          component="form"
          sx={{
            "& :not(style)": { width: "35ch" },
          }}
          autoComplete="off"
          onSubmit={handleSubmit}
        >
          <TextField
            id="database-name"
            label="Database Name"
            helperText="Name of SQLite database file"
            variant="outlined"
            required
            sx={{ color: 'text.primary' }}
            onKeyDown={focusNext}
          />
          <br />
          <TextField
            id="session-name"
            label="Session Name"
            variant="outlined"
            required
            onKeyDown={focusNext}
          />
          <br />
          <TextField
            id="ip-range-1"
            label="IP Range 1"
            variant="outlined"
            required
            onKeyDown={focusNext}
          />
          <br />
          <TextField
            id="ip-range-2"
            label="IP Range 2"
            variant="outlined"
            onKeyDown={focusNext}
          />
          <br />
          <TextField
            id="ip-range-3"
            label="IP Range 3"
            variant="outlined"
          />
          <br />

          <Button variant="outlined" type="submit">
            Submit
          </Button>
        </FormControl>

      </ThemeProvider>
    </div>
  );
}

export default App;
