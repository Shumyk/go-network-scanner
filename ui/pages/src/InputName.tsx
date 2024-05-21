import { TextField } from "@mui/material";

const InputName = ({
  type,
  helperText,
  onChange,
  onKeyDown,
}: {
  type: string;
  helperText?: string;
  onChange: React.ChangeEventHandler<HTMLInputElement>;
  onKeyDown: React.KeyboardEventHandler<HTMLInputElement>;
}) => {
  return (
    <TextField
    id={type + "-name"}
    label={type[0].toUpperCase() + type.slice(1) + " Name"}
    helperText={helperText}
    variant="outlined"
    fullWidth
    required
    onChange={onChange}
    onKeyDown={onKeyDown}
  />
  );
};

export default InputName;
