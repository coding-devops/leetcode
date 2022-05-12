package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd_go_env := exec.Command("git", "grep", "-nI")
	stdout, _ := cmd_go_env.StdoutPipe()
	//stdin, _ := cmd_go_env.StdinPipe()
	if err := cmd_go_env.Start(); err != nil {
		return
	}
	a1 := make([]byte, 1024)
	n, _ := stdout.Read(a1)
	fmt.Println(n)
	fmt.Printf("asd", a1[:n])
	cmd_go_env.Wait()
}

func isanagram(nums []int, target int) []int {
	a := make(map[int]int)
	for index := range nums {
		tmp := target - nums[index]
		if re, ok := a[tmp]; ok {
			var results []int
			results[0] = re
			results[1] = a[nums[index]]
		}
		a[nums[index]] = index
	}
	return nil
}
