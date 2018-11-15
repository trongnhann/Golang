package devide

func D(N, k uint32, ok bool) ([]uint32, bool) {
	var (
		N1, i uint32
		val   []uint32
	)
	if (k == 0) || (((N%1000 != 0) || (N < 1000)) && (ok == true)) {
		return val, false
	} else {
		if ok {
			N1 = N / 1000
		} else {
			N1 = N
		}

		x := (N1 / k)
		y := N1 % k

		if ok {
			for i = 0; i < k; i++ {
				if i < y {
					val = append(val, (x+1)*1000)
				} else {
					val = append(val, x*1000)
				}
			}
		} else {
			for i = 0; i < k; i++ {
				if i < y {
					val = append(val, x+1)
				} else {
					val = append(val, x)
				}
			}
		}
		return val, true
	}
}
