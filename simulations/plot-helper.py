import os
import sys

type = sys.argv[1]

if type == "E":
	args = "-p0 1 2 -xlabel \"$ T$\" -ylabel \"$ E$\""
elif type == "Mag":
	args = "-p0 1 3 -xlabel \"$ T$\" -ylabel \"$ m$\""
elif type == "C":
	args = "-p0 1 4 -xlabel \"$ T$\" -ylabel \"$ C$\""
elif type == "X":
	args = "-p0 1 5 -xlabel \"$ T$\" -ylabel \"$ X$\""
else:
	assert(0)

def getTableWidth(fileName):
	parts = fileName.split(".")
	if len(parts) != 2 or parts[1] != "table":
		return None
	nameParts = parts[0].split("temp")
	if len(nameParts) != 2 or nameParts[0] != "":
		return None
	try:
		print fileName
		return int(nameParts[1])
	except:
		return None

tableNames, titleArgs = [], []

for file in os.listdir("."):
	n = getTableWidth(file)
	if n is not None:
		tableNames.append(file)
		titleArgs.append("-title \"Lattice Width: %d\"" % n)

for (tableName, titleArg) in zip(tableNames, titleArgs):
	print "~/code/python/cmd-plot/plot.py %s %s %s" % (tableName, args, titleArg)
	os.system("~/code/python/cmd-plot/plot.py %s %s %s &" % (tableName, args, titleArg))
