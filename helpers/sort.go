package helpers

func MergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	var mid = len(arr) / 2
	var left = MergeSort(arr[0:mid])
	var right = MergeSort(arr[mid:])
	return mergeHelper(left, right)

}

func mergeHelper(arr1 []int, arr2 []int) []int {
	var i, j int
	var result []int
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}

	for i < len(arr1) {
		result = append(result, arr1[i])
		i++
	}

	for j < len(arr2) {
		result = append(result, arr2[j])
		j++
	}
	return result
}
