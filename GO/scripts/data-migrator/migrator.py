from cassandra.cluster import Cluster
from cassandra.decoder import dict_factory
from config import CASS_CONFIG, MONGO_CONFIG
from pymongo import Connection


class Cassandra(object):

	def __init__(self):
		"""
		Cassandara Setup
		"""
		cluster = Cluster(CASS_CONFIG['cluster'], protocol_version=CASS_CONFIG['protocol_version'])
		self.session = cluster.connect(CASS_CONFIG['keyspace'])
		self.session.row_factory = dict_factory

	def get_records(self):
		"""
		fetches the records from the specified coloumnfamily
		"""
		orders_stmt = self.session.prepare('SELECT * FROM {}'.format(CASS_CONFIG['cache']))
		return self.session.execute(orders_stmt, [])

	def get_count(self):
		"""
		gets the counts of coloumnfamily
		"""
		orders_stmt = self.session.prepare('SELECT count(*) FROM {}'.format(CASS_CONFIG['cache']))
		return self.session.execute(orders_stmt, [])


def get_mongo_db():
	"""
	returns the Mongo collection
	"""
	client = Connection(MONGO_CONFIG['host'], MONGO_CONFIG['port'])
	db = client[MONGO_CONFIG['db']]
	if MONGO_CONFIG['enable_auth']:
		db.authenticate(MONGO_CONFIG['user'], MONGO_CONFIG['pass'])
	return db[MONGO_CONFIG['collection']]


if __name__ == '__main__':

	collection = get_mongo_db()
	CASS = Cassandra()
	cass_items = CASS.get_records()
	cass_count = CASS.get_count()
	missed_item = []
	event_ids = []

	for line in cass_items:
		line['entryTimestamp'] = line.pop('entrytimestamp')
		line['oppId'] = line.pop('opp_id')
		line['eventId'] = line.pop('key')
		line['accountId'] = line.pop('account_id')
		line['salesRep'] = line.pop('sales_rep')
		collection.insert(line)
		event_ids.append(line['eventId']) #adding event id to list for validation

	for item in collection.find():
		if item['eventId'] not in event_ids:
			missed_item.append(item['eventId'])

	mongo_count = collection.count()
	if cass_count[0]['count'] != mongo_count or missed_item:
		print {"Missing Records": missed_item}
	else:
		print {"mongo_count": mongo_count, "cass_count": cass_count[0]['count']}
		print "Data Migrated Succesfully"




	#collection.insert(data) # inserts the list of records into mongo
