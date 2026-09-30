class Solution:
    def validTree(self, n: int, edges: List[List[int]]) -> bool:
        {
            0: {1,2,3},
            1: {0, 4},
            2: {0},
            3: {0},
            4: {1},
        }

        {
            0: {1},
            1: {0,2,3,4},
            2: {1,3},
            3: {2,1},
            4: {1},
        }
        adj_map = [[] for _ in range(n)]
        for node_a, node_b in edges:
            adj_map[node_a].append(node_b)
            adj_map[node_b].append(node_a)
        visited = set()
        if self.dfs(adj_map, 0, -1, visited, set()) == False:
            return False
        print(len(visited))
        return len(visited) == n

    def dfs(
        self,
        adj_map: List[List[int]],
        current: int,
        parent: int,
        visited: set,
        path: set
    ) -> bool:
        print(current, parent, visited, path)
        path.add(current)
        visited.add(current)
        print(adj_map[current])
        for child in adj_map[current]:
            if child == parent:
                continue
            elif child in path:
                print(11)
                return False
            elif child not in visited:
                if self.dfs(adj_map, child, current, visited, path) == False:
                    print(22)
                    return False
        path.remove(current)
        return True



"""
 0 --- 1 --- 4
 |\
 | \
 2  3
"""

"""
0 --- 1 --- 2 --- 3
     /|           |
    4 -------------
"""
