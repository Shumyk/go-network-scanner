import React, { useEffect } from "react";

import "./App.css";
import Inputs from "./input/Inputs";
import { ThemeProvider } from "@mui/material";
import ScanGridTable from "./table/ScanGridTable";
import { ConfigProvider } from "antd";
import themes from "./common/themes";
import Storage from "./common/storage";
import setupEvents from "./common/Events";

function App() {
  const [configured, setConfigured] = React.useState(Storage.isConfigured());

  useEffect(() => setupEvents(), []);

  return (
    <div className="App">
      {!configured && (
        <ThemeProvider theme={themes.mui}>
          <Inputs
            onConfigured={() => {
              setConfigured(true);
              Storage.configured();
            }}
          />
        </ThemeProvider>
      )}

      {configured && (
        <ConfigProvider theme={themes.antd}>
          <ScanGridTable />
        </ConfigProvider>
      )}
    </div>
  );
}

export default App;
