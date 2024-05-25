import { HostDTO, ScanDTO } from "./dataTransferObject";

const ips = [
  "34.134.56.2",
  "192.34.65.1",
  "127.0.0.1",
  "193.23.43.43",
  "23.65.23.65",
];

const random = (max: number): number => Math.random() * max;

const create = (): HostDTO[] => {
  const data: HostDTO[] = [];
  for (let i = 0; i < 5; ++i) {
    data.push({
      ip: ips[i],
      os: i % 2 === 0 ? "Windows" : "Linux",
      hostnameLabel: "hostname" + i,
      customLabel: "custom" + i,
      services: [
        {
          ip: ips[i],
          port: random(65535).toFixed(0),
          banner: "Apache",
          firstDiscovered: new Date().toLocaleTimeString(),
          lastDiscovered: new Date().toLocaleTimeString(),
        },
        {
          ip: ips[i],
          port: random(65535).toFixed(0),
          banner: "Nginx",
          firstDiscovered: new Date().toLocaleTimeString(),
          lastDiscovered: new Date().toLocaleTimeString(),
        },
        {
          ip: ips[i],
          port: random(65535).toFixed(0),
          banner: "OpenSSH",
          firstDiscovered: new Date().toLocaleTimeString(),
          lastDiscovered: new Date().toLocaleTimeString(),
        },
      ],
      firstDiscovered: new Date().toLocaleTimeString(),
      lastDiscovered: new Date().toLocaleTimeString() + "",
    });
  }
  return data;
};

const data: HostDTO[] = create();

export function sampleData(): ScanDTO {
  const cloned = [...data];
  vandalize(cloned);
  return new ScanDTO(cloned);

  function vandalize(data: HostDTO[]): void {
    shuffle(data);
    data.forEach((host) => shuffle(host.services));

    if (Math.random() > 0.5) {
      // Remove some services and hosts.;
      data.pop();
      data.pop();

      data[0].services.pop();
      data[1].services.pop();
      data[2].services.pop();
    }
  }
  function shuffle(array: any[]): void {
    let currentIndex = array.length;
    // While there remain elements to shuffle...
    while (currentIndex !== 0) {
      // Pick a remaining element...
      let randomIndex = Math.floor(Math.random() * currentIndex);
      currentIndex--;
      // And swap it with the current element.
      [array[currentIndex], array[randomIndex]] = [
        array[randomIndex],
        array[currentIndex],
      ];
    }
  }
}
