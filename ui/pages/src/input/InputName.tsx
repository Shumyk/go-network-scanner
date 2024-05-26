import { TextField } from "@mui/material";

const InputName = ({
  type,
  value,
  helperText,
  focus,
  onChange,
  onKeyDown,
}: {
  type: string;
  value?: string;
  helperText?: string;
  focus?: boolean;
  onChange: React.ChangeEventHandler<HTMLInputElement>;
  onKeyDown: React.KeyboardEventHandler<HTMLInputElement>;
}) => {
  return (
    <TextField
      id={type + "-name"}
      value={value}
      label={type[0].toUpperCase() + type.slice(1) + " Name"}
      helperText={helperText}
      variant="outlined"
      inputRef={(input) => focus && input?.focus()}
      fullWidth
      required
      onChange={onChange}
      onKeyDown={onKeyDown}
    />
  );
};

export default InputName;
