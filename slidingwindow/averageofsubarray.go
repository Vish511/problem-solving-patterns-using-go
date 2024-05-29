package slidingwindow

/* Given an array, find the average of all contiguous subarrays of size ‘K’ in it.

Let’s understand this problem with a real input:

Array: [1, 3, 2, 6, -1, 4, 1, 8, 2], K=5
Output: [2.2, 2.8, 2.4, 3.6, 2.8]

*/

func AvgOfSubArrayOfSizeK(arr []int, k int) []float64 {
	if len(arr) < k {
		return nil
	}
	var windowStart int
	var result []float64
	var sum int
	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		sum += arr[windowEnd]
		if windowEnd-windowStart+1 == k {
			result = append(result, float64(sum)/float64(k))
			sum -= arr[windowStart]
			windowStart++
		}
	}
	return result
}
