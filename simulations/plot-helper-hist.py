import os
import sys

fileName = sys.argv[1]

if sys.argv[2] == "E":
	os.system(("""~/code/python/cmd-plot/plot.py %s""" +
						 """ -p0 1 2 b "10 Sweeps" -p1 1 4 g "100 Sweeps" """ +
						 """ -p2 1 6 c "1,000 Sweeps" -p3 1 8 m "10,000 Sweeps" """ +
						 """-p4 1 10 r "100,000 Sweeps" """ +
						 """-p5 1 12 k "1000,000 Sweeps" """ +
						 """ -nw""") % fileName)
elif sys.argv[2] == "C":
	os.system(("""~/code/python/cmd-plot/plot.py %s""" +
						 """ -p0 1 3 b "10 Sweeps" -p1 1 5 g "100 Sweeps" """ +
						 """ -p2 1 7 c "1,000 Sweeps" -p3 1 9 m "10,000 Sweeps" """ +
						 """-p4 1 11 r "100,000 Sweeps" """ +
						 """-p5 1 13 k "1000,000 Sweeps" """ +
						 """-ne""") % fileName)
else:
	assert(0)
