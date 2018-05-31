# This script is used to Migrate Data from NIS collections to RPI collections

import logging
import pymongo
import time

MONGO_URL = 'mongodb://localhost:27017'
DB_SOURCE = 'nis'
DB_DESTINATION = 'rpis'

# DB
mc = pymongo.MongoClient(MONGO_URL)
nis_database = mc[DB_SOURCE]
rpis_database = mc[DB_DESTINATION]

def get_nis_db_collection(collection):
    return nis_database[collection]

def get_rpis_db_collection(collection):
    return rpis_database[collection]

log = logging.getLogger('load_devices_data')

def insert_devices(devices_collection):
    rpis_devices_list = []
    rpis_devices = {}
    provider = {}
    location ={}
    product = {}
    productCatalogDetails= {}
    offerServiceDetails= {}
    inventoryState = {}
    deviceService = {}
    core = {}
    devices = {}
    links = {}
    account = {}
    quote = {}
    automationEvent = {}
    deviceState = {}
    lastModified = {}
    created = {}
    log.info('Total count before insertion ... {}'.format(devices_collection.count()))
    nis_device_collection = get_nis_db_collection('devices')
    devices_list = nis_device_collection.find()
    for device in devices_list:
        rpis_devices['type'] = device.get('device_type')
        rpis_devices['macAddress'] = device.get('mac_address')
        rpis_devices['comment'] = device.get('comments')
        rpis_devices['deleted'] = device.get('is_deleted')
        provider['href'] = None
        provider['name'] = None
        provider['id'] = None
        rpis_devices['provider'] = provider
        location['aggrZone'] = device.get('aggr_zone')
        location['datacenterId'] = None
        location['datacenter'] = device.get('dc')
        location['cabinet'] = None
        location['cabinetStartingSpace'] = device.get('starting_space')
        rpis_devices['location'] = location
        productCatalogDetails['href'] = "productCatalogLink"
        productCatalogDetails['id'] = device.get('product_catalog_item_id')
        productCatalogDetails['offeringDescription'] = "48GB Dual Processor Hex Core Dedicated Server"
        offerServiceDetails['href'] = "offerlink"
        offerServiceDetails['productId'] = device.get('product_id')
        offerServiceDetails['productName'] = device.get('product_name')
        product['productCatalogDetails'] = productCatalogDetails
        product['offerServiceDetails'] = offerServiceDetails
        rpis_devices['product'] = product
        inventoryState['state'] = None
        inventoryState['userId'] = device.get('user_id')
        inventoryState['comment'] = None
        inventoryState['timestamp'] = time.strftime("%c")
        deviceService['href'] = None
        deviceService['id'] = None
        core['href'] = None
        core['id'] = None
        devices['deviceService'] = deviceService
        devices['core'] = core
        inventoryState['device'] = devices
        links['account'] = None
        links['device'] = None
        account['links'] = links
        account['accountId'] = None
        account['deviceId'] = device.get('device_id')
        inventoryState['account'] = account
        quote['href'] = None
        quote['id'] = None
        quote['salesPersonUserId'] = None
        quote['opportunity'] = None
        inventoryState['quote'] = quote
        automationEvent['href'] = ".../events/45bd38cb-f401-4873-a7a5-d5afa2fcc0cd"
        automationEvent['eventId'] = "45bd38cb-f401-4873-a7a5-d5afa2fcc0cd"
        inventoryState['automationEvent'] = automationEvent
        rpis_devices['inventoryState'] = inventoryState
        deviceState['state'] = None
        deviceState['comment'] = None
        deviceState['userId'] = device.get('user_id')
        deviceState['timestamp'] = time.strftime("%c")
        rpis_devices['deviceState'] = deviceState
        lastModified['userId'] = device.get('user_id')
        lastModified['timestamp'] = time.strftime("%c")
        rpis_devices['lastModified'] = lastModified
        created['userId'] = device.get('user_id')
        created['timestamp'] = time.strftime("%c")
        rpis_devices['created'] = created
        rpis_devices_list.append(rpis_devices)

    for value in rpis_devices_list:
        if '_id' in value:
            del value['_id']
        result = devices_collection.insert(value)
        if not result:
            raise Exception('devices not inserted ')
    log.info('DONE :-) Total count after insertion ... {}'.format(devices_collection.count()))


def main():
    devices_collection = get_rpis_db_collection('devices')
    log.info('Updating rpis devices database!')
    insert_devices(devices_collection)


if __name__ == '__main__':
    logging.basicConfig(level=logging.DEBUG)
    main()
