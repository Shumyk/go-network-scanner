import { Badge, Dropdown, Image, Space, TableColumnsType } from "antd";
import { HostType, ServiceType } from "./model";
import { DownOutlined } from "@ant-design/icons";

const items = [
  { key: "1", label: "Action 1" },
  { key: "2", label: "Action 2" },
];

const hostColumns: TableColumnsType<HostType> = [
  { title: "IP", dataIndex: "ip", key: "ip" },
  {
    title: "OS",
    dataIndex: "os",
    key: "os",
    align: "center",
    render: (os: string) => {
      if (os.includes("Windows")) {
        return <Image src="/windows.png" width={23} preview={false} />;
      }
      if (os.includes("Linux")) {
        return <Image src="/linux.png" width={23} preview={false} />;
      }
      if (os.includes("Mac")) {
        return <Image src="/apple.png" width={23} preview={false} />;
      }
      return os;
    },
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

const serviceColumns: TableColumnsType<ServiceType> = [
  { title: "Port", dataIndex: "port", key: "port" },
  { title: "Banner", dataIndex: "banner", key: "banner" },
  {
    title: "Status",
    key: "state",
    render: () => <Badge status="success" text="Finished" />,
  },
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

export { hostColumns, serviceColumns };
