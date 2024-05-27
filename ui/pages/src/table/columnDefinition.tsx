import { OsIcon, renderIcon } from "./Icons";
import { TableColumnsType } from "antd";
import { HostType, ServiceType } from "./dataTableObject";

const hostColumns: TableColumnsType<HostType> = [
  {
    title: "❇️",
    dataIndex: "icon",
    key: "icon",
    align: "center",
    width: 48 + 12, // = 60px, 48px is the width of nested table margin
    className: "icon-cell",
    render: (_, record: HostType) => renderIcon(record),
  },
  {
    title: "IP",
    dataIndex: "ip",
    key: "ip",
    align: "center",
    width: 181, // = 120 + 60 (icon col width) + 1 (margin)
  },
  {
    title: "OS",
    dataIndex: "os",
    key: "os",
    align: "center",
    className: "icon-cell",
    width: 50, // just enough to fit the icon
    render: (os: string) => <OsIcon os={os} />,
  },
  {
    title: "Hostname",
    dataIndex: "hostnameLabel",
    key: "hostnameLabel",
    align: "center",
  },
  {
    title: "Custom Label",
    dataIndex: "customLabel",
    key: "customLabel",
    align: "center",
  },
  {
    title: "Discovered",
    dataIndex: "firstDiscovered",
    key: "firstDiscovered",
    align: "center",
    width: 180,
  },
  {
    title: "Last seen",
    dataIndex: "lastDiscovered",
    key: "lastDiscovered",
    align: "center",
    width: 181, // ideal to fit date time, plus 1 for margin
  },
];

const serviceColumns = (): TableColumnsType<ServiceType> => {
  return [
    {
      title: "❇️",
      align: "center",
      className: "empty-cell",
      width: 60, // the same as the width of the icon in hostColumns
      render: (_, record: ServiceType) => renderIcon(record),
    },
    {
      title: "Port",
      dataIndex: "port",
      key: "port",
      align: "center",
      width: 132, // = (icon width + ip width) - indent width - nested table padding - icon width - margin = (181 + 60) - 32 - 16 - 60 - 1
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
      width: 180,
    },
    {
      title: "Last seen",
      dataIndex: "lastDiscovered",
      key: "lastDiscovered",
      align: "center",
      width: 180,
    },
  ];
};

export { hostColumns, serviceColumns };
