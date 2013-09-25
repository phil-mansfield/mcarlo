from __future__ import print_function

import random
import time


t0 = time.time()

for i in range(1000 * 1000): random.random()

t1 = time.time()

print("%g ns per op." % ((t1 - t0)/(1000 * 1000.0) * 1e9))
