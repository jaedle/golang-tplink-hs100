'use strict';

const simulator = require('tplink-smarthome-simulator');
const Device = simulator.Device;
// const UdpServer = simulator.UdpServer;

const device = new Device({
    model: 'hs110v2',
    address: '127.0.0.1',
    port: 9999,
    responseDelay: 0
});
device.start();

// This uses a hardcoded mac and deviceId.
// devices.push(new Device({ model: 'hs100', data: { mac: '50:c7:bf:8f:58:18', deviceId: '12345' } }));

// UdpServer.start();