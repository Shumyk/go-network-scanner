import { HostDTO } from "./dataTransferObject";
import { HostType } from "./dataTableObject";

export function sampleData(): HostType[] {
  const data: HostDTO[] = [];
  for (let i = 0; i < 3; ++i) {
    const ip = "193.23.43." + i;
    data.push({
      ip: ip,
      os: i % 2 === 0 ? "Windows" : "Linux",
      hostnameLabel: "hostname" + i,
      customLabel: "custom" + i,
      services: [
        {
          ip: "193.23.43." + i,
          port: "80",
          banner: "Apache",
          firstDiscovered: new Date().toLocaleTimeString(),
          lastDiscovered: new Date().toLocaleTimeString(),
        },
        {
          ip: "193.23.43." + i,
          port: "443",
          banner: "Nginx",
          firstDiscovered: new Date().toLocaleTimeString(),
          lastDiscovered: new Date().toLocaleTimeString(),
        },
      ],
      firstDiscovered: new Date().toLocaleTimeString(),
      lastDiscovered: new Date().toLocaleTimeString() + "",
    });
  }
  return data.map((d) => HostType.found(d));
}
