#!/usr/bin/env python
"""
Configuration for script to deal with tasks for newly added collection
Script: scripts.create_inventory_exclusion_script
"""

MONGO_CONNECTION_CONFIG = {
        'db': 'nis',
        'host': 'localhost',
        'port': 27017,
        'user': 'nis',
        'pass': '',
        'enable_auth': False
}

default_inventory_exclusions_document = {
    "owner_email": "ai_team@rackspace.com",
    "value": {
              "datacenter": [],
              "aggr_zone": [],
              "cabinet": [],
              "device_product": []
             },
    "created_by": "first_time_default_user",
    "created_timestamp_utc": None,
    "modified_by": None,
    "modified_timestamp_utc": None
}
