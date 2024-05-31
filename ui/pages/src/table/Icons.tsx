import { Type } from "./dataTableObject";
import {
  WindowsIcon,
  LinuxIcon,
  AppleIcon,
  OperativeSystemIcon,
  ComputerIcon,
  ServiceIcon,
} from "./Svg";

const OsIcon = ({ os }: { os: string }) => {
  if (os.includes("Windows")) {
    return <WindowsIcon os={os} />;
  }
  if (os.includes("Linux")) {
    return <LinuxIcon os={os} />;
  }
  if (os.includes("macOS") || os.includes("iOS")) {
    return <AppleIcon os={os} />;
  }
  return <OperativeSystemIcon os={os} />;
};

const renderIcon = (type: Type) => {
  return {
    props: { className: type.classNames() },
    children: icon(),
  };

  function icon() {
    if (type.icon === "computer") return <ComputerIcon />;
    else return <ServiceIcon />;
  }
};

export { OsIcon, renderIcon };
