# The missionaries and cannibals

Three missionaries and three cannibals are on one side of a river, along with a boat that can hold one or two people. Find a way to get everyone to the other side without ever leaving a group of missionaries in one place outnumbered by the cannibals in that place. 

### Implementaion

The problem can be broken down into a graph. The starting point is the root of the node and the possible actions that can be made from that node are the children. A simple BFS search finds the first optimal solution, which is in 11 steps.

                                3M/3C
                                /   \
                             <0,1> <1,0> ....
                              /       \  
                           2M/3C     3M/2C ....
                           ....         ...

### Installation

Install [Go](https://golang.org/doc/install)

```sh
$ git clone git@github.com:zachsitler/go-cannibals-missionaries.git
$ cd go-cannibals-missionaries
$ go install
$ go run cannibals-missionaries.go
```