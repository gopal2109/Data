# -*- coding: utf-8 -*-
import datetime
import pymongo

from bson.objectid import ObjectId

# CONFIG
MONGO_URL = 'mongodb://localhost:27017'
DB = 'configurations'
COLLECTION_NAME = 'regex'

# DB
mc = pymongo.MongoClient(MONGO_URL)
database = mc[DB]

# Name of user added/updated document
user = 'bhaw2019'


def get_db_collection(collection):
    return database[collection]

rules_collection = get_db_collection(COLLECTION_NAME)

# Rules for F5
f5 = [
    {
        "rule_name": "auth tacacs",
        "device_type": "f5",
        "rule_type": "line:startswith",
        "regex": "\\s*(\\S+)\\s+(\\S+).*$",
        "replace": "f5_replace_line_starts_with_util",
        "is_active": True,
        "is_function": True
    },
    {
        "rule_name": "configsync",
        "device_type": "f5",
        "rule_type": "line:startswith",
        "regex": "\\s*(\\S+)\\s+(\\S+).*$",
        "replace": "f5_replace_line_starts_with_util",
        "is_active": True,
        "is_function": True
    },
    {
        "rule_name": "snmpd",
        "device_type": "f5",
        "rule_type": "line:startswith",
        "regex": "\\s*(\\S+)\\s+(\\S+).*$",
        "replace": "f5_replace_line_starts_with_util",
        "is_active": True,
        "is_function": True
    },
    {
        "rule_name": "syslog",
        "device_type": "f5",
        "rule_type": "line:startswith",
        "regex": "\\s*(\\S+)\\s+(\\S+).*$",
        "replace": "f5_replace_line_starts_with_util",
        "is_active": True,
        "is_function": True
    },
    {
        "rule_name": "password replace",
        "device_type": "f5",
        "rule_type": "line:replace",
        "regex": "^.*password.*$",
        "replace": "password  <REMOVED>",
        "is_active": True,
        "is_function": False
    }
]


def get_max_rule_id():

    rule = list(rules_collection.find().sort([('rule_id', -1)]).limit(1))
    rule_id = rule[0]['rule_id']
    return rule_id


def add_rules():
    starts = get_max_rule_id() + 1
    for rule in f5:

        rule.update(
                {'_id': ObjectId(),
                 'rule_id': starts,
                 'created_by': user,
                 'updated_by': user,
                 'added_on': datetime.datetime.utcnow(),
                 'updated_on': datetime.datetime.utcnow()
                 }
        )
        rules_collection.insert(rule)

        starts += 1


def main():

    print('Adding rules for f5 device to database: {}, collection: {}'.format(DB, COLLECTION_NAME))
    add_rules()
    print('Done:)')


if __name__ == '__main__':

    main()
