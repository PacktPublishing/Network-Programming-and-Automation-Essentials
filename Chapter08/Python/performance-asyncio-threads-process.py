from pythonping import ping
from multiprocessing import Process
import time
import threading
import sys
import asyncio


def myping(host):
    response = ping(host)


class myPing(threading.Thread):
    def __init__(self, host):
        threading.Thread.__init__(self)
        self.host = host

    def run(self):
        response = ping(self.host)


def main(total):
    print("Multi-threading test ", end="")
    start = time.time()
    threads = []
    for i in range(total):
        t = myPing("127.0.0.1")
        threads.append(t)
        t.start()
    for t in threads:
        t.join()
    print("--- duration %.3f seconds" % (time.time() - start))

    print("Multi-processing test", end="")
    start = time.time()
    processes = []
    for i in range(total):
        p = Process(target=myping, args=("127.0.0.1",))
        processes.append(p)
        p.start()
    for p in processes:
        p.join()
    print("--- duration %.3f seconds" % (time.time() - start))


async def myping_async(host):
    response = ping(host)


async def main_async(total):
    print("Asyncio test", end="")
    start = time.time()
    coroutines = []
    for i in range(total):
        coroutines.append(asyncio.ensure_future(myping_async("127.0.0.1")))
    for coroutine in coroutines:
        await coroutine
    print("--- duration %.3f seconds" % (time.time() - start))


if __name__ == "__main__":
    if len(sys.argv) != 2 or (int(sys.argv[1]) > 100 or int(sys.argv[1]) < 1):
        print("Need one argument number between 1 and 100")
        sys.exit(1)
    total = int(sys.argv[1])
    main(total)
    asyncio.run(main_async(total))
