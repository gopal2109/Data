// This script lists the current aggrzone for a given list of devices

// USAGE from Mongo Shell command line (execute and quit):
// mongo nis -u <username> -p "<password>" --eval "var devices [<csv of device integers>]" list-aggr-zone-for-devices.js

// To run from within Mongo Shell, uncomment the devices variable declaration and include the list of devices in the list value
//var devices = []

var results = db.devices.find({device_id: {$in: devices}}, {device_id:1, aggr_zone:1, cabinet_id:1}).sort({device_id:1});
results.forEach( function(d) {
	var output = {"device_id":d.device_id, "cabinet_id":d.cabinet_id, "aggr_zone":d.aggr_zone};
	print (JSON.stringify(output));
});
