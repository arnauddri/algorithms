package merge

func sort(arr []int) {
	var s = make([]int, len(arr)/2+1)
	if len(arr) < 2 {
		return
	}

	mid := len(arr) / 2

	sort(arr[:mid])
	sort(arr[mid:])

	if arr[mid-1] <= arr[mid] {
		return
	}

	copy(s, arr[:mid])

	l, r := 0, mid

	for i := 0; ; i++ {
		if s[l] <= arr[r] {
			arr[i] = s[l]
			l++

			if l == mid {
				break
			}
		} else {
			arr[i] = arr[r]
			r++
			if r == len(arr) {
				copy(arr[i+1:], s[l:mid])
				break
			}
		}
	}
	return
}
