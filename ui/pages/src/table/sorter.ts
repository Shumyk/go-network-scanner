import { HostType } from "./dataTableObject";

export const sort = (types: HostType[]): HostType[] => {
  types
    .sort((a, b) => ipNumber(a.ip) - ipNumber(b.ip))
    .forEach((host) =>
      host.services.sort((a, b) => Number(a.port) - Number(b.port))
    );
  return types;

  function ipNumber(ip: string): number {
    return Number(
      ip
        .split(".")
        .map((part) => `00${part}`.slice(-3))
        .join("")
    );
  }
};
