#!/usr/bin/python3

from scrapli.driver import GenericDriver

TARGET = {
    "host": "10.0.4.1",
    "auth_username": "user",
    "auth_password": "pw",
    "auth_strict_key": False,
}

with GenericDriver(**TARGET) as conn:
    command_return = conn.send_command("uptime")

print(command_return.result)
