const ScanGrid = () => {
  const addScannedDevice = (device: string) => {
    console.log("Adding scanned device:", device);
    const scanList = document.getElementById("scan-list");
    const scanItem = document.createElement("p");
    scanItem.innerHTML = device;
    scanList?.appendChild(scanItem);
  };
  addScannedDevice("Scanned Devices:");

  return (
    <div>
      <h1>Scan Grid</h1>
      <div id="scan-list"></div>
    </div>
  );
};

export default ScanGrid;
