#!/usr/bin/env python
"""
Script to deal with tasks to create or for newly added collection
dev: Sachin Deep
Prerequisites:
- MongoDB with NIS db should up and running.
- Configuration (inventory_exclusion_config.py) for script
Usage: python  path_to_script (scripts/create_inventory_exclusion_script.py)
"""

import time

from pymongo import MongoClient
from inventory_exclusion_config import MONGO_CONNECTION_CONFIG, \
default_inventory_exclusions_document


class NisMongoClient(object):

    _instance = None

    def __new__(cls, *args, **kwargs):
        if not cls._instance:
            cls._instance = super(NisMongoClient, cls).__new__(
                                cls, *args, **kwargs)
        return cls._instance

    @classmethod
    def create_db_collection(cls, collection_name):
        """
        Create collection in db and insert document.
        Collections in MongoDB are created lazily so not actually performed
        any operations on the MongoDB server.
        Collections are created when the first document is inserted into them.
        """
        client = MongoClient(MONGO_CONNECTION_CONFIG['host'],
                             MONGO_CONNECTION_CONFIG['port'])
        db = client[MONGO_CONNECTION_CONFIG['db']]

        if MONGO_CONNECTION_CONFIG['enable_auth']:
            db.authenticate(MONGO_CONNECTION_CONFIG['user'],
                            MONGO_CONNECTION_CONFIG['pass'])

        if collection_name not in db.collection_names():
            collection_for_update = db[collection_name]
            print("Collection going to create- "+collection_name)
            default_inventory_exclusions_document['created_timestamp_utc'] = time.asctime()
            collection_for_update.insert(default_inventory_exclusions_document)
            print("Collection inserted with default document "+collection_name)
        else:
            print("Collection is already present.")

NisMongoClient.create_db_collection('inventory_exclusions')
