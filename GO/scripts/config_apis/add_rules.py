import logging
import datetime
import pymongo

from bson.objectid import ObjectId

# CONFIG
MONGO_URL = 'mongodb://localhost:27017'
DB = 'configurations'
COLLECTION_NAME = 'regular_expressions'

# DB
mc = pymongo.MongoClient(MONGO_URL)
database = mc[DB]

# name of user added/updated document
user = 'bhaw2019'


def get_db_collection(collection):
    return database[collection]

_rules = []

# rule_id starts with
starts = 1000

log = logging.getLogger('config.add_rules')

ASA = [
    ('subsection:replace', '^(object-group.*network.*rackspace\-ips.*)', ''),
    ('subsection:replace', '^(object-group.*network.*rackspace\-netops.*)', ''),
    ('subsection:replace', '^(object-group.*network.*rackspace\-monitoring.*)', ''),
    ('subsection:replace', '^(object-group.*network.*rackspace\-rba.*)', ''),
    ('subsection:replace', '^(object-group.*network.*rackspace\-nest.*)', ''),
    ('subsection:replace', '^(object-group.*network.*rackspace\-infrastructure.*)', ''),
    ('subsection:replace', '^(object-group.*network.*intensive\-infrastructure.*)', ''),
    ('subsection:replace', '^(object-group.*network.*rackspace\-sitescope.*)', ''),
    ('subsection:replace', '^(object-group.*network.*rackspace\-segsupport.*)', ''),
    ('subsection:replace', '^(object-group.*network.*RACKSPACE\-BASTIONS.*)', ''),
    ('subsection:replace', '^(object-group.*network.*rackspace\-\S+.*)', ''),
    ('subsection:replace', '^(object-group.*network.*ALERT-LOGIC-IPs-OUTBOUND.*)', ''),
    ('subsection:replace', '^(object-group.*network.*ALERT-LOGIC-IPs.*)', ''),
    ('subsection:replace', '^(object-group.*network.*RAXGEN_v1_NET_NEXTGEN_BAST.*)', ''),
    ('subsection:replace', '^(object-group.*network.*RAXGEN_v1_NET_NEST.*)', ''),
    ('subsection:replace', '^(object-group.*network.*RACKSPACE-NEST.*)', ''),
    ('subsection:replace', '^(object-group.*network.*RACKSPACE-PATCHING.*)', ''),
    ('subsection:replace', '^(object-group.*network.*rs-zabbix.*)', ''),
    ('subsection:replace', '^(object-group.*network.*rackspace-denied-access.*)', ''),
    ('line:replace', '^enable.*password.*|^username .* privilege 15|^passwd.*|^username (smbproserv|enable_15|pix2) password.*', ''),
    ('line:replace', '^logging host outside (\d+.\d+.\d+.\d+).*', 'asa_logging_host_outside', 'IS_FUNCTION' ),
    ('subsection:replace', '^aaa-server.*\(outside\) host 10\.\d+\.\d+\.\d+', 'asa_aaa_server_filter_for_outside_host', 'IS_FUNCTION'),
    ('line:replace', '^snmp-server host outside (\d+.\d+.\d+.\d+) poll.*', 'snmp-server host outside <REMOVED>'),
    ('line:replace', '^snmp-server host outside 64.39.1.231 poll.*|^snmp-server host outside 64.39.1.234 poll.*|^snmp-server host outside 72.3.130.47 poll.*', 'asa_snmp_server_host_outside_rsvoyence', 'IS_FUNCTION'),
    ('line:replace', '^snmp-server host outside 72.3.130.48 poll.*|^snmp-server host outside 72.32.192.147 poll.*|^snmp-server host outside 173.203.5.102 poll.*', 'asa_snmp_server_host_outside_rsvoyence', 'IS_FUNCTION'),
    ('line:replace', '^snmp-server host outside.*rsvoyence.*|^snmp-server host outside 74.205.2.142 poll.*|^snmp-server host outside 10.\d+.109.(68|70|72|74|76) poll.*', 'asa_snmp_server_host_outside_rsvoyence', 'IS_FUNCTION'),
    ('line:replace', '^.*snmp-server community rsvoyence.*', '.*snmp-server community <REMOVED>'),
    ('line:replace', '^.*ssh\s+(\d+\.\d+\.\d+\.\d+)\s+(\d+\.\d+\.\d+\.\d+)\s+(\S+).*', 'asa_ssh_outside_remove', 'IS_FUNCTION'),
    ('line:replace', 'name\s+\d+\.\d+\.\d+\.\d+\s+(hybridControl.*|RCAutomation.*| CloudIP|.*hybrid.*|.*cloud.*|.*rackconnect.*|.*rackspace.*)',
     "name <REMOVED> {0} description"),
    ('line:replace', 'http\s+(\d+\.\d+\.\d+\.\d+)\s+(\S+).*$', 'http <REMOVED> {1}'),
    ('line:replace', 'username (smbproserv|enable_15|pix2|nglab-admin) attributes.*', ''),
    ('line:replace', '^ntp server.*', 'ntp server <REMOVED> source outside'),
    ('line:replace', 'failover key.*', 'failover key <REMOVED>'),
    ('line:replace', '^tftp-server outside.*|^Cryptochecksum.*', ''),
    ('subsection:replace', '^(object-group.*network.*RAXGEN_v1_NET_NETSEC-INFRA.*)', ''),
    ('subsection:replace', '^aaa-server.*(rsa_radius|rsa_radius2).* ', 'asa_aaa_server_filter', 'IS_FUNCTION'),
    ('subsection:replace', '^tunnel-group.*', 'asa_tunnel_group_filter', 'IS_FUNCTION')
]


