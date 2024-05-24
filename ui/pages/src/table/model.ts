class ScanDTO {
  constructor(public data: HostDTO[]) {}
}

class TableData {
  public data: HostType[] = [];

  merge(newData: HostDTO[]): TableData {
    const merged: HostType[] = [];

    // Iterate through the newly scanned data to understand which hosts are new or existing.
    for (const newHost of newData) {
      const prevHost = this.findByIp(this.data, newHost);

      if (prevHost) {
        // Entity continue to be present, transiting to ONLINE state.
        const online: HostType = this.mergeServices(prevHost, newHost);
        online.status = ONLINE;
        merged.push(online);
      } else {
        // New host found, mark all its services as FOUND.
        merged.push(HostType.found(newHost));
      }
    }

    // Iterate through the existing data to understand which hosts are lost.
    for (const prevHost of this.data) {
      const newHost = this.findByIp(newData, prevHost);
      // Entry was present before and it's present now. It was handled above.
      if (newHost) continue;

      // Entry is lost, transiting to next state depending on the current status.
      this.markAs(prevHost, this.lostStatusTransition(prevHost.status));
      merged.push(prevHost);
    }

    const mergedTable = new TableData();
    mergedTable.data = merged;
    console.info("Merged data: ", mergedTable);
    return mergedTable;
  }

  findByIp<T extends HasIP, S extends HasIP>(
    hosts: T[],
    target: T | S
  ): T | undefined {
    return hosts.find((h) => h.ip === target.ip);
  }

  findByPort<T extends HasPort, S extends HasPort>(
    services: T[],
    target: T | S
  ): T | undefined {
    return services.find((s) => s.port === target.port);
  }

  // Depending on the current status, decide to which status to transition to.
  lostStatusTransition(status: TypeStatus): TypeStatus {
    // Currently, entity is present, so it should transition to LOST.
    if (status === FOUND) return LOST;
    else if (status === ONLINE) return LOST;
    // If it is absent already, keep OFFLINE status.
    else if (status === LOST) return OFFLINE;
    else if (status === OFFLINE) return OFFLINE;
    return status;
  }

  foundStatusTransition(previous: any): TypeStatus {
    // Entity present, so it should transition to ONLINE.
    if (previous) return ONLINE;
    // Previous entry is absent, so it's newly found entity.
    return FOUND;
  }

  mergeServices(prevHost: HostType, newHost: HostDTO): HostType {
    const merged: HostType = HostType.of(newHost);

    // Iterate through the newly scanned data to understand which services are new or existing.
    for (const newService of newHost.services) {
      const prevService = this.findByPort(prevHost.services, newService);
      merged.services.push(
        ServiceType.of(newService, this.foundStatusTransition(prevService))
      );
    }

    for (const prevService of prevHost.services) {
      const newService = this.findByPort(newHost.services, prevService);
      // Entry was present before and it's present now. It was handled above.
      if (newService) continue;

      // Entry not present in the new data. It was present and lost.
      prevService.status = this.lostStatusTransition(prevService.status);
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
const UNKNOWN = "unknown";
type TypeStatus = "found" | "lost" | "online" | "offline" | "unknown";

interface HostDTO {
  ip: string;
  os: string;
  hostnameLabel: string;
  customLabel: string;

  services: ServiceDTO[];

  firstDiscovered: string;
  lastDiscovered: string;
}

interface ServiceDTO {
  ip: string;
  port: string;
  banner: string;

  firstDiscovered: string;
  lastDiscovered: string;
}

type HasIP = { ip: string };
type HasPort = { port: string };

class HostType {
  static of(host: HostDTO): HostType {
    return new this(
      host.ip, // Use a unique identifier for the key
      host.ip,
      host.os,
      host.hostnameLabel,
      host.customLabel,
      UNKNOWN,
      [],
      host.firstDiscovered,
      host.lastDiscovered
    );
  }
  static found(host: HostDTO): HostType {
    return new this(
      host.ip, // Use a unique identifier for the key
      host.ip,
      host.os,
      host.hostnameLabel,
      host.customLabel,
      FOUND,
      host.services.map((s) => ServiceType.of(s, FOUND)),
      host.firstDiscovered,
      host.lastDiscovered
    );
  }

  constructor(
    public key: React.Key,
    public ip: string,
    public os: string,
    public hostnameLabel: string,
    public customLabel: string,

    public status: TypeStatus = UNKNOWN,
    public services: ServiceType[] = [],

    public firstDiscovered: string,
    public lastDiscovered: string
  ) {}

  public classNames(): string {
    return this.status;
  }
}

class ServiceType {
  static of(service: ServiceDTO, status: TypeStatus = UNKNOWN): ServiceType {
    return new this(
      service.ip + service.port,
      service.ip,
      service.port,
      service.banner,
      service.firstDiscovered,
      service.lastDiscovered,
      status
    );
  }

  constructor(
    public key: React.Key,
    public ip: string,
    public port: string,
    public banner: string,

    public firstDiscovered: string,
    public lastDiscovered: string,

    public status: TypeStatus = UNKNOWN
  ) {}

  public classNames(): string {
    return this.status;
  }
}

export { ScanDTO, HostType, TableData };
export type { HostDTO, ServiceType };
