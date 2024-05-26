import { HostType } from "./table/dataTableObject";

const SCANNED = "scanned";
const CONFIGURED = "configured";

abstract class Storage {
  public static isConfigured(): boolean {
    return sessionStorage.getItem(CONFIGURED) === "true";
  }
  public static configured(): void {
    sessionStorage.setItem(CONFIGURED, "true");
  }

  public static setScanned(data: HostType[]): void {
    sessionStorage.setItem(SCANNED, JSON.stringify(data));
  }
    public static scanned(): HostType[] {
        const scanned = sessionStorage.getItem(SCANNED);
        if (!scanned) return [];
        return JSON.parse(scanned).map((host: HostType) => HostType.assign(host));
    }
}

export default Storage;
