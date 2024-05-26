import { Theme, createTheme } from "@mui/material";
import { pink } from "@mui/material/colors";
import { ThemeConfig } from "antd";

class Themes {
  public mui: Theme;
  public antd: ThemeConfig;

  constructor() {
    this.mui = createTheme({
      palette: {
        mode: "dark",
        primary: {
          main: pink[200],
        },
      },
    });
    this.antd = {
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
    };
  }
}

const themes = new Themes();
export default themes;