BROCADE = [
    ('line:replace', r'^\s*aaa\s+(\S+).*$', ''),
    ('line:replace', r'^\s*ip\s+route\s+(\d+\.\d+\.\d+\.\d+)\s+(\d+\.\d+\.\d+\.\d+)\s+(\d+\.\d+\.\d+\.\d+).*$', 'brocade_ip_route', 'IS_FUNCTION'),
    ('line:replace', r'^\s*username\s+(\S+).*$', 'username <REMOVED>'),
    ('line:replace', r'^\s*enable\s+(\S+)\s+\d+.*$', ''),
    ('line:replace', r'^\s*tacacs-server\s+(\S+)\s+\d+.*$', 'tacacs-server <REMOVED>'),
    ('line:replace', r'^\s*tacacs-server\s+key\s+\d+\s+(\S+).*$', 'tacacs-server key <REMOVED>'),
    ('line:replace', r'^\s*snmp-server.*$', 'snmp-server <REMOVED>'),
    ('line:replace', r'^\s*sntp\s+server.*$', 'sntp server <REMOVED>'),
    ('line:replace', r'crypto-ssl certificate generate secret_data*', ''),
    ('line:replace', r'.*super-user-password.*$', 'super-user-password <REMOVED>'),
    ('section:replace', (r'---- BEGIN SSH2 PUBLIC KEY ----', r'---- END SSH2 PUBLIC KEY ----'), 'brocade_reflect', 'IS_FUNCTION'),
]


JUNIPER = [
    ('line:replace', '^set system root-authentication encrypted-password (\S+)$', 'set system root-authentication encrypted-password <REMOVED>'), # rule 1
    ('line:replace', '^set system tacplus.*$', 'set system tacplus <REMOVED>'), # rule 2
    ('line:replace', '^set system login.*$', 'set system login <REMOVED>'), # rule 3
    ('line:replace', '^set system ntp server .*$', 'set system ntp server <REMOVED>'), # rule 4
    ('line:replace', '^set snmp community (\w+) .*$', 'set snmp community <REMOVED>'), # rule 5
    ('line:replace', '^set snmp trap-group RAXTRAPS categories (.*)$', 'set snmp trap-group <REMOVED> categories {0}'), # rule 6
    ('line:replace', '^set snmp trap-group RAXTRAPS targets .*$', 'set snmp trap-group <REMOVED>'),# rule 7
    ('line:replace', '^set system syslog host .*$', 'set system syslog host <REMOVED>'),
    ('line:replace', '^set system tacplus.*$', 'set system tacplus <REMOVED>'),
    ('line:replace', '^set policy-options prefix-list (RAX.*) .*$', 'set policy-options prefix-list {0} <REMOVED>'), # rule 8
    ('line:replace', '^set policy-options prefix-list (RS\-MGM.*) .*$', 'set policy-options prefix-list {0} <REMOVED>'), # rule 9
    ('line:replace', '^set security address-book global address-set (RACKSPACE\-IPs) address .*$', 'set security address-book global address-set {0} address <REMOVED>'),
    ('line:replace', '^set security address-book global address-set (RACKSPACE\-NETOPS) address .*$', 'set security address-book global address-set {0} address <REMOVED>'),
    ('line:replace', '^set security address-book global address-set (RACKSPACE\-MONITORING) address .*$', 'set security address-book global address-set {0} address <REMOVED>'),
    ('line:replace', '^set security address-book global address-set (RACKSPACE\-RBA) address .*$', 'set security address-book global address-set {0} address <REMOVED>'),
    ('line:replace', '^set security address-book global address-set (RACKSPACE\-SITESCOPE) address .*$', 'set security address-book global address-set {0} address <REMOVED>'),
    ('line:replace', '^set security address-book global address-set (INTENSIVE\-INFRASTRUCTURE) address .*$', 'set security address-book global address-set {0} address <REMOVED>'),
    ('line:replace', '^set security address-book global address-set (rackspace\-infrastructure) address .*$', 'set security address-book global address-set {0} address <REMOVED>'),
    ('line:replace', '^set security address-book global address-set (rackspace\-sitescope) address .*$', 'set security address-book global address-set {0} address <REMOVED>'),
    ('line:replace', '^set security address-book global address-set (rackspace\-segsupport) address .*$', 'set security address-book global address-set {0} address <REMOVED>'),
    ('line:replace', '^set security address-book global address-set (RACKSPACE\-BASTIONS) address .*$', 'set security address-book global address-set {0} address <REMOVED>'),
    ('line:replace', '^set security address-book global address-set (rackspace\-\S+) address .*$', 'set security address-book global address-set {0} address <REMOVED>'),
    ('line:replace', '^set security address-book global address-set (ALERT\-LOGIC\-IPs\-OUTBOUND) address .*$', 'set security address-book global address-set {0} address <REMOVED>'),
    ('line:replace', '^set security address-book global address-set (ALERT\-LOGIC\-IPs) address .*$', 'set security address-book global address-set {0} address <REMOVED>'),
    ('line:replace', '^set security address-book global address-set (rs\-zabbix) address .*$', 'set security address-book global address-set {0} address <REMOVED>'),
    ('line:replace', '^set security address-book global address-set (rackspace\-denied\-access) address .*$', 'set security address-book global address-set {0} address <REMOVED>'),
    ('line:replace', '^(.*)\s+(\d+\.\d+\.\d+\.\d+\/\d+) (\d+\.\d+\.\d+\.\d+\/\d+)(.*)?$', 'juniper_rex_inner_fun1', 'IS_FUNCTION'),
    ('line:replace', '^(.*)\s+(\d+\.\d+\.\d+\.\d+\/\d+) next\-hop (\d+\.\d+\.\d+\.\d+).*$', 'juniper_rex_inner_fun2', 'IS_FUNCTION'),
    ('line:replace', '^(.*)\s+(\d+\.\d+\.\d+\.\d+\/\d+) next\-hop (\S+).*$', 'juniper_rex_inner_fun3', 'IS_FUNCTION'),
    ('line:replace', '^(.*)\s+(\d+\.\d+\.\d+\.\d+\/\d+)(.*)$', 'juniper_rex_inner_fun4', 'IS_FUNCTION'),
    ('line:replace', '^(.*)\s+(\d+\.\d+\.\d+\.\d+)(.*)?$', 'juniper_rex_inner_fun5', 'IS_FUNCTION')
]


