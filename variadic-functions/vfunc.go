package main
import "fmt"

func main(){

	nums := []int{3,5,9,10}

	maxNum := max(1,2,3,4)

	maxNum2 := max(nums...)

	fmt.Println(maxNum, maxNum2)

}

func max( nums ...int)(int) {

	max := nums[0];

	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return max

}