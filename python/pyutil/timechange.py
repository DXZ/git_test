# 5993c7d255972fb4


# 5993c999feb38414

# 2018-11-12 15:15:51.725246339


import datetime
import time
a = {}
format = "%Y-%m-%d %H:%M:%S.%f"


def string2timestamp(strValue):
    d = datetime.datetime.strptime(strValue, "%Y-%m-%d %H:%M:%S.%f")
    t = d.timetuple()
    timeStamp = int(time.mktime(t))
    timeStamp = float(str(timeStamp) + str("%06d" % d.microsecond))/1000000
    return timeStamp


fr = open('result.txt','w')
with open('temp.log','r') as f:
    lines = f.readlines()
    for line in lines:
        item = line.split(" ")
        taskid = item[6].strip()
        timestr = item[1] + " " + item[2]
        print timestr,taskid
        if line[0] == '-':
            a[taskid] = timestr
            continue
        else:
            start = a[taskid]
            start = string2timestamp(start[:26])
            endtime = string2timestamp(timestr[:26])
            gap = float(endtime - start)
            if gap <= 0.005:
                gap = 0
            print(gap)
            fr.write(taskid+"\t"+timestr+"\t"+str(gap)+"\n")



