import subprocess

TARGET = "yahoo.com"

command = ["ping", "-c", "1", TARGET]
response = subprocess.call(command, stdout=subprocess.DEVNULL)

if response == 0:
    print(TARGET, "OK")
else:
    print(TARGET, "FAILED")
