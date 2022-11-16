from pythonping import ping

TARGET = "yahoo.com"

response = ping(TARGET, count=1)
if response.success:
    print(TARGET, "OK")
else:
    print(TARGET, "FAILED")