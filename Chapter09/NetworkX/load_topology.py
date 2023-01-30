import networkx as nx
import yaml

G = nx.Graph()
devices = {}

with open("topology.yaml", "r") as file:
    yfile = yaml.safe_load(file)

for i, x in enumerate(yfile["devices"]):
    devices[x] = i
    G.add_node(i, name=x)
for link in yfile["links"]:
    G.add_edge(devices[link[0]], devices[link[1]])
