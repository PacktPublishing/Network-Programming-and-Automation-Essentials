from jinja2 import Environment, FileSystemLoader
import yaml

env = Environment(loader=FileSystemLoader("."))
template = env.get_template("cisco_template_python.txt")

with open("router_definitions.yaml") as f:
    routers = yaml.safe_load(f)

for router in routers:
    router_conf = router["name"] + "_router_config.txt"
    with open(router_conf, "w") as f:
        f.write(template.render(router))
