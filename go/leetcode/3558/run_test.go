package _test

import (
	"fmt"
	"testing"
)


func TestXxx(t *testing.T) {
	edges :=[][]int{{1,2},{1,3},{2,4}}
	res:=assignEdgeWeights(edges)
	fmt.Println(res)
}

func assignEdgeWeights(edges [][]int) int {
    n:= len(edges) + 1
	g := make([][]int,n + 1)
	for _,e := range edges{
		u,v := e[0],e[1]
		g[u]=append(g[u], v)
		g[v]= append(g[v], u)
	}
	maxDep := dfs(g,1,0)
	return powMod(2,maxDep -1)

}
const MOD = 1_000_000_007
func powMod(a,b int) int{
	res:=1
	a %=MOD

	for b > 0 {
		if b & 1 == 1{
			res = res * a % MOD
		}
		a = a * a %	MOD
		b >>= 1
	}
	return res
}

func dfs(g [][]int,x,f int) int{
	maxDep:= 0

	for _,y:= range g[x]{
		if y == f{
			continue
		}
		maxDep=max(maxDep,dfs(g,y,x) +1)
	}
	return maxDep
}