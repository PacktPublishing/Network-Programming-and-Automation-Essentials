from aioping import ping
import asyncio

TARGETS = ["yahoo.com", "google.com", "cisco.com", "cern.ch"]


async def myping(host):
    try:
        delay = await ping(host)
        print("%s OK, latency is %.3f ms" % (host, delay * 1000))
    except TimeoutError:
        print(host, "FAILED")


async def main():
    coroutines = []
    for target in TARGETS:
        coroutines.append(asyncio.ensure_future(myping(target)))
    for coroutine in coroutines:
        await coroutine


if __name__ == "__main__":
    asyncio.run(main())
