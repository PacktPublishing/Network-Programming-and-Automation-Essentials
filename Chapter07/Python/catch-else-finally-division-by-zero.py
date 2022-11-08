

def division(q, d):
    return q/d

try:
    result = division(10, 1)
except ZeroDivisionError:
    print("Error: We should not divide by zero")
else:
    print("Division succeded, result is:", result)
finally:
    print("done")
