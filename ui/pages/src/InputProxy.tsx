import { TextField } from "@mui/material";

const InputProxy = ({
  id,
  visible,
  onKeyDown,
}: {
  id: number;
  visible: boolean;
  onKeyDown: any;
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
      onKeyDown={onKeyDown}
    />
  );
};

export default InputProxy;
