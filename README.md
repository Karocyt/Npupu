# N-puzzle solver
A* solver for n-puzzle (sliding tiles, mystic square...) using different heuristics, relying only on the standard library.

## Input file format
If you want to provide your own input files, the first line should be the width of your puzzle. The following lines are the multiline space separated puzzle itself. "0" being the empty tile with all tiles ranging from 1 to size*size-1.

    > cat example.txt
    3
    2 6 7
    5 8 4
    1 3 0

By default, the goal state to reach is a snail one :

| 1 | 2 | 3 |
| - | - | - |
| 8 | 0 | 4 |
| 7 | 6 | 5 |

A classic mode is also available (ascending order with empty tile at the very end)

In case you don't provide a file, our generator will take care of you, please refer to the Usage section for options.

## Compilation
    go build

## Usage
    > ./npuzzle -help
    -classic
        uses an ascendant order solution instead of a snail one
    -display
            force print of full solution in any case
    -f string
            filename of your input file
    -h int
            Available heuristics:
                    1: Linear Conflicts
                    2: Manhattan Distance
                    3: Euclidean Distance
                    4: Tiles-Out-Of-Place
            (default 1)
    -mix int
            numbers of scramblings of the generated puzzle if filename omitted (default 200)
    -ref
            adds uniform-cost search for reference
    -shortest
            uses A* algorithm to find an optimal solution
    -size int
            size of the generated puzzle if filename omitted, (3 minimum) (default 3)
    -vs
            compare greedy search and Astar performance

## Structure
This program is structured in 4 independant Go Modules:
- **parser**: reads the input file and check syntax
- **heuristics**: define the canonical heuristic function prototype and contains all our different heuristics
- **sortedhashedtree**: define our data structure with a red-black self-balancing binary search tree for quick access to the lowest score node, with a hashtable to fast access to a specific node
- **solver**: Implements a greedy search algorithm. The "A Star" side of things being handled through the heuristics functions. 

### Algorithms & Heuristics:
Our **Greedy Search** (default) is an *Informed Search* that always explores the open state with the lowest score provided by the heuristic function until to find the goal. Always trying to get closer to the goal at any cost, a greedy search can hence find the **easiest to find path** given its heuristic. With a constant heuristic and a *First In First Out* (stack-like) storage, it should degenerate to a *DFS algorithm*. With a constant heuristic and a *Last In Last Out* storage, it should degenerate to a *BFS algorithm*, and even a *Uniform Cost Uninformed Search* if our case, thanks to our constant cost (we want a solution using the less turns, and all moves takes one turn, no move can take more or less).
A **A&#42;** algorithm uses the same heuristic principle for *the path ahead of us* but adds it the length of the path already traveled (depth of the state from root in our case) to take total distance into account and find an **optimal path**. Doing so, it also keep a lot of open states and ends up being memory inneficient. With a constant heuristic, it should always degenerate into a *Uniform-Cost Uninformed Search*.

In our implementation, we only have a greedy algorithm taking a single heuristic, and the *A star* side of things (adding depth) is done in a dedicated heuristic function. This is why we always have 2 functions for each heuristic: a heuristic function and its *A star* version.

Given the flexibility of our datastructures, we made available **Greedy** mode as a default, to provide a result as fast as possible. Through command-line options, you will also fin **Uniform-Cost** and **A star**

From worst to best, this is the heuristics we implemented, from lazy to almigthy:
1. **Misplaced tiles**: number of tiles that are not at the good position
2. **Euclidean distance**: sum of the flight distances of each tile to its final position
3. **Mannathan/Taxicab distance**: sum of the taxicab distances of each tile to its final position (with only horizontal/vertical travel)
4. **Linear Collisions**: takes the mannathan distance and adds 2 moves when tiles are in linear conflict (when 2 tiles are in the same final row or column wich is their final one, but should be in reverse order on the other axis. Hence, *at least* 2 moves would be needed for the dodging operation)

All these heuristics are *Admissible*, meaning that they assure you to find the shortest path in an Astar implementation faster than a Uniform Cost Search by having 3 properties:
- They give a meaningful information
- The goal state is the one with the lowest score
- They are "optimistic": the score difference with the goal state is never higher than the number of moves required to reach this state

If they all garantee to come to the shortest solution, the later are better *informants* for the search as their score is a closer approximation of the number of moves needed. They effectively cut the clutter and allow us to explore the more usefull states first. This results in a way smaller memory footprint and faster execution times by reducing considerably the number of opened states.
The accuracy of a heuristic could be benchmarked by the difference between the score of the starting state and the number of moves effectively needed to solve the puzzle. This should also be reflected in the length of the path found by a greedy search. (the perfect heuristic should find an optimal path with a greedy search)

### Adding a heuristic
Add your functions to the heuristics package and register those in heuristics.go structure, take a look at some easy heuristics (Mannathan, TOoP), the code is self explanatory.

## Authors
* **Kevin Azoulay** @ 42 Lyon
* **Mickaël Mignot** @ 42 Lyon
