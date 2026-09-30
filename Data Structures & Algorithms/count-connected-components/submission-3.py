class Solution:
    def countComponents(self, n: int, edges: List[List[int]]) -> int:
        adj_list = [[] for _ in range(n)]
        for node_a, node_b in edges:
            adj_list[node_a].append(node_b)
            adj_list[node_b].append(node_a)

        visited = set()
        counter = 0
        for i in range(len(adj_list)):
            if i in visited:
                continue
            self.dfs(adj_list, i, -1, visited, set())
            counter += 1
        return counter


    def dfs(
        self,
        adj_list: List[List[int]],
        current_node: int,
        origin_node: int,
        visited: set,
        path: set,
    ):
        # path.add(current_node)
        visited.add(current_node)
        for connection in adj_list[current_node]:
            if connection in visited:
                continue
            self.dfs(adj_list, connection, current_node, visited, path)

        # path.remove(current_node)


"""
visited set 0, 1, 2
path set
counter 1

[[0,1], [0,2]]
[[1,2], [0], [0]]

0 -- 1
|
2

visited set 0,1,2,3,4,5
path set
counter 2

[[0,1], [1,2], [2,3], [4,5]]
[[1], [0,2], [1,3], [2], [5], [4]]

0 -- 1 -- 2 -- 3        4 -- 5

"""