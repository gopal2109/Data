# This script is used to load inventoryState data in the collection.

import logging
import pymongo

MONGO_URL = 'mongodb://localhost:27017'
DB = 'rpis'

# DB
mc = pymongo.MongoClient(MONGO_URL)
database = mc[DB]


def get_db_collection(collection):
    return database[collection]


log = logging.getLogger('load_inventory_state_data')

__inventoyState = [
    {
        "state": "AVAILABLE",
        "userId": "automation",
        "comment": "IIR Complete. Device available for sale.",
        "timestamp": "1/10/2016T13:00:00Z",
        "automationEvent": {
            "href": ".../events/45bd38cb-f401-4873-a7a5-d5afa2fcc0cd",
            "eventId": "45bd38cb-f401-4873-a7a5-d5afa2fcc0cd"
        }
    }, {
        "state": "DECOMMISSIONED",
        "userId": "automation",
        "comment": "Starting Decom in the Rack process for this device.",
        "timestamp": "1/9/2016T11:11:11Z",
        "automationEvent": {
            "href": ".../events/45bd38cb-f401-4873-a7a5-d5afa2fcc0cd",
            "eventId": "45bd38cb-f401-4873-a7a5-d5afa2fcc0cd"
        }
    }, {
        "state": "ALLOCATED",
        "userId": "automation",
        "comment": "TEST",
        "timestamp": "1/4/2016T11:00:00Z",
        "device": {
            "deviceService": {
                "href": "deviceServiceLink",
                "id": "12312312312"
            },
            "core": {
                "href": "core.rackspace.com/device/12312312312",
                "id": "12312312312"
            }
        },
        "account": {
            "href": "core/account/612345",
            "accountId": 612345
        },
        "quote": {
            "href": ".../quotes/9898989",
            "id": "9898989",
            "salesPersonUserId": "jim1234",
            "opportunity": "9898989898"
        },
        "automationEvent": {
            "href": ".../events/45bd38cb-f401-4873-a7a5-d5afa2fcc0cd",
            "eventId": "45bd38cb-f401-4873-a7a5-d5afa2fcc0cd"
        }
    }, {
        "state": "MAINTENACE",
        "userId": "joe2345",
        "comment": "Noticed blinking drive on device.",
        "timestamp": "1/3/2016T08:00:00Z",
        "automationEvent": {
            "href": ".../events/45bd38cb-f401-4873-a7a5-d5afa2fcc0cd",
            "eventId": "45bd38cb-f401-4873-a7a5-d5afa2fcc0cd"
        }
    }, {
        "state": "SUSPENDED",
        "userId": "automation",
        "comment": "RAID battery failure",
        "timestamp": "1/2/2016T22:00:00Z",
        "automationEvent": {
            "href": ".../events/45bd38cb-f401-4873-a7a5-d5afa2fcc0cd",
            "eventId": "45bd38cb-f401-4873-a7a5-d5afa2fcc0cd"
        }
    }, {
        "state": "AVAILABLE",
        "userId": "automation",
        "comment": "IIR complete. Marking device as available for inventory.",
        "timestamp": "1/1/2016T11:12:00Z",
        "automationEvent": {
            "href": ".../events/45bd38cb-f401-4873-a7a5-d5afa2fcc0cd",
            "eventId": "45bd38cb-f401-4873-a7a5-d5afa2fcc0cd"
        }
    }, {
        "state": "NEW",
        "userId": "mercury",
        "comment": "TEST",
        "timestamp": "1/1/2016T11:11:00Z"
    }
]


def insert_inventory_state_collection(inventory_state_collection):
    for data in __inventoyState:
        result = inventory_state_collection.insert(data)
        if not result:
            raise Exception('inventoryState data not inserted ')
    log.info('DONE :-) Total count after insertion ... {}'.format(inventory_state_collection.count()))


def main():
    inventory_state_collection = get_db_collection('inventoryState')
    log.info('Updating rpis inventoryState database!')
    insert_inventory_state_collection(inventory_state_collection)


if __name__ == '__main__':
    logging.basicConfig(level=logging.DEBUG)
    main()
