const disableKeyReload = (e: KeyboardEvent) => {
  if (isF5(e) || isCtrlR(e) || isCtrlShiftR(e) || isCmdR(e) || isCmdShiftR(e)) {
    e.preventDefault();
    return;
  }

  function isF5(event: KeyboardEvent): boolean {
    return event.key === "F5";
  }
  function isCtrlR(event: KeyboardEvent): boolean {
    return event.ctrlKey && event.key === "r";
  }
  function isCtrlShiftR(event: KeyboardEvent): boolean {
    return event.ctrlKey && event.shiftKey && event.key === "R";
  }
  function isCmdR(event: KeyboardEvent): boolean {
    return event.metaKey && event.key === "r";
  }
  function isCmdShiftR(event: KeyboardEvent): boolean {
    return event.metaKey && event.shiftKey && event.key === "R";
  }
};

const setupEvents = () => {
  window.addEventListener("keydown", disableKeyReload);
  return () => window.removeEventListener("keydown", disableKeyReload);
};

export default setupEvents;
