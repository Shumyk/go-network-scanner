class ScanDTO {
  constructor(public data: HostDTO[]) {}
}

class TableData {
  public data: HostType[] = [];

  merge(newData: HostDTO[]): TableData {
    const merged: HostType[] = [];

    // Iterate through the newly scanned data to understand which hosts are new or existing.
    for (const newHost of newData) {
      const index = this.data.findIndex((p) => newHost.ip === p.ip);
      if (index === -1) {
        // NEW HOST FOUND
        const found: HostType = new HostType(newHost);
        this.markAs(found, FOUND);
        merged.push(found);
      } else {
        // HOST ONLINE - merge its services
        const online: HostType = this.mergeServices(this.data[index], newHost);
        online.status = ONLINE;
        merged.push(online);
      }
    }

    // Iterate through the existing data to understand which hosts are lost.
    for (const prevHost of this.data) {
      const newHost = newData.find((n) => prevHost.ip === n.ip);
      // ONLINE - ignore as it's done above
      if (newHost) continue;

      // Host is lost.
      // If the host was present, transition to LOST status.
      if (prevHost.status === FOUND) prevHost.status = LOST;
      else if (prevHost.status === ONLINE) prevHost.status = LOST;
      // If the host was lost already, transition or keep to OFFLINE status.
      else if (prevHost.status === LOST) prevHost.status = OFFLINE;
      else if (prevHost.status === OFFLINE) prevHost.status = OFFLINE;

      this.markAs(prevHost, prevHost.status);
      merged.push(prevHost);
    }

    const mergedTable = new TableData();
    mergedTable.data = merged;
    console.info("Merged data: ", mergedTable);
    return mergedTable;
  }

  mergeServices(prevHost: HostType, newHost: HostDTO): HostType {
    const merged: HostType = new HostType(newHost, []);
    merged.services = [];

    for (const newService of newHost.services) {
      const prevService = prevHost.services.find(
        (s) => s.port === newService.port
      );

      if (prevService) newService.status = ONLINE;
      else newService.status = FOUND;

      merged.services.push(newService);
    }

    for (const prevService of prevHost.services) {
      const newService = newHost.services.find(
        (s) => s.port === prevService.port
      );
      // ONLINE - ignore as it's done above
      if (newService) continue;

      // Service not present in the new data:
      // Mark the service as lost if it was found or online before.
      if (prevService.status === FOUND) prevService.status = LOST;
      else if (prevService.status === ONLINE) prevService.status = LOST;
      // Mark the service as offline if it was lost or offline before.
      else if (prevService.status === LOST) prevService.status = OFFLINE;
      else if (prevService.status === OFFLINE) prevService.status = OFFLINE;

      merged.services.push(prevService);
    }

    return merged;
  }

  markAs(h: HostType, status: TypeStatus): void {
    h.status = status;
    h.services.forEach((service) => (service.status = status));
  }
}

const FOUND = "found";
const LOST = "lost";
const ONLINE = "online";
const OFFLINE = "offline";
type TypeStatus = "found" | "lost" | "online" | "offline" | "unknown";

interface HostDTO {
  // key: React.Key;
  ip: string;
  os: string;
  hostnameLabel: string;
  customLabel: string;

  services: ServiceType[];

  firstDiscovered: string;
  lastDiscovered: string;
  scannedAt: string;

  // status: TypeStatus;
}

class HostType {
  key: React.Key;
  ip: string;
  os: string;
  hostnameLabel: string;
  customLabel: string;

  services: ServiceType[];

  firstDiscovered: string;
  lastDiscovered: string;
  scannedAt: string;

  status: TypeStatus;

  constructor(host: HostDTO, services: ServiceType[] = []) {
    this.key = host.ip; // Use a unique identifier for the key
    this.ip = host.ip;
    this.os = host.os;
    this.hostnameLabel = host.hostnameLabel;
    this.customLabel = host.customLabel;
    this.services = services || host.services;
    this.firstDiscovered = host.firstDiscovered;
    this.lastDiscovered = host.lastDiscovered;
    this.scannedAt = host.scannedAt;
    this.status = "unknown";
  }

  public classNames(): string {
    return this.status;
  }
}

interface ServiceType {
  key: React.Key;
  ip: string;
  port: string;
  banner: string;

  firstDiscovered: string;
  lastDiscovered: string;
  scannedAt: string;

  status: TypeStatus;
}

export { ScanDTO, HostType, TableData };
export type { HostDTO, ServiceType };
