import React, { Dispatch, FC, ReactElement, SetStateAction } from "react";
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

const Inputs: FC<{ configured: Dispatch<SetStateAction<boolean>> }> = ({
  configured,
}): ReactElement => {
  const [proxy, setProxy] = React.useState(false);

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

  function handleSubmit(event: React.FormEvent<HTMLFormElement>): void {
    event.preventDefault();
    configured(true);
  }

  const toggleProxy = (event: React.ChangeEvent<HTMLInputElement>): void => {
    setProxy(event.target.checked);
  };

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
                helperText="Name of SQLite database file"
                onKeyDown={focusNext}
              />
            </Grid>
            <Grid item xs={5}></Grid>
          </Grid>

          {/* Session Name grid */}
          <Grid container item spacing={2}>
            <Grid item xs={2}></Grid>
            <Grid item xs={5}>
              <InputName type="session" onKeyDown={focusNext} />
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
              <InputIpRange id={1} required={true} onKeyDown={focusNext} />
            </Grid>
            <Grid item xs={5}>
              <InputProxy id={1} visible={proxy} onKeyDown={focusNext} />
            </Grid>
          </Grid>

          {/* IP Range 2 grid */}
          <Grid container item spacing={2}>
            <Grid item xs={2}></Grid>
            <Grid item xs={5}>
              <InputIpRange id={2} onKeyDown={focusNext} />
            </Grid>
            <Grid item xs={5}>
              <InputProxy id={2} visible={proxy} onKeyDown={focusNext} />
            </Grid>
          </Grid>

          {/* IP Range 3 grid */}
          <Grid container item spacing={2}>
            <Grid item xs={2}></Grid>
            <Grid item xs={5}>
              <InputIpRange id={3} onKeyDown={focusNext} />
            </Grid>
            <Grid item xs={5}>
              <InputProxy id={3} visible={proxy} onKeyDown={focusNext} />
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
