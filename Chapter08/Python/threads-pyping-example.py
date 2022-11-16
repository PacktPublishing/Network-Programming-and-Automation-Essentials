from pythonping import ping
import threading

TARGETS = ["yahoo.com", "google.com", "cisco.com", "cern.ch"]

class myPing(threading.Thread):
    def __init__(self, host):
        threading.Thread.__init__(self)
        self.host = host

    def run(self):
        response = ping(self.host)
        if response.success:
            print("%s OK, latency is %.2fms" % (self.host, response.rtt_avg_ms))
        else:
            print(self.host, "FAILED")

def main():
    for host in TARGETS:
        myPing(host).start()

if __name__ == "__main__":
    main()
