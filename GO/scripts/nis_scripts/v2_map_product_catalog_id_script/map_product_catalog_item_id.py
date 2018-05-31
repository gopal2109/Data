import logging
import pymongo

# CONFIG

MONGO_URL = 'mongodb://localhost:27017'
DB = 'nis'

# DB
mc = pymongo.MongoClient(MONGO_URL)
database = mc[DB]

def get_db_collection(collection):
    return database[collection]


log = logging.getLogger('nis_api.add_offer_id')

__mapping = [
    {
        'product_catalog_item_id': 967,
        'product_catalog_item_name': 'Rapid Deployment - 12GB Single Processor Quad Core Dedicated Server',
        'product_id': '9253c89e-b535-4d4a-8fc8-0745035c4263'
    },
    {
        'product_catalog_item_id': 968,
        'product_catalog_item_name': 'Rapid Deployment - 24GB Single Processor Hex Core Dedicated Server',
        'product_id': '030b8f2c-e5fd-4779-9aab-f117f2004b55'
    },
    {
        'product_catalog_item_id': 969,
        'product_catalog_item_name': 'Rapid Deployment - 24GB Dual Processor Quad Core Dedicated Server',
        'product_id': 'b89bf6dd-42c4-457d-bdbc-9174ee7f3554'
    },
    {
        'product_catalog_item_id': 970,
        'product_catalog_item_name': 'Rapid Deployment - 48GB Dual Processor Hex Core Dedicated Server',
        'product_id': '902e1779-1356-4a95-85fa-b50001ef874c'
    },
    {
        'product_catalog_item_id': 971,
        'product_catalog_item_name': 'Rapid Deployment - 128GB Dual Processor Hex Core Dedicated Server',
        'product_id': '6afc88a2-2e22-4bde-8545-a027533ec594'
    },
    {
        'product_catalog_item_id': 2629,
        'product_catalog_item_name': 'Rapid Deployment - 32GB Single Processor Quad Core Dedicated Server',
        'product_id': '027a0992-7ca3-11e4-b116-123b93f75cba'
    },
    {
        'product_catalog_item_id': 2632,
        'product_catalog_item_name': 'Rapid Deployment - 32GB Single Processor Hex Core Dedicated Server',
        'product_id': '5f5fc742-7cbb-11e4-b116-123b93f75cba'
    },
    {
        'product_catalog_item_id': 2634,
        'product_catalog_item_name': 'Rapid Deployment - 128GB Dual Processor Quad Core Dedicated Server',
        'product_id': '5f5fc9a4-7cbb-11e4-b116-123b93f75cba'
    },
    {
        'product_catalog_item_id': 3025,
        'product_catalog_item_name': 'Rapid Deployment - 32GB Single Processor Hex Core Dedicated Server Haswell',
        'product_id': '3e0bfb6d-dc02-4145-a367-55ebc449a605'
    },
    {
        'product_catalog_item_id': 3067,
        'product_catalog_item_name': 'Rapid Deployment - 64GB Single Processor Octo Core Dedicated Server Haswell',
        'product_id': 'dca86a19-3033-4133-86d9-bb4afbf2acb7'
    },
    {
        'product_catalog_item_id': 3071,
        'product_catalog_item_name': 'Rapid Deployment - 128GB Dual Processor Octo Core Dedicated Server Haswell',
        'product_id': '8c12a932-6368-401d-a6a1-fee550962607'
    },
    {
        'product_catalog_item_id': 2712,
        'product_catalog_item_name': 'Rapid Deployment - Cisco ASA 5505 Unlimited',
        'product_id': '027a1004-7ca3-11e4-b116-123b93f75cba'
    },
    {
        'product_catalog_item_id': 2002,
        'product_catalog_item_name': 'Rapid Deployment - Cisco ASA 5505 Firewall Sec Plus',
        'product_id': 'e8cfa8ec-af7c-4d3e-9659-0fc6017ad8eb'
    },
    {
        'product_catalog_item_id': 2636,
        'product_catalog_item_name': 'Rapid Deployment - Cisco ASA 5510 Sec Plus Firewall',
        'product_id': 'e9164316-7ca3-11e4-b116-123b93f75cba'
    },
    {
        'product_catalog_item_id': 3195,
        'product_catalog_item_name': 'Rapid Deployment - Cisco ASA 5515 X Firewall',
        'product_id': '11054e52-4aaa-4dcc-afa4-b75195486bcf'
    }
]


def update_product(product_collection):
    for m in __mapping:
        log.info('Updating: {}'.format(m['product_id']))
        result = product_collection.update({'product_id': m['product_id']},
                                           {
                                               '$set': {
                                                   'product_catalog_item_id': m['product_catalog_item_id'],
                                                   'product_catalog_item_name':m['product_catalog_item_name']
                                               }
                                           })
        if not result['n'] == 1:
            raise Exception('product does not exist : {}'.format(m['product_id']))
    log.info('Done')


def update_devices(devices_collections):
    for m in __mapping:
        log.info('Updating records for product: {}'.format(m['product_id']))
        result = devices_collections.update({'product_id': m['product_id']},
                                            {
                                                '$set': {
                                                    'product_catalog_item_id': m['product_catalog_item_id']
                                                }
                                            }, multi=True)
        log.info('Matched {} documents, made {} updates'.format(result['n'],
                                                                result['nModified']))
    log.info('Done')


def main():
    product_collection = get_db_collection('products')
    devices_collection = get_db_collection('devices')

    log.info('Updating database!')

    update_product(product_collection)
    update_devices(devices_collection)


if __name__ == '__main__':
    logging.basicConfig(level=logging.DEBUG)
    main()

