import netmiko 
 
host = { 
    "host": "10.0.4.1", 
    "username": "netlab", 
    "password": "netlab", 
    "device_type": "linux_ssh", 
} 
 
with netmiko.ConnectHandler(**host) as netcon: 
    output = netcon.send_command("uptime") 

print(output) 
