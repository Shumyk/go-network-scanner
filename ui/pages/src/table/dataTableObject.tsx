import { HostDTO, ServiceDTO } from "./dataTransferObject";

// ### Entity Statuses.
// - `FOUND` - Entity was found. It was absent in the previous scan and present in the current scan.
const FOUND = "found";
// - `LOST` - Entity was lost. It was present in the previous scan and absent in the current scan.
const LOST = "lost";
// - `ONLINE` - Entity is online. It was present in the previous scan and present in the current scan.
const ONLINE = "online";
// - `OFFLINE` - Entity is offline. It was absent in the previous scan and absent in the current scan.
const OFFLINE = "offline";
const UNKNOWN = "unknown";
type TypeStatus = "found" | "lost" | "online" | "offline" | "unknown";

interface Type {
  icon: string;
  classNames(): string;
}

class HostType implements Type {
  static assign(host: HostType): HostType {
    return new this(
      host.key,
      host.ip,
      host.os,
      host.hostnameLabel,
      host.customLabel,
      host.status,
      host.services.map((s) => ServiceType.assign(s)),
      host.firstDiscovered,
      host.lastDiscovered
    );
  }
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

  icon: string = "computer.png";
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

class ServiceType implements Type {
  static assign(service: ServiceType): ServiceType {
    return new this(
      service.key,
      service.ip,
      service.port,
      service.banner,
      service.firstDiscovered,
      service.lastDiscovered,
      service.status
    );
  }
  static of(service: ServiceDTO, status: TypeStatus = UNKNOWN): ServiceType {
    return new this(
      service.ip + ":" + service.port,
      service.ip,
      service.port,
      service.banner,
      service.firstDiscovered,
      service.lastDiscovered,
      status
    );
  }

  icon: string = "service.png";
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

export { HostType, ServiceType };
export type { TypeStatus, Type };
export { ONLINE, OFFLINE, FOUND, LOST, UNKNOWN };
