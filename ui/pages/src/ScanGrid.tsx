import { useEffect, useState, useCallback } from "react";
import ScanGridItem from "./ScanGridItem";

const ScanGrid = () => {
  const [scannedDevices, setScannedDevices] = useState<string[]>([]);

  const addScannedDevice = useCallback(
    (scans: string): void => {
      console.log("Received scans", scans);
      const devices = JSON.parse(scans);
      console.log("Parsed devices", devices);
      // setScannedDevices([...scannedDevices, device]);
    },
    [scannedDevices]
  );

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
        {scannedDevices.map((device) => (
          <ScanGridItem device={device} />
        ))}
      </div>
    </div>
  );
};

export default ScanGrid;
