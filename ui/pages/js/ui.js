let count = document.getElementById("count").innerHTML;
console.log("Initial count: " + count);

function SetCount(number) {
    document.getElementById("count").innerHTML = number;
    console.log("New count: " + number);
    count = number;
}

function AddDevice(device) {
    const list = document.getElementById("scan-list");
    const newDevice = document.createElement("p");
    newDevice.innerHTML = device;
    list.appendChild(newDevice);
}