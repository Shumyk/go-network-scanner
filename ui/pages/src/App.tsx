import React from "react";

import "./App.css";
import Inputs from "./Inputs";
import { ThemeProvider, createTheme } from "@mui/material";
import { pink } from "@mui/material/colors";
import ScanGridTable from "./ScanGridTable";

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
  const [configured, setConfigured] = React.useState(true);
  const handleConfigured = () => setConfigured(true);

  return (
    <div className="App">
      <ThemeProvider theme={theme}>
        {!configured && <Inputs onConfigured={handleConfigured} />}
        {configured && <ScanGridTable />}
      </ThemeProvider>
    </div>
  );
}

export default App;
