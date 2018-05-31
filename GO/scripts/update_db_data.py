"""
Script to updata a device record for NWT-420
"""

from pymongo import MongoClient
from pymongo.errors import (AutoReconnect,
                            ConfigurationError,
                            ConnectionFailure
                            )

DB_CONFIG = {
    'URI': 'mongodb://USERNAME:PASSWORD@HOST_NAME:PORT_NUMBER/DB_NAME'
}

# Add product name as key and number of segments as value
# format
# { <device_id>: (field, value) }

DATA_TO_UPDATE = {
    ('device_id', 742082): {"aggr_zone": "LON3:Public:Zone351-2"},
}

COLLECTION_NAME = 'devices'


class NisMongoClient(object):

    @classmethod
    def update_collection_data(cls):
        """
        Connect with db, find given collection for update
        Add new fields with default values in existing collection
        """
        try:
            # creating db connection.
            client = MongoClient(DB_CONFIG['URI'])
            db = client.get_default_database()
            collection = db[COLLECTION_NAME]
            
            if len(DATA_TO_UPDATE):
                print('Updating...')
                for _filter, data_to_update in DATA_TO_UPDATE.items():
                    print("Updating device: ", _filter)
                    collection.update(
                        {
                            _filter[0]: _filter[1]
                        },
                        {
                            '$set': data_to_update
                        }
                    )
                print('Updated')
            else:
                print('No data to update')
        except (AutoReconnect, ConfigurationError, ConnectionFailure) as e:
            print("Error:: "+str(e)+"")

NisMongoClient.update_collection_data()
