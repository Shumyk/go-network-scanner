import { HostDTO, ServiceDTO } from "./dataTransferObject";

const FOUND = "found";
const LOST = "lost";
const ONLINE = "online";
const OFFLINE = "offline";
const UNKNOWN = "unknown";
type TypeStatus = "found" | "lost" | "online" | "offline" | "unknown";

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

export { HostType, ServiceType };
export type { TypeStatus };
export { ONLINE, OFFLINE, FOUND, LOST, UNKNOWN };