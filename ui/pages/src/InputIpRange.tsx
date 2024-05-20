import { TextField } from "@mui/material";

const InputIpRange = ({ id, required = false, onKeyDown }: { id: number; required?: boolean; onKeyDown: any }) => {
  return (
    <TextField
      id={"ip-range-" + id}
      label={"IP Range " + id}
      variant="outlined"
      fullWidth
      required={required}
      onKeyDown={onKeyDown}
    />
  );
};

export default InputIpRange;
