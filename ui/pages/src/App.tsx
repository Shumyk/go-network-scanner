import React from "react";

import "./App.css";
import Inputs from "./Inputs";
import { ThemeProvider, createTheme } from "@mui/material";
import { pink } from "@mui/material/colors";
import ScanGrid from "./ScanGrid";

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
  const [configured, setConfigured] = React.useState(false);

  return (
    <div className="App">
      <ThemeProvider theme={theme}>
        {!configured && <Inputs configured={setConfigured} />}
        {configured && <ScanGrid />}
      </ThemeProvider>
    </div>
  );
}

export default App;
