# This script is used to load thresholds data in the collection.

import logging
import pymongo

MONGO_URL = 'mongodb://localhost:27017'
DB = 'rpis'

# DB
mc = pymongo.MongoClient(MONGO_URL)
database = mc[DB]


def get_db_collection(collection):
    return database[collection]

log = logging.getLogger('load_thresholds_data')


__thresholds = [
    {
        "offering": {
            "href": "product/catalog/offering/111",
            "offeringId": 111
        },
        "datacenterThresholds": [{
            "datacenterId": 12,
            "datacenterAbbreviation": "DFW3",
            "warning": 100,
            "critical": 45
        }, {
            "datacenterId": 13,
            "datacenterAbbreviation": "IAD3",
            "warning": 75,
            "critical": 30
        }, {
            "datacenterId": 17,
            "datacenterAbbreviation": "LON5",
            "warning": 60,
            "critical": 50
        }]
    },
    {
        "offering": {
            "href": "product/catalog/offering/967",
            "offeringId": 967
        },
        "datacenterThresholds": [{
            "datacenterId": 12,
            "datacenterAbbreviation": "DFW3",
            "warning": 100,
            "critical": 45
        }, {
            "datacenterId": 13,
            "datacenterAbbreviation": "IAD3",
            "warning": 75,
            "critical": 30
        }, {
            "datacenterId": 17,
            "datacenterAbbreviation": "LON5",
            "warning": 60,
            "critical": 50
        }]
    },
    {
        "offering": {
            "href": "product/catalog/offering/968",
            "offeringId": 968
        },
        "datacenterThresholds": [{
            "datacenterId": 12,
            "datacenterAbbreviation": "DFW3",
            "warning": 100,
            "critical": 45
        }, {
            "datacenterId": 13,
            "datacenterAbbreviation": "IAD3",
            "warning": 75,
            "critical": 30
        }, {
            "datacenterId": 17,
            "datacenterAbbreviation": "LON5",
            "warning": 60,
            "critical": 50
        }]
    },
    {
        "offering": {
            "href": "product/catalog/offering/969",
            "offeringId": 969
        },
        "datacenterThresholds": [{
            "datacenterId": 12,
            "datacenterAbbreviation": "DFW3",
            "warning": 100,
            "critical": 45
        }, {
            "datacenterId": 13,
            "datacenterAbbreviation": "IAD3",
            "warning": 75,
            "critical": 30
        }, {
            "datacenterId": 17,
            "datacenterAbbreviation": "LON5",
            "warning": 60,
            "critical": 50
        }]
    },
    {
        "offering": {
            "href": "product/catalog/offering/970",
            "offeringId": 970
        },
        "datacenterThresholds": [{
            "datacenterId": 12,
            "datacenterAbbreviation": "DFW3",
            "warning": 100,
            "critical": 45
        }, {
            "datacenterId": 13,
            "datacenterAbbreviation": "IAD3",
            "warning": 75,
            "critical": 30
        }, {
            "datacenterId": 17,
            "datacenterAbbreviation": "LON5",
            "warning": 60,
            "critical": 50
        }]
    },
    {
        "offering": {
            "href": "product/catalog/offering/2629",
            "offeringId": 2629
        },
        "datacenterThresholds": [{
            "datacenterId": 12,
            "datacenterAbbreviation": "DFW3",
            "warning": 100,
            "critical": 45
        }, {
            "datacenterId": 13,
            "datacenterAbbreviation": "IAD3",
            "warning": 75,
            "critical": 30
        }, {
            "datacenterId": 17,
            "datacenterAbbreviation": "LON5",
            "warning": 60,
            "critical": 50
        }]
    },
    {
        "offering": {
            "href": "product/catalog/offering/2632",
            "offeringId": 2632
        },
        "datacenterThresholds": [{
            "datacenterId": 12,
            "datacenterAbbreviation": "DFW3",
            "warning": 100,
            "critical": 45
        }, {
            "datacenterId": 13,
            "datacenterAbbreviation": "IAD3",
            "warning": 75,
            "critical": 30
        }, {
            "datacenterId": 17,
            "datacenterAbbreviation": "LON5",
            "warning": 60,
            "critical": 50
        }]
    },
    {
        "offering": {
            "href": "product/catalog/offering/2634",
            "offeringId": 2634
        },
        "datacenterThresholds": [{
            "datacenterId": 12,
            "datacenterAbbreviation": "DFW3",
            "warning": 100,
            "critical": 45
        }, {
            "datacenterId": 13,
            "datacenterAbbreviation": "IAD3",
            "warning": 75,
            "critical": 30
        }, {
            "datacenterId": 17,
            "datacenterAbbreviation": "LON5",
            "warning": 60,
            "critical": 50
        }]
    },
    {
        "offering": {
            "href": "product/catalog/offering/3025",
            "offeringId": 3025
        },
        "datacenterThresholds": [{
            "datacenterId": 12,
            "datacenterAbbreviation": "DFW3",
            "warning": 100,
            "critical": 45
        }, {
            "datacenterId": 13,
            "datacenterAbbreviation": "IAD3",
            "warning": 75,
            "critical": 30
        }, {
            "datacenterId": 17,
            "datacenterAbbreviation": "LON5",
            "warning": 60,
            "critical": 50
        }]
    },
    {
        "offering": {
            "href": "product/catalog/offering/3067",
            "offeringId": 3067
        },
        "datacenterThresholds": [{
            "datacenterId": 12,
            "datacenterAbbreviation": "DFW3",
            "warning": 100,
            "critical": 45
        }, {
            "datacenterId": 13,
            "datacenterAbbreviation": "IAD3",
            "warning": 75,
            "critical": 30
        }, {
            "datacenterId": 17,
            "datacenterAbbreviation": "LON5",
            "warning": 60,
            "critical": 50
        }]
    },
    {
        "offering": {
            "href": "product/catalog/offering/3071",
            "offeringId": 3071
        },
        "datacenterThresholds": [{
            "datacenterId": 12,
            "datacenterAbbreviation": "DFW3",
            "warning": 100,
            "critical": 45
        }, {
            "datacenterId": 13,
            "datacenterAbbreviation": "IAD3",
            "warning": 75,
            "critical": 30
        }, {
            "datacenterId": 17,
            "datacenterAbbreviation": "LON5",
            "warning": 60,
            "critical": 50
        }]
    },
    {
        "offering": {
            "href": "product/catalog/offering/2712",
            "offeringId": 2712
        },
        "datacenterThresholds": [{
            "datacenterId": 12,
            "datacenterAbbreviation": "DFW3",
            "warning": 100,
            "critical": 45
        }, {
            "datacenterId": 13,
            "datacenterAbbreviation": "IAD3",
            "warning": 75,
            "critical": 30
        }, {
            "datacenterId": 17,
            "datacenterAbbreviation": "LON5",
            "warning": 60,
            "critical": 50
        }]
    },
    {
        "offering": {
            "href": "product/catalog/offering/2002",
            "offeringId": 2002
        },
        "datacenterThresholds": [{
            "datacenterId": 12,
            "datacenterAbbreviation": "DFW3",
            "warning": 100,
            "critical": 45
        }, {
            "datacenterId": 13,
            "datacenterAbbreviation": "IAD3",
            "warning": 75,
            "critical": 30
        }, {
            "datacenterId": 17,
            "datacenterAbbreviation": "LON5",
            "warning": 60,
            "critical": 50
        }]
    },
    {
        "offering": {
            "href": "product/catalog/offering/2636",
            "offeringId": 2636
        },
        "datacenterThresholds": [{
            "datacenterId": 12,
            "datacenterAbbreviation": "DFW3",
            "warning": 100,
            "critical": 45
        }, {
            "datacenterId": 13,
            "datacenterAbbreviation": "IAD3",
            "warning": 75,
            "critical": 30
        }, {
            "datacenterId": 17,
            "datacenterAbbreviation": "LON5",
            "warning": 60,
            "critical": 50
        }]
    },
    {
        "offering": {
            "href": "product/catalog/offering/3195",
            "offeringId": 3195
        },
        "datacenterThresholds": [{
            "datacenterId": 12,
            "datacenterAbbreviation": "DFW3",
            "warning": 100,
            "critical": 45
        }, {
            "datacenterId": 13,
            "datacenterAbbreviation": "IAD3",
            "warning": 75,
            "critical": 30
        }, {
            "datacenterId": 17,
            "datacenterAbbreviation": "LON5",
            "warning": 60,
            "critical": 50
        }]
    }
]


def insert_thresholds(thresholds_collection):
    log.info('Total count before insertion ... {}'.format(thresholds_collection.count()))
    for threshold in __thresholds:
        log.info('Inserting Thresholds for {}'.format(threshold['offering']['offeringId']))
        result = thresholds_collection.insert(threshold)
        if not result:
            raise Exception('Thresholds not inserted for {}'.format(threshold['offering']['offeringId']))
    log.info('DONE :-) Total count after insertion ... {}'.format(thresholds_collection.count()))


def main():
    thresholds_collection = get_db_collection('thresholds')
    log.info('Updating thresholds database!')
    insert_thresholds(thresholds_collection)


if __name__ == '__main__':
    logging.basicConfig(level=logging.DEBUG)
    main()
