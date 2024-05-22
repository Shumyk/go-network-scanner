import React, { useCallback, useEffect, useState } from "react";
import { DownOutlined } from "@ant-design/icons";
import type { TableColumnsType } from "antd";
import { Badge, ConfigProvider, Dropdown, Space, Table } from "antd";
import "./ScanGridTable.css";

class ScanDTO {
  constructor(public data: HostType[]) {}

  hydrate(): HostType[] {
    // Assign unique keys to HostType and ServiceType
    return this.data.map((host) => ({
      ...host,
      key: host.ip, // Use a unique identifier for the key
      services: host.services.map((service) => ({
        ...service,
        key: `${host.ip}-${service.port}`, // Use a combination of host IP and port for the key
      })),
    }));
  }
}

interface HostType {
  key: React.Key;
  ip: string;
  os: string;
  hostnameLabel: string;
  customLabel: string;

  services: ServiceType[];

  firstDiscovered: string;
  lastDiscovered: string;
}

interface ServiceType {
  key: React.Key;
  ip: string;
  port: string;
  banner: string;

  firstDiscovered: string;
  lastDiscovered: string;
}

const items = [
  { key: "1", label: "Action 1" },
  { key: "2", label: "Action 2" },
];

const ScanGridTable: React.FC = () => {
  const [rows, setRows] = useState<HostType[]>([]);

  const addScannedDevice = useCallback((json: string): void => {
    const scan: ScanDTO = JSON.parse(json);
    console.log("Received scanned devices", scan);

    const data = new ScanDTO(scan.data).hydrate();
    setRows(data);
  }, []);

  useEffect(() => {
    // @ts-ignore
    window.addScannedDevice = (device) => {
      addScannedDevice(device);
    };
  }, [addScannedDevice]);

  const expandedRowRender = (host: HostType) => {
    const columns: TableColumnsType<ServiceType> = [
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

    return (
      <Table
        columns={columns}
        dataSource={host.services}
        pagination={false}
        bordered
      />
    );
  };

  const columns: TableColumnsType<HostType> = [
    { title: "IP", dataIndex: "ip", key: "ip" },
    { title: "OS", dataIndex: "os", key: "os" },
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

  return (
    <ConfigProvider
      theme={{
        token: {
          colorTextBase: "#ffffff",
          colorBgBase: "#614b9e",
          colorPrimary: "#eb2f96",
          colorInfo: "#e271af",
        },
        components: {
          Table: {
            borderColor: "#ccc1ec",
          },
        },
      }}
    >
      <Table
        columns={columns}
        expandable={{
          expandedRowRender,
          defaultExpandAllRows: true,
          rowExpandable: () => true,
        }}
        dataSource={sampleData()}
        bordered={true}
      />
    </ConfigProvider>
  );
};

export default ScanGridTable;

function sampleData(): HostType[] {
  const data: HostType[] = [];
  for (let i = 0; i < 3; ++i) {
    data.push({
      key: i.toString(),
      ip: "193.23.43." + i,
      os: i % 2 === 0 ? "Windows" : "Linux",
      hostnameLabel: "hostname" + i,
      customLabel: "custom" + i,
      services: [
        {
          key: i.toString(),
          ip: "193.23.43." + i,
          port: "80",
          banner: "Apache",
          firstDiscovered: new Date().toLocaleTimeString(),
          lastDiscovered: new Date().toLocaleTimeString(),
        },
        {
          key: i.toString(),
          ip: "193.23.43." + i,
          port: "443",
          banner: "Nginx",
          firstDiscovered: new Date().toLocaleTimeString(),
          lastDiscovered: new Date().toLocaleTimeString(),
        },
      ],
      firstDiscovered: new Date().toLocaleTimeString(),
      lastDiscovered: new Date().toLocaleTimeString(),
    });
  }
  return data;
}
