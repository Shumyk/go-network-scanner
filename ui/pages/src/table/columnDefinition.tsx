import { OsIcon } from "./OsIcon";
import { Image, TableColumnsType } from "antd";
import { HostType, ServiceType } from "./dataTableObject";

const hostColumns: TableColumnsType<HostType> = [
  {
    title: "❇️",
    dataIndex: "icon",
    key: "icon",
    align: "center",
    width: 48 + 12, // = 60px, 48px is the width of nested table margin
    className: "icon-cell",
    render: (_, record: HostType) => {
      return {
        props: { className: record.classNames() },
        children: <Image src="/computer.png" width={35} preview={false} />,
      };
    },
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
    width: 170,
  },
  {
    title: "Last seen",
    dataIndex: "lastDiscovered",
    key: "lastDiscovered",
    align: "center",
    width: 171, // ideal to fit date time, plus 1 for margin
  },
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
      width: 60, // the same as the width of the icon in hostColumns
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
      width: 170,
    },
    {
      title: "Last seen",
      dataIndex: "lastDiscovered",
      key: "lastDiscovered",
      align: "center",
      width: 170,
    },
  ];
};

export { hostColumns, serviceColumns };
