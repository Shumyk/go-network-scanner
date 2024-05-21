import React, { FC, ReactElement } from "react";
import {
  Box,
  FormControl,
  FormControlLabel,
  Grid,
  Switch,
} from "@mui/material";
import Button from "@mui/material/Button";
import InputProxy from "./InputProxy";
import InputIpRange from "./InputIpRange";
import InputName from "./InputName";

const Inputs: FC<{ onConfigured: () => void }> = ({
  onConfigured,
}): ReactElement => {
  const [databaseName, setDatabaseName] = React.useState("test.db"); // TODO: change to 'scans.db'
  const [sessionName, setSessionName] = React.useState("");
  const [ipRange1, setIpRange1] = React.useState("192.168.0.0/24");
  const [ipRange2, setIpRange2] = React.useState("");
  const [ipRange3, setIpRange3] = React.useState("");
  const [enableProxy, setEnableProxy] = React.useState(false);
  const [proxy1, setProxy1] = React.useState("");
  const [proxy2, setProxy2] = React.useState("");
  const [proxy3, setProxy3] = React.useState("");

  function focusNext(event: React.KeyboardEvent<HTMLDivElement>): void {
    if (event.key === "Enter") {
      const fields = Array.from(document.querySelectorAll("input")) || [];
      let position = fields.indexOf(event.target as HTMLInputElement) + 1;

      while (fields[position]?.disabled && position <= fields.length - 1) {
        position++;
      }
      fields[position]?.focus();
    }
  }

  const toggleProxy = (event: React.ChangeEvent<HTMLInputElement>): void => {
    setEnableProxy(event.target.checked);
  };

  function handleSubmit(event: React.FormEvent<HTMLFormElement>): void {
    event.preventDefault();

    const data = {
      "databaseName": databaseName,
      "sessionName": sessionName,
      "ipRange1": ipRange1,
      "ipRange2": ipRange2,
      "ipRange3": ipRange3,
      "enableProxy": enableProxy,
      "proxy1": proxy1,
      "proxy2": proxy2,
      "proxy3": proxy3,
    };
    console.log(data);
    // @ts-ignore
    window.webui.bootstrap(JSON.stringify(data));

    onConfigured();
  }

  return (
    <FormControl component="form" autoComplete="off" onSubmit={handleSubmit}>
      <Box sx={{ flexGrow: 1 }} margin={10}>
        <Grid container spacing={2}>
          {/* Database Name grid */}
          <Grid container item spacing={2} rowSpacing={2}>
            <Grid item xs={2}></Grid>
            <Grid item xs={5}>
              <InputName
                type="database"
                value={databaseName}
                helperText="Name of SQLite database file"
                onChange={(e) => setDatabaseName(e.target.value)}
                onKeyDown={focusNext}
              />
            </Grid>
            <Grid item xs={5}></Grid>
          </Grid>

          {/* Session Name grid */}
          <Grid container item spacing={2}>
            <Grid item xs={2}></Grid>
            <Grid item xs={5}>
              <InputName
                type="session"
                onChange={(e) => setSessionName(e.target.value)}
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
                control={<Switch id="enable-proxy" onChange={toggleProxy} />}
                label="Proxy"
              />
            </Grid>
            <Grid item xs={5}>
              <InputIpRange
                id={1}
                required={true}
                value={ipRange1}
                onChange={(e) => setIpRange1(e.target.value)}
                onKeyDown={focusNext}
              />
            </Grid>
            <Grid item xs={5}>
              <InputProxy
                id={1}
                visible={enableProxy}
                onChange={(e) => setProxy1(e.target.value)}
                onKeyDown={focusNext}
              />
            </Grid>
          </Grid>

          {/* IP Range 2 grid */}
          <Grid container item spacing={2}>
            <Grid item xs={2}></Grid>
            <Grid item xs={5}>
              <InputIpRange
                id={2}
                onChange={(e) => setIpRange2(e.target.value)}
                onKeyDown={focusNext}
              />
            </Grid>
            <Grid item xs={5}>
              <InputProxy
                id={2}
                visible={enableProxy}
                onChange={(e) => setProxy2(e.target.value)}
                onKeyDown={focusNext}
              />
            </Grid>
          </Grid>

          {/* IP Range 3 grid */}
          <Grid container item spacing={2}>
            <Grid item xs={2}></Grid>
            <Grid item xs={5}>
              <InputIpRange
                id={3}
                onChange={(e) => setIpRange3(e.target.value)}
                onKeyDown={focusNext}
              />
            </Grid>
            <Grid item xs={5}>
              <InputProxy
                id={3}
                visible={enableProxy}
                onChange={(e) => setProxy3(e.target.value)}
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
  );
};

export default Inputs;
