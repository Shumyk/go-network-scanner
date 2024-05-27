import { HostType } from "./dataTableObject";
import { ScanDTO } from "./dataTransferObject";
import { merge } from "./merger";
import { sort } from "./sorter";
import Storage from "../common/storage";

class ScannedDataListener {
  private static _instance: ScannedDataListener;

  private setTableData: React.Dispatch<React.SetStateAction<HostType[]>>;

  private constructor(
    setTableData: React.Dispatch<React.SetStateAction<HostType[]>>
  ) {
    this.setTableData = setTableData;
  }

  public static instance(
    setTableData: React.Dispatch<React.SetStateAction<HostType[]>>
  ): ScannedDataListener {
    if (!ScannedDataListener._instance)
      ScannedDataListener._instance = new ScannedDataListener(setTableData);
    return ScannedDataListener._instance;
  }

  public process(current: HostType[], json: string): void {
    const scan: ScanDTO = JSON.parse(json);
    console.log("Received scanned devices", scan);

    const merged: HostType[] = merge(current, scan.data);
    const sorted: HostType[] = sort(merged);
    this.setTableData(sorted);

    Storage.setScanned(sorted);
  }
}

export default ScannedDataListener;
