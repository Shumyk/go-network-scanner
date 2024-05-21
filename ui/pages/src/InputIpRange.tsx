import { TextField } from "@mui/material";
import React from "react";

const InputIpRange = ({
  id,
  required = false,
  value = "",
  onChange,
  onKeyDown,
}: {
  id: number;
  required?: boolean;
  value?: string;
  onChange: React.ChangeEventHandler<HTMLInputElement>;
  onKeyDown: React.KeyboardEventHandler<HTMLInputElement>;
}) => {
  return (
    <TextField
      id={"ip-range-" + id}
      label={"IP Range " + id}
      value={value}
      variant="outlined"
      fullWidth
      required={required}
      onChange={onChange}
      onKeyDown={onKeyDown}
    />
  );
};

export default InputIpRange;
