import React, { useEffect } from "react";

import "./App.css";
import Inputs from "./input/Inputs";
import { ThemeProvider } from "@mui/material";
import ScanGridTable from "./table/ScanGridTable";
import { ConfigProvider } from "antd";
import themes from "./themes";

function App() {
  const parseConfigured = (): boolean =>
    sessionStorage.getItem("configured") === "true";

  const [configured, setConfigured] = React.useState<boolean>(
    parseConfigured()
  );

  const disableKeyReload = (e: KeyboardEvent) => {
    if (isF5(e) || isCtrlR(e) || isCtrlShiftR(e) || isCmdR(e)) {
      e.preventDefault();
      return;
    }

    function isF5(event: KeyboardEvent): boolean {
      return event.key === "F5";
    }
    function isCtrlR(event: KeyboardEvent): boolean {
      return event.ctrlKey && event.key === "r";
    }
    function isCtrlShiftR(event: KeyboardEvent): boolean {
      return event.ctrlKey && event.shiftKey && event.key === "R";
    }
    function isCmdR(event: KeyboardEvent): boolean {
      return event.metaKey && event.key === "r";
    }
  };

  useEffect(() => {
    window.addEventListener("keydown", disableKeyReload);
    return () => window.removeEventListener("keydown", disableKeyReload);
  }, []);

  return (
    <div className="App">
      {!configured && (
        <ThemeProvider theme={themes.mui}>
          <Inputs
            onConfigured={() => {
              setConfigured(true);
              sessionStorage.setItem("configured", "true");
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
