import { HostType } from "./model";

export function sampleData(): HostType[] {
  const data: HostType[] = [];
  for (let i = 0; i < 3; ++i) {
    const ip = "193.23.43." + i;
    data.push({
      key: i.toString(),
      ip: ip,
      os: i % 2 === 0 ? "Windows" : "Linux",
      hostnameLabel: "hostname" + i,
      customLabel: "custom" + i,
      services: [
        {
          key: ip + "80",
          ip: "193.23.43." + i,
          port: "80",
          banner: "Apache",
          firstDiscovered: new Date().toLocaleTimeString(),
          lastDiscovered: new Date().toLocaleTimeString(),
        },
        {
          key: ip + "443",
          ip: "193.23.43." + i,
          port: "443",
          banner: "Nginx",
          firstDiscovered: new Date().toLocaleTimeString(),
          lastDiscovered: new Date().toLocaleTimeString(),
        },
      ],
      firstDiscovered: new Date().toLocaleTimeString(),
      lastDiscovered: new Date().toLocaleTimeString(),
    });
  }
  return data;
}
