"""
Get current Thread
"""

import threading
import time


def worker():
    print threading.currentThread().getName(), 'start'
    time.sleep(3)
    print threading.currentThread().getName(), 'end'


def service():
    print threading.currentThread().getName(), 'start'
    time.sleep(3)
    print threading.currentThread().getName(), 'end'

t1 = threading.Thread(name='a_worker', target=worker())
t2 = threading.Thread(target=service())

t1.start()
t2.start()

