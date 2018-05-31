"""
Script to deal with tasks for add/update number_of_segments in products
collection for FIREWALL type product
"""

from pymongo import MongoClient
from pymongo.errors import (AutoReconnect,
                            ConfigurationError,
                            ConnectionFailure
                            )

DB_CONFIG = {
    'URI': 'mongodb://USERNAME:PASSWORD@HOST_NAME:POST_NUMBER/DB_NAME'
}

# Add product name as key and number of segments as value
PRODUCTS_TO_UPDATE = {
    'Cisco ASA 5505 Firewall': 3,
    'Cisco ASA 5505 Unlimited (Orange Sticker)': 3,
    'Cisco ASA 5510 Sec+': 3,
    'Cisco ASA 5515': 9
}

COLLECTION_NAME = 'products'


class NisMongoClient(object):

    @classmethod
    def add_field_to_db_collection(cls):
        """
        Connect with db, find given collection for update
        Add new fields with default values in existing collection
        """
        try:
            # creating db connection.
            client = MongoClient(DB_CONFIG['URI'])
            db = client.get_default_database()

            if len(PRODUCTS_TO_UPDATE):
                print('Started updating documents')
                for product_name, segments in PRODUCTS_TO_UPDATE.items():
                    db[COLLECTION_NAME].update(
                        {
                            'product_name': product_name
                        },
                        {'$set': {'number_of_segments': segments,
                                  'is_multi_segment': True}
                        }
                    )
                print('Done :)')
            else:
                print('No product to update')
        except (AutoReconnect, ConfigurationError, ConnectionFailure) as e:
            print("Error:: "+str(e)+"")

NisMongoClient.add_field_to_db_collection()
