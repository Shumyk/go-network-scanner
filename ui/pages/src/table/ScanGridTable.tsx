import React, { useEffect, useState } from "react";
import { Divider, Empty, Spin, Table } from "antd";
import "./ScanGridTable.css";
import { sampleData } from "./dataGenerator";
import { hostColumns, serviceColumns } from "./columnDefinition";
import { HostType } from "./dataTableObject";
import Storage from "../common/storage";
import ScannedDataListener from "./ScannedDataListener";

const ScanGridTable: React.FC = () => {
  const [tableData, setTableData] = useState<HostType[]>(Storage.scanned());

  // TODO: add proxy to the backend.
  // TODO: remove scannedAt

  useEffect(() => {
    const listener = ScannedDataListener.instance(setTableData);
    // @ts-ignore
    window.addScannedDevice = (json) => listener.process(tableData, json);

    // For development, trigger sample data event after 10 seconds.
    // @ts-ignore
    !window.webui &&
      setTimeout(
        () => listener.process(tableData, JSON.stringify(sampleData())),
        10000
      );

    // @ts-ignore
    return () => (window.addScannedDevice = undefined);
  }, [tableData, setTableData]);

  const nestedTableRender = (showHeader: boolean) => {
    return (host: HostType) => {
      if (!host.services.length)
        return <Empty image={Empty.PRESENTED_IMAGE_SIMPLE} />;
      return (
        <Table
          columns={serviceColumns()}
          dataSource={host.services}
          bordered
          size="small"
          pagination={false}
          showHeader={showHeader}
        />
      );
    };
  };

  return (
    <>
      <Table
        title={() => "Scanned devices"}
        columns={hostColumns}
        dataSource={!tableData.length ? [] : [tableData[0]]}
        expandable={{
          showExpandColumn: false,
          expandedRowRender: nestedTableRender(true),
          expandedRowKeys: [tableData[0]?.key],
        }}
        bordered={true}
        pagination={false}
        locale={{ emptyText: <Spin size="large" /> }}
      />

      {tableData.slice(1).map((host) => (
        <>
          <Divider key={host.key + "divider"} />
          <Table
            key={host.key}
            columns={hostColumns}
            dataSource={[host]}
            expandable={{
              showExpandColumn: false,
              expandedRowRender: nestedTableRender(false),
              expandedRowKeys: tableData.map((host) => host.key),
            }}
            bordered={true}
            pagination={false}
            showHeader={false}
          />
        </>
      ))}
    </>
  );
};

export default ScanGridTable;
