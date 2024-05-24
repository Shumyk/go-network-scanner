class ScanDTO {
  constructor(public data: HostDTO[]) {}
}

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

export { ScanDTO };
export type { HostDTO, ServiceDTO };