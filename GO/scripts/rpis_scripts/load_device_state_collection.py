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


log = logging.getLogger('load_device_state_data')

__deviceState = [
    {
        "state": "TEST",
        "comment": "Test comment",
        "timestamp": "1/1/1111T11:11:11Z"
    },
    {
        "state": "PRODUCTION",
        "comment": "Prod comment",
        "timestamp": "1/1/1111T11:11:11Z"
    },
    {
        "state": "PRE-PRODUCTION",
        "comment": "Pre-Prod comment",
        "timestamp": "1/1/1111T11:11:11Z"
    },
    {
        "state": "TEST",
        "comment": "Test comment",
        "timestamp": "1/1/1111T11:11:11Z"
    },
    {
        "state": "PRODUCTION",
        "comment": "Prod comment",
        "timestamp": "1/1/1111T11:11:11Z"
    },
    {
        "state": "PRE-PRODUCTION",
        "comment": "Pre-Prod comment",
        "timestamp": "1/1/1111T11:11:11Z"
    }
]


def insert_device_state_collection(device_state_collection):
    for device in __deviceState:
        result = device_state_collection.insert(device)
        if not result:
            raise Exception('deviceState data not inserted ')
    log.info('DONE :-) Total count after insertion ... {}'.format(device_state_collection.count()))


def main():
    device_state_collection = get_db_collection('deviceState')
    log.info('Updating rpis deviceState database!')
    insert_device_state_collection(device_state_collection)


if __name__ == '__main__':
    logging.basicConfig(level=logging.DEBUG)
    main()
