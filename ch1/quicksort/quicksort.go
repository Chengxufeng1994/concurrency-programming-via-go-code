package quicksort

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[j], arr[low] = arr[low], arr[j]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return i
}

func quicksort(arr []int, low, high int) {
	if low >= high {
		return
	}

	p := partition(arr, low, high)
	quicksort(arr, low, p-1)
	quicksort(arr, p+1, high)
}

func quicksortConcurrent(arr []int, low, high int, done chan struct{}) {
	if low >= high {
		done <- struct{}{}
		return
	}

	p := partition(arr, low, high)
	childDone := make(chan struct{}, 2)
	go quicksortConcurrent(arr, low, p-1, childDone)
	go quicksortConcurrent(arr, p+1, high, childDone)
	<-childDone
	<-childDone
	done <- struct{}{}
}

func quicksortConcurrentParallel(arr []int, low, high int, done chan struct{}, depth int) {
	if low >= high {
		done <- struct{}{}
		return
	}
	depth--
	p := partition(arr, low, high)
	if depth > 0 {
		childDone := make(chan struct{}, 2)
		go quicksortConcurrentParallel(arr, low, p-1, childDone, depth)
		go quicksortConcurrentParallel(arr, p+1, high, childDone, depth)
		<-childDone
		<-childDone
	} else {
		quicksort(arr, low, p-1)
		quicksort(arr, p+1, high)

	}
	done <- struct{}{}
}
