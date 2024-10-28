# Assignment 15: Maze Solver with BFS and DFS

## Objective

This assignment will teach you to implement **Breadth-First Search (BFS)** and **Depth-First Search (DFS)** algorithms to find a path through a maze. You will implement both algorithms to solve the same problem and compare their performance and path results.

## Definitions

- **Breadth-First Search (BFS)**: BFS is a graph traversal algorithm that explores all nodes at the present "depth" level before moving on to nodes at the next level. It’s ideal for finding the shortest path in an unweighted graph, as it checks all possible paths layer by layer.

- **Depth-First Search (DFS)**: DFS is a graph traversal algorithm that explores as far as possible along each branch before backtracking. DFS is useful for exploring all possible paths but may not find the shortest path in an unweighted graph.

## Instructions

1. Create a file named `maze_solver.go`.
2. Inside the file, define a `Maze` struct to represent the maze as a 2D grid. Use `0` for open paths and `1` for walls. Define the start and end points within the maze.
3. Implement a `SolveMazeBFS` function using BFS to find a path from the start to the end of the maze.
4. Implement a `SolveMazeDFS` function using DFS to find a path from the start to the end of the maze.
5. Both functions should:
   - Print the path from start to end (or indicate no path exists).
   - Track and print the number of nodes visited during traversal.
6. Demonstrate both algorithms on the same maze and compare the output.

## Example Maze

The maze is represented as a 2D slice of integers. `0` represents open paths, and `1` represents walls.

```go
maze := [][]int{
    {0, 1, 0, 0, 0},
    {0, 1, 0, 1, 0},
    {0, 0, 0, 1, 0},
    {0, 1, 1, 1, 0},
    {0, 0, 0, 0, 0},
}
```

**Start Point**: `(0, 0)`  
**End Point**: `(4, 4)`

## Example Output

```text
Using BFS:
Path found: [(0,0) -> (1,0) -> (2,0) -> (2,1) -> (2,2) -> (3,2) -> (4,2) -> (4,3) -> (4,4)]
Nodes visited: 15

Using DFS:
Path found: [(0,0) -> (1,0) -> (2,0) -> (3,0) -> (4,0) -> (4,1) -> (4,2) -> (4,3) -> (4,4)]
Nodes visited: 20
```
