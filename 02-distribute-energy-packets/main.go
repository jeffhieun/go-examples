/*
*
# **Distribute Energy Packets**
You are given `n` packets, where `packets[i]` represents the size of the i-th energy packet.
Each packet can be split into any number of smaller sub-packets, but you cannot merge packets together.
You are also given `agents`, representing the number of agents who must each receive a packet of the same size.
Each agent must receive their allocation from one sub-packet only, and some packets may go unused.
Return the maximum number of units each agent can receive equally.
---
### Example 1
**Input:**
```
packets = [7, 10, 4]
agents = 4
```
**Output:**
```
4
```
**Explanation:**
* Divide 7 → 4 + 3
* Divide 10 → 4 + 4 + 2
* Now we have 4 packets of size 4
* Assign 4, 4, 4, 4 to 4 agents
*
*/
package main

import "fmt"

func main() {
	packets := []int64{7, 10, 4}
	agents := int64(4)

	result := solution(packets, agents)

	fmt.Printf("Maxium equal enery per agent: %d\n", result)
}

// -- Solution ---
func solution(packets []int64, agents int64) int64 {
	var left int64 = 1
	var right int64 = 0

	for _, p := range packets {
		if p > right {
			right = p
		}
	}

	var result int64 = 0

	for left <= right {
		mid := left + (right-left)/2

		if canDistribute(packets, agents, mid) {
			result = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}

func canDistribute(packets []int64, agents int64, size int64) bool {
	var count int64 = 0
	for _, p := range packets {
		count += p / size

		if count >= agents {
			return true
		}
	}

	return false
}
