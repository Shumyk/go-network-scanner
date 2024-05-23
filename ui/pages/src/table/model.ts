class ScanDTO {
  constructor(public data: HostType[]) {}

  hydrate(): HostType[] {
    // Assign unique keys to HostType and ServiceType
    return this.data.map((host) => ({
      ...host,
      key: host.ip, // Use a unique identifier for the key
      services: host.services.map((service) => ({
        ...service,
        key: `${host.ip}-${service.port}`, // Use a combination of host IP and port for the key
      })),
    }));
  }
}

interface HostType {
  key: React.Key;
  ip: string;
  os: string;
  hostnameLabel: string;
  customLabel: string;

  services: ServiceType[];

  firstDiscovered: string;
  lastDiscovered: string;
  scannedAt: string;
}

interface ServiceType {
  key: React.Key;
  ip: string;
  port: string;
  banner: string;

  firstDiscovered: string;
  lastDiscovered: string;
}

export { ScanDTO };
export type { HostType, ServiceType };
