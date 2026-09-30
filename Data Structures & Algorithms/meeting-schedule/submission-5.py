"""
Definition of Interval:
class Interval(object):
    def __init__(self, start, end):
        self.start = start
        self.end = end
"""

class Solution:
    def canAttendMeetings(self, intervals: List[Interval]) -> bool:
        if len(intervals) == 0:
            return True

        def getStartTime(interval):
            return interval.start

        sortedIntervals = sorted(intervals, key=getStartTime)
        prevInterval = sortedIntervals[0]
        for currInterval in range(1, len(sortedIntervals)):
            if prevInterval.end > sortedIntervals[currInterval].start:
                return False
            prevInterval = sortedIntervals[currInterval]
        return True
