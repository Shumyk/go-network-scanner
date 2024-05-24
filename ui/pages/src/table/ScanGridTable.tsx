import React, { useCallback, useEffect, useState } from "react";
import { ConfigProvider, Empty, Table } from "antd";
import "./ScanGridTable.css";
import { HostType, ScanDTO, TableData } from "./model";
import { sampleData } from "./dataGenerator";
import { hostColumns, serviceColumns } from "./columnDefinition";

const ScanGridTable: React.FC = () => {
  const [tableData, setTableData] = useState<TableData>(new TableData());

  const addScannedDevice = useCallback(
    (json: string): void => {
      const scan: ScanDTO = JSON.parse(json);
      console.log("Received scanned devices", scan);

      const merged: TableData = tableData.merge(scan.data);
      setTableData(merged);
    },
    [tableData, setTableData]
  );

  useEffect(() => {
    // @ts-ignore
    window.addScannedDevice = (device) => {
      addScannedDevice(device);
    };
  }, [addScannedDevice]);

  const expandedRowRender = (
    host: HostType,
    _number: number,
    _indent: number,
    expanded: boolean
  ) => {
    console.info("expandedRowRender", host, expanded, host.services);
    console.info("expanded", expanded);
    if (!host.services.length) return <Empty image={Empty.PRESENTED_IMAGE_SIMPLE} />; 
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
        dataSource={[...tableData.data]}
        bordered={true}
      />
    </ConfigProvider>
  );
};

export default ScanGridTable;
