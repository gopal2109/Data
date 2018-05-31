"""
This script is used to add new key(s) in redis database.
Run it only if need to add a new key.
"""

import sys
import os
from redis import StrictRedis
from argparse import ArgumentParser

class RedisAddKeys(object):

    def __init__(self, host):
        self.__redis = StrictRedis(host)

        """
        add key and value in dictionary and add to records list

        [
            {
                'key': 'product_id',
                'value': 13809
            },
            {
                'key': 'device_type',
                'val': 'server'
            }
        ]

        """

        self.records = [
            {
                'key': 'multisegment',
                'value': {
                    'is_active': True,
                    'feature_data': {
                        'maximum_cloud_segments': 9,
                        'dedicated_segments': [
                            'FW-INSIDE',
                            'FW-DMZ'
                        ]
                    }
                }
            }
        ]

        self.add_records()

    def add_records(self):
        keys = []
        # adding entry
        for entry in self.records:
            self.__redis.set(entry['key'], entry['value'])
            keys.append(entry['key'])

        # displays added keys
        print('Printing added key(s)')
        result = self.__redis.mget(keys)
        print(result)


if __name__ == '__main__':
    parser = ArgumentParser(description='Redi key creater')
    parser.add_argument('-s', '--redis', help='redis host')
    args = parser.parse_args()
    rk = RedisAddKeys(host=args.redis)
    rk.add_records()
