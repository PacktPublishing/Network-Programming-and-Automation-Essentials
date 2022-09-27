#!/usr/bin/python3

from scrapli.driver import GenericDriver

MY_DEVICE = {
    "host": "10.0.4.1",
    "auth_username": "user",
    "auth_password": "pw",
    "auth_strict_key": False,
}

with GenericDriver(**MY_DEVICE) as conn:
    result = conn.send_command("show version")
print(result.result)
