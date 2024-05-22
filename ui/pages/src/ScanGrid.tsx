import { useEffect, useState, useCallback } from "react";
import {
  DataGrid,
  GridRowsProp,
  GridRowModel,
  GridColDef,
} from "@mui/x-data-grid";

const columns: GridColDef[] = [
  { field: "ip", headerName: "IP", width: 120 },
  { field: "port", headerName: "Port", width: 50 },
  { field: "service", headerName: "Service", width: 350 },
];

interface ScansDTO {
  data: Device[];
}

interface Device {
  ip: string;
  port: string;
  service: string;
}

const ScanGrid = () => {
  const [rows, setRows] = useState<GridRowsProp>([]);

  const addScannedDevice = useCallback((json: string): void => {
    const scans: ScansDTO = JSON.parse(json);
    console.log("Received scanned devices", scans);

    const scannedDevices: GridRowModel[] = scans.data.map((device, id) => {
      return {
        id: id,
        ip: device.ip,
        port: device.port,
        service: device.service,
      };
    });

    setRows(scannedDevices);
  }, []);

  useEffect(() => {
    // @ts-ignore
    window.addScannedDevice = (device) => {
      addScannedDevice(device);
    };
  }, [addScannedDevice]);

  return (
    <div>
      <h1>Scan Grid</h1>
      <div id="scan-list">
        <DataGrid rows={rows} columns={columns} />
        {/* {scannedDevices.map((device) => (
          <ScanGridItem device={device} />
        ))} */}
      </div>
    </div>
  );
};

export default ScanGrid;
