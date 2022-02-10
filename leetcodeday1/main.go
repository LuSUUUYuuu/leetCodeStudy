package main

func main() {

	nums := [3]int{1, 2, 3}
	target := 4

	i := twoSum(nums[:], target)

	print(i)
}

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}
