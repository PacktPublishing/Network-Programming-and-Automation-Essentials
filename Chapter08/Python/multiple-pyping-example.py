from pythonping import ping
from multiprocessing import Process

TARGETS = ["yahoo.com", "google.com", "cisco.com", "cern.ch"]

def myping(host):
    response = ping(host)
    if response.success:
        print("%s OK, latency is %.2fms" % (host, response.rtt_avg_ms))
    else:
        print(host, "FAILED")

def main():
    for host in TARGETS:
        Process(target=myping, args=(host,)).start()

if __name__ == "__main__":
        main()
