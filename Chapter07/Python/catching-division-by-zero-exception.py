

def division(q, d):
    return q/d

try:
    print(division(1, 0))
except ZeroDivisionError:
    print("Error: We should not divide by zero")

