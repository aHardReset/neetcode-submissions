from collections import defaultdict

class Solution:

    def findOrder(self, numCourses: int, prerequisites: List[List[int]]) -> List[int]:
        adj = self.getAdjMatrix(prerequisites)
        res = list()
        for course in range(numCourses):
            visited = set()
            if not self.dfs(adj, course, visited, res):
                return []

        taken = set()
        order = []
        for course in res:
            if course not in taken:
                order.append(course)
                taken.add(course)
        return order

    def dfs(self, adj, course, visited, res):
        if course in visited:
            return False
        visited.add(course)
        for prerequisite in adj[course]:
            if not self.dfs(adj, prerequisite, visited, res):
                return False
        res.append(course)
        visited.remove(course)
        return True

    @staticmethod
    def getAdjMatrix(prerequisites):
        adj = defaultdict(list)
        for course, prerequisite in prerequisites:
            adj[course].append(prerequisite)
        return adj
