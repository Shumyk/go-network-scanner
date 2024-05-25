import React from "react";
import { Image, TableColumnsType, Tooltip } from "antd";
import { HostType, ServiceType } from "./dataTableObject";

const OsIcon: React.FC<{
  os: string;
}> = ({ os }) => {
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

const hostColumns: TableColumnsType<HostType> = [
  {
    title: "❇️",
    dataIndex: "icon",
    key: "icon",
    align: "center",
    className: "icon-cell",
    render: (_, record: HostType) => {
      return {
        props: { className: record.classNames() },
        children: <Image src="/computer.png" width={35} preview={false} />,
      };
    },
  },
  { title: "IP", dataIndex: "ip", key: "ip" },
  {
    title: "OS",
    dataIndex: "os",
    key: "os",
    align: "center",
    className: "icon-cell",
    render: (os: string) => <OsIcon os={os} />,
  },
  { title: "Hostname", dataIndex: "hostnameLabel", key: "hostnameLabel" },
  { title: "Custom Label", dataIndex: "customLabel", key: "customLabel" },
  {
    title: "Discovered",
    dataIndex: "firstDiscovered",
    key: "firstDiscovered",
  },
  { title: "Last seen", dataIndex: "lastDiscovered", key: "lastDiscovered" },
];

const IconCell = ({ record }: { record: ServiceType }) => {
  // const [found, setFound] = React.useState("");
  // const [lost, setLost] = React.useState("");

  // useEffect(() => {
  //   if (record.scannedAt === record.lastDiscovered) {
  //     setFound("found");
  //   } else {
  //     setLost("lost");
  //   }
  // }, [record.scannedAt, record.lastDiscovered]);

  // const classes = `service-icon-cell ${found} ${lost}`;
  return (
    // <div className="icon-cell" style={{ width: 35}}>
    <Image src="/service.png" width={35} preview={false} />
    // </div>
  );
};

const serviceColumns = (): TableColumnsType<ServiceType> => {
  return [
    {
      title: "❇️",
      align: "center",
      className: "empty-cell",
    },
    {
      title: "",
      dataIndex: "icon",
      key: "icon",
      align: "center",
      className: "icon-cell",
      render: (_, record: ServiceType) => {
        return {
          props: { className: record.classNames() },
          children: <IconCell record={record} />,
        };
      },
    },
    {
      title: "Port",
      dataIndex: "port",
      key: "port",
      align: "center",
    },
    {
      title: "Banner",
      dataIndex: "banner",
      key: "banner",
      align: "center",
    },
    {
      title: "Discovered",
      dataIndex: "firstDiscovered",
      key: "firstDiscovered",
      align: "center",
    },
    {
      title: "Last seen",
      dataIndex: "lastDiscovered",
      key: "lastDiscovered",
      align: "center",
    },
  ];
};

export { hostColumns, serviceColumns };
