import { TextField } from "@mui/material";

const InputName = ({
  type,
  helperText,
  onKeyDown,
}: {
  type: string;
  helperText?: string;
  onKeyDown: any;
}) => {
  return (
    <TextField
    id={type + "-name"}
    label={type[0].toUpperCase() + type.slice(1) + " Name"}
    helperText={helperText}
    variant="outlined"
    fullWidth
    required
    onKeyDown={onKeyDown}
  />
  );
};

export default InputName;
