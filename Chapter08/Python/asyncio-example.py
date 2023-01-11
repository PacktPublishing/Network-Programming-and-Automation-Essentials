from pythonping import ping
import asyncio

TARGETS = ["yahoo.com", "google.com", "cisco.com", "cern.ch"]


async def myping(host):
    response = await ping(host)
    if response.success:
        print("%s OK, latency is %.2fms" % (host, response.rtt_avg_ms))
    else:
        print(host, "FAILED")


async def main():
    coroutines = []
    for target in TARGETS:
        coroutines.append(asyncio.ensure_future(myping(target)))
    for coroutine in coroutines:
        await coroutine


if __name__ == "__main__":
    asyncio.run(main())
