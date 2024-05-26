import { TextField } from "@mui/material";
import React from "react";

const InputProxy = ({
  id,
  visible,
  onChange,
  onKeyDown,
}: {
  id: number;
  visible: boolean;
  onChange: React.ChangeEventHandler<HTMLInputElement>;
  onKeyDown: React.KeyboardEventHandler<HTMLInputElement>;
}) => {
  return (
    <TextField
      id={"proxy-" + id}
      label={"Proxy " + id}
      variant="outlined"
      fullWidth
      sx={{
        opacity: visible ? 1 : 0,
        transition: "all 0.5s",
        visibility: visible ? "visible" : "hidden",
      }}
      onChange={onChange}
      onKeyDown={onKeyDown}
    />
  );
};

export default InputProxy;
