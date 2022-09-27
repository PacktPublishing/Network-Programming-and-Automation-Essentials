#!/usr/bin/python3

from scrapli.driver import GenericDriver

TARGET_DEVICE = {
    "host": "10.0.4.1",
    "auth_username": "user",
    "auth_password": "pw",
    "auth_strict_key": False,
}

with GenericDriver(**TARGET_DEVICE) as conn:
    command_return = conn.send_command("show version")

print(command_return.result)
