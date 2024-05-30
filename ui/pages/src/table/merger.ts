import {
  FOUND,
  HostType,
  LOST,
  OFFLINE,
  ONLINE,
  ServiceType,
  TypeStatus,
} from "./dataTableObject";
import { HostDTO, ServiceDTO } from "./dataTransferObject";

// ## Business logic.
// Merge the old data with the new data.
// The old data is the data that was previously scanned.
// The new data is the data that was scanned now.
//
// The merging process is done in the following way:
// - Iterate through the newly scanned data to understand which hosts are new or existing.
// - Iterate through the existing data to understand which hosts are lost.
// - Transition the status of the entities depending on the current status.
function merge(oldData: HostType[], newData: HostDTO[]): HostType[] {
  const merged: HostType[] = [];
  // Iterate through the newly scanned data to understand which hosts are new or existing.
  for (const newHost of newData) processNewHost(newHost);
  // Iterate through the existing data to understand which hosts are lost.
  for (const prevHost of oldData) processLostHost(prevHost);
  console.info("Merged data: ", merged);
  return merged;

  // # Processing functions.
  function processNewHost(newHost: HostDTO) {
    const prevHost = findByIp(oldData, newHost);
    if (prevHost) {
      // Entity continue to be present, transiting to ONLINE state.
      const online: HostType = mergeServices(prevHost, newHost);
      online.status = existingPreviousTransition(prevHost.status);
      merged.push(online);
    } else {
      // New host is found, mark all its services as FOUND.
      merged.push(HostType.found(newHost));
    }
  }
  function processLostHost(prevHost: HostType) {
    const newHost = findByIp(newData, prevHost);
    // Entry was present before and it's present now. It was handled above.
    if (newHost) return;
    // Entry is lost, transiting to next state depending on the current status.
    markAs(prevHost, absentStatusTransition(prevHost.status));
    merged.push(prevHost);
  }
}

// Merge the old services with the new services.
// Mechanism is the same as for the hosts.
function mergeServices(prevHost: HostType, newHost: HostDTO): HostType {
  const merged: HostType = HostType.of(newHost);
  // Iterate through the newly scanned data to understand which services are new or existing.
  for (const newService of newHost.services) processNewService(newService);
  // Iterate through the existing services to understand which services are lost.
  for (const prevService of prevHost.services) processLostService(prevService);
  return merged;

  // # Processing functions.
  function processNewService(newService: ServiceDTO) {
    const prevService = findByPort(prevHost.services, newService);
    const newStatus: TypeStatus = presentStatusTransition(prevService);

    merged.services.push(ServiceType.of(newService, newStatus, prevService?.firstDiscovered));
  }
  function processLostService(prevService: ServiceType) {
    const newService = findByPort(newHost.services, prevService);
    // Entry was present before and it's present now. It was handled in processNewService(..).
    if (newService) return;
    // Entry not present in the new data. It was present and lost.
    prevService.status = absentStatusTransition(prevService.status);
    merged.services.push(prevService);
  }
}

// ## Status transition functions.
// Depending on the current status, decide to which status to transition to.
//
// Entity is found in current scan, decide to which status to transition to.
function presentStatusTransition(previous: any): TypeStatus {
  if (previous)
    return existingPreviousTransition(previous.status);
  // Previous entry is absent, so it's newly found entity.
  return FOUND;
}
// Entity was found previous scan, so it should transition:
//      LOST or OFFLINE  -> to FOUND.
//      else             -> to ONLINE.
function existingPreviousTransition(previous: TypeStatus): TypeStatus {
  if (previous === LOST) return FOUND;
  if (previous === OFFLINE) return FOUND;
  return ONLINE;
}
// Entity is absent in current scan, decide to which status to transition to.
function absentStatusTransition(status: TypeStatus): TypeStatus {
  // Currently, entity is present, so it should transition to LOST.
  if (status === FOUND) return LOST;
  else if (status === ONLINE) return LOST;
  // If it is absent already, keep OFFLINE status.
  else if (status === LOST) return OFFLINE;
  else if (status === OFFLINE) return OFFLINE;
  return status;
}

// ## Helper types.
type HasIP = { ip: string };
type HasPort = { port: string };

// ## Helper functions.
function findByIp<T extends HasIP, S extends HasIP>(
  hosts: T[],
  target: T | S
): T | undefined {
  return hosts.find((h) => h.ip === target.ip);
}
function findByPort<T extends HasPort, S extends HasPort>(
  services: T[],
  target: T | S
): T | undefined {
  return services.find((s) => s.port === target.port);
}
function markAs(h: HostType, status: TypeStatus): void {
  h.status = status;
  h.services.forEach((service) => (service.status = status));
}

export { merge };
