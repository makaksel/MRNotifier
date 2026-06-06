package quicksort

// Quicksort sorts an array of integers using the quicksort algorithm.
func Quicksort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }

    pivot := arr[len(arr)/2]
    left := []int{}
    right := []int{}

    for _, value := range arr {
        if value < pivot {
            left = append(left, value)
        } else if value > pivot {
            right = append(right, value)
        }
    }

    left = Quicksort(left)
    right = Quicksort(right)

    return append(append(left, pivot), right...)
}
