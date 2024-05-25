import React from "react";

import "./App.css";
import Inputs from "./Inputs";
import { ThemeProvider, createTheme } from "@mui/material";
import { pink } from "@mui/material/colors";
import ScanGridTable from "./table/ScanGridTable";
import { ConfigProvider } from "antd";

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
      {!configured && (
        <ThemeProvider theme={theme}>
          <Inputs onConfigured={handleConfigured} />
        </ThemeProvider>
      )}
      
      {configured && (
        <ConfigProvider
          theme={{
            token: {
              colorTextBase: "#ffffff",
              colorBgBase: "#614b9e",
              colorPrimary: "#eb2f96",
              colorInfo: "#e271af",
            },
            components: {
              Table: {
                borderColor: "#ccc1ec",
              },
            },
          }}
        >
          <ScanGridTable />
        </ConfigProvider>
      )}
    </div>
  );
}

export default App;
