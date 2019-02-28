'use strict';

const Device = require('tplink-smarthome-simulator').Device;

const device = new Device({
    model: 'hs110',
    port: 9999,
    responseDelay: 100
});

device.start();
