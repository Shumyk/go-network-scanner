import { HostType } from "./dataTableObject";
import { ScanDTO } from "./dataTransferObject";
import { merge } from "./merger";
import { sort } from "./sorter";
import Storage from "../common/storage";

class ScannedDataListener {
  private static _instance: ScannedDataListener;

  private tableData: HostType[];
  private setTableData: React.Dispatch<React.SetStateAction<HostType[]>>;

  private constructor(
    tableData: HostType[],
    setTableData: React.Dispatch<React.SetStateAction<HostType[]>>
  ) {
    this.tableData = tableData;
    this.setTableData = setTableData;
  }

  public static instance(
    tableData: HostType[],
    setTableData: React.Dispatch<React.SetStateAction<HostType[]>>
  ): ScannedDataListener {
    if (!ScannedDataListener._instance)
      ScannedDataListener._instance = new ScannedDataListener(
        tableData,
        setTableData
      );
    return ScannedDataListener._instance;
  }

  public bind(): void {
    // @ts-ignore
    window.addScannedDevice = this.process;
  }

  public unbind(): void {
    // @ts-ignore
    window.addScannedDevice = undefined;
  }

  public process(json: string): void {
    const scan: ScanDTO = JSON.parse(json);
    console.log("Received scanned devices", scan);

    const merged: HostType[] = merge(this.tableData, scan.data);
    const sorted: HostType[] = sort(merged);
    this.setTableData(sorted);

    Storage.setScanned(sorted);
  }
}

export default ScannedDataListener;
