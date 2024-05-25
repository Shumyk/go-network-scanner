import React, { useCallback, useEffect, useState } from "react";
import { ConfigProvider, Empty, Spin, Table } from "antd";
import "./ScanGridTable.css";
import { merge } from "./merger";
import { sampleData } from "./dataGenerator";
import { hostColumns, serviceColumns } from "./columnDefinition";
import { HostType } from "./dataTableObject";
import { ScanDTO } from "./dataTransferObject";

const ScanGridTable: React.FC = () => {
  const [tableData, setTableData] = useState<HostType[]>([]);

  const addScannedDevice = useCallback(
    (json: string): void => {
      const scan: ScanDTO = JSON.parse(json);
      console.log("Received scanned devices", scan);

      const merged: HostType[] = merge(tableData, scan.data);
      setTableData(merged);
    },
    [tableData, setTableData]
  );

  useEffect(() => {
    // @ts-ignore
    window.addScannedDevice = (device) => {
      addScannedDevice(device);
    };

    setTimeout(() => setTableData(sampleData()), 2000);
  }, [addScannedDevice]);

  const expandedRowRender = (
    host: HostType,
    _number: number,
    _indent: number,
    expanded: boolean
  ) => {
    // console.info("expandedRowRender", host, expanded, host.services);
    console.info(`key ${host.key}, row ${_number},expanded ${expanded}`);

    if (!host.services.length)
      return <Empty image={Empty.PRESENTED_IMAGE_SIMPLE} />;
    return (
      <Table
        columns={serviceColumns()}
        dataSource={host.services}
        pagination={false}
        bordered
        size="small"
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
        title={() => "Scanned devices"}
        columns={hostColumns}
        expandable={{
          expandedRowRender,
          expandedRowKeys: tableData.map((host) => host.key),
        }}
        dataSource={tableData}
        bordered={true}
        locale={{
          emptyText: <Spin size="large" tip="Scanning..." />,
        }}
      /> 
    </ConfigProvider>
  );
};

export default ScanGridTable;
