import React, { useCallback, useEffect, useState } from "react";
import { ConfigProvider, Table } from "antd";
import "./ScanGridTable.css";
import { HostType, ScanDTO } from "./model";
import { sampleData } from "./dataGenerator";
import { hostColumns, serviceColumns } from "./columnDefinition";

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
    return (
      <Table
        columns={serviceColumns()}
        dataSource={host.services}
        pagination={false}
        bordered
      />
    );
  };

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
        columns={hostColumns}
        expandable={{
          expandedRowRender,
          defaultExpandAllRows: true,
          rowExpandable: () => true,
        }}
        // dataSource={rows}
        dataSource={rows}
        bordered={true}
      />
    </ConfigProvider>
  );
};

export default ScanGridTable;
