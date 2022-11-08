from pysnmp.hlapi import *

snmpIt = getCmd(SnmpEngine(),
                  CommunityData('public'),
                  UdpTransportTarget(('10.0.4.1', 161)),
                  ContextData(),
                  ObjectType(ObjectIdentity('SNMPv2-MIB', 'sysUpTime', 0)))

errEngine, errAgent, errorIndex, vars = next(snmpIt)

if errEngine:
    print("Got engine error:", errEngine)
elif errAgent:
    print("Got agent error:", errAgent.prettyPrint())
else:
    for var in vars:
        print(' = '.join([x.prettyPrint() for x in var]))
