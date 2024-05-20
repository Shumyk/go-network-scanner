import React from "react";

import "./App.css";
import Inputs from "./Inputs";
import {
  ThemeProvider,
  createTheme,
} from "@mui/material";
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
  return (
    <div className="App">
      <ThemeProvider theme={theme}>
        <Inputs />
      </ThemeProvider>
    </div>
  );
}

export default App;