# formatting mongodb document for ASA
for asa in ASA:

    try:
        asa[3]
        is_function = True
    except:
        is_function = False

    rule = {
        '_id': ObjectId(),
        'rule_id': starts,
        'rule_type': asa[0],
        'rule_name': '',
        'device_type': 'ASA',
        'regex': asa[1],
        'replace': asa[2],
        'is_active': True,
        'is_function': is_function,
        'created_by': user,
        'updated_by': user,
        'added_on': datetime.datetime.utcnow(),
        'updated_on': datetime.datetime.utcnow()
    }
    _rules.append(rule)
    starts += 1

# formatting mongodb document for JUNIPER device
for juni in JUNIPER:

    try:
        juni[3]
        is_function = True
    except:
        is_function = False

    rule = {
        '_id': ObjectId(),
        'rule_id': starts,
        'rule_type': juni[0],
        'rule_name': '',
        'device_type': 'JUNIPER',
        'regex': juni[1],
        'replace': juni[2],
        'is_active': True,
        'is_function': is_function,
        'created_by': user,
        'updated_by': user,
        'added_on': datetime.datetime.utcnow(),
        'updated_on': datetime.datetime.utcnow()
    }
    _rules.append(rule)
    starts += 1

# formatting mongodb document for BROCADE device
for broc in BROCADE:

    try:
        broc[3]
        is_function = True
    except:
        is_function = False

    rule = {
        '_id': ObjectId(),
        'rule_id': starts,
        'rule_type': broc[0],
        'rule_name': '',
        'device_type': 'BROCADE',
        'regex': broc[1],
        'replace': broc[2],
        'is_function': is_function,
        'is_active': True,
        'created_by': user,
        'updated_by': user,
        'added_on': datetime.datetime.utcnow(),
        'updated_on': datetime.datetime.utcnow()
    }
    _rules.append(rule)
    starts += 1

logging.info('Total number of rules: {}'.format(_rules))


# add rules to database
def add_rules(rules_collection):
    for rule in _rules:
        log.info('adding rule: {}'.format(rule['rule_id']))

        rules_collection.insert(rule)

    log.info('Done')


def main():

    log.info('Adding rules to database: {}, collection: {}'.format(DB, COLLECTION_NAME))
    rules_collection = get_db_collection(COLLECTION_NAME)
    add_rules(rules_collection)
    print('Done :)')


if __name__ == '__main__':

    logging.basicConfig(level=logging.DEBUG)
    main()
