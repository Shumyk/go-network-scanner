import { Image, Tooltip } from "antd";

const OsIcon: React.FC<{ os: string }> = ({ os }) => {
  let image = "/operative-system.png";
  if (os.includes("Windows")) {
    image = "/windows.png";
  }
  if (os.includes("Linux")) {
    image = "/linux.png";
  }
  if (os.includes("macOS") || os.includes("iOS")) {
    image = "/apple.png";
  }
  return (
    <Tooltip title={os} placement="top" color="#e271af">
      <Image src={image} width={35} preview={false} />
    </Tooltip>
  );
};

export { OsIcon };