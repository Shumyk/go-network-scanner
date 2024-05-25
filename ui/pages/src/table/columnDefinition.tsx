import { OsIcon } from "./OsIcon";
import { Image, TableColumnsType } from "antd";
import { HostType, ServiceType } from "./dataTableObject";

const hostColumns: TableColumnsType<HostType> = [
  {
    title: "❇️",
    dataIndex: "icon",
    key: "icon",
    align: "center",
    width: "5%",
    className: "icon-cell",
    render: (_, record: HostType) => {
      return {
        props: { className: record.classNames() },
        children: <Image src="/computer.png" width={35} preview={false} />,
      };
    },
  },
  { title: "IP", dataIndex: "ip", key: "ip", width: "14.1%" },
  {
    title: "OS",
    dataIndex: "os",
    key: "os",
    align: "center",
    className: "icon-cell",
    render: (os: string) => <OsIcon os={os} />,
  },
  {
    title: "Hostname",
    dataIndex: "hostnameLabel",
    key: "hostnameLabel",
    align: "center",
    width: "21.5%",
  },
  {
    title: "Custom Label",
    dataIndex: "customLabel",
    key: "customLabel",
    align: "center",
    width: "21.5%",
  },
  {
    title: "Discovered",
    dataIndex: "firstDiscovered",
    key: "firstDiscovered",
    align: "center",
    width: "14.5%",
  },
  {
    title: "Last seen",
    dataIndex: "lastDiscovered",
    key: "lastDiscovered",
    align: "center",
    width: "14.5%",
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
      width: "5%",
    },
    {
      title: "",
      dataIndex: "icon",
      key: "icon",
      align: "center",
      className: "icon-cell",
      width: "5%",
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
      width: "10%",
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
      width: "15%",
    },
    {
      title: "Last seen",
      dataIndex: "lastDiscovered",
      key: "lastDiscovered",
      align: "center",
      width: "15%",
    },
  ];
};

export { hostColumns, serviceColumns };
