import { Image, Tooltip } from "antd";
import { Type } from "./dataTableObject";

const OsIcon = ({ os }: { os: string }) => {
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

const IconCell = ({ src }: { src: string }) => {
  return <Image src={`/${src}`} width={35} preview={false} />;
};

const renderIcon = (type: Type) => {
  return {
    props: { className: type.classNames() },
    children: <IconCell src={type.icon} />,
  };
};

export { OsIcon, IconCell, renderIcon };
