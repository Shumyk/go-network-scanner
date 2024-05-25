import React, { useCallback, useEffect, useState } from "react";
import { Divider, Empty, Spin, Table } from "antd";
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

  const expandedRowRender = (showHeader: boolean) => {
    return (
      host: HostType,
      _number: number,
      _indent: number,
      expanded: boolean
    ) => {
      console.info(`key ${host.key}, row ${_number},expanded ${expanded}`);

      if (!host.services.length)
        return <Empty image={Empty.PRESENTED_IMAGE_SIMPLE} />;
      return (
        <Table
          columns={serviceColumns()}
          dataSource={host.services}
          pagination={false}
          bordered
          showHeader={showHeader}
          size="small"
        />
      );
    };
  };

  return (
    <>
      <Table
        title={() => "Scanned devices"}
        columns={hostColumns}
        expandable={{
          expandedRowRender: expandedRowRender(true),
          expandedRowKeys: tableData.map((host) => host.key),
        }}
        dataSource={tableData.length > 0 ? [tableData[0]] : []}
        bordered={true}
        locale={{
          emptyText: <Spin size="large" tip="Scanning..." />,
        }}
        pagination={false}
      />

      {tableData.slice(1).map((host) => (
        <>
          <Divider key={host.key + 'divider'} />
          <Table
            key={host.key}
            columns={hostColumns}
            showHeader={false}
            expandable={{
              expandedRowRender: expandedRowRender(false),
              expandedRowKeys: tableData.map((host) => host.key),
            }}
            dataSource={[host]}
            bordered={true}
            pagination={false}
          />
        </>
      ))}
    </>
  );
};

export default ScanGridTable;
