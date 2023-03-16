package main

import (
    "bufio"
    "fmt"
    "os"
)

var (
    visited []bool
    graph   [][]int
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)

    n, m := readInt(scanner), readInt(scanner)

    graph = make([][]int, n+1)
    for i := 1; i <= n; i++ {
        graph[i] = make([]int, 0)
    }

    for i := 0; i < m; i++ {
        u, v := readInt(scanner), readInt(scanner)
        graph[u] = append(graph[u], v)
        graph[v] = append(graph[v], u)
    }

    visited = make([]bool, n+1)

    dfs(1)

    count := 0
    for i := 1; i <= n; i++ {
        if visited[i] {
            count++
        }
    }
    fmt.Println(count)
    for i := 1; i <= n; i++ {
        if visited[i] {
            fmt.Print(i, " ")
        }
    }
}

func dfs(u int) {
    visited[u] = true
    for _, v := range graph[u] {
        if !visited[v] {
            dfs(v)
        }
    }
}

func readInt(scanner *bufio.Scanner) int {
    scanner.Scan()
    res := 0
    for _, c := range scanner.Bytes() {
        res = 10*res + int(c-'0')
    }
    return res
}
