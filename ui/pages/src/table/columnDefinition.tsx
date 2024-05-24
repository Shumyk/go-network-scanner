import React, { useEffect } from "react";
import { Dropdown, Image, Space, TableColumnsType, Tooltip } from "antd";
import { HostType, ServiceType } from "./model";
import { DownOutlined } from "@ant-design/icons";

const items = [
  { key: "1", label: "Action 1" },
  { key: "2", label: "Action 2" },
];

const OsIcon: React.FC<{
  os: string;
}> = ({ os }) => {
  let image = "/computer.png";
  if (os.includes("Windows")) {
    image = "/windows.png";
  }
  if (os.includes("Linux")) {
    image = "/linux.png";
  }
  if (os.includes("macOS")) {
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
        props: {
          className: `${
            record.scannedAt === record.lastDiscovered ? "found" : "lost"
          } `,
        },
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
  { title: "Action", key: "operation", render: () => <a>Publish</a> },
];

const IconCell = ({ record }: { record: ServiceType }) => {
  const [found, setFound] = React.useState("");
  const [lost, setLost] = React.useState("");

  useEffect(() => {
    if (record.scannedAt === record.lastDiscovered) {
      setFound("found");
    } else {
      setLost("lost");
    }
  }, [record.scannedAt, record.lastDiscovered]);

  // const classes = `service-icon-cell ${found} ${lost}`;
  return (
    // <div className={classes}>
    <Image src="/service.png" width={35} preview={false} />
    // </div>
  );
};

const serviceColumns = (): TableColumnsType<ServiceType> => {
  return [
    {
      title: "❇️",
      dataIndex: "icon",
      key: "icon",
      align: "center",
      // className: "icon-cell",
      render: (_, record: ServiceType) => {
        setTimeout(() => {}, 8000);
        return {
          props: {
            className: `${
              record.scannedAt === record.lastDiscovered
                ? "icon-cell found"
                : "icon-cell lost"
            } `,
          },
          children: <IconCell record={record} />,
        };
      },
    },
    { title: "Port", dataIndex: "port", key: "port" },
    { title: "Banner", dataIndex: "banner", key: "banner" },
    {
      title: "Discovered",
      dataIndex: "firstDiscovered",
      key: "firstDiscovered",
    },
    {
      title: "Last seen",
      dataIndex: "lastDiscovered",
      key: "lastDiscovered",
    },
    {
      title: "Action",
      key: "operation",
      render: () => (
        <Space size="middle">
          <a>Pause</a>
          <a>Stop</a>
          <Dropdown menu={{ items }}>
            <a>
              More <DownOutlined />
            </a>
          </Dropdown>
        </Space>
      ),
    },
  ];
};

export { hostColumns, serviceColumns };
