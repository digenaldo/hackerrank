package alg

//SockMerchant ...
func SockMerchant(n int32, ar []int32) int32 {
	var aux = ar
	var c int32
	c = 0
	m1 := make(map[int32]int32)
	for i := 0; i < len(ar); i++ {
		for j := 0; j < len(aux); j++ {
			if aux[j] == ar[i] {
				c++
			}
		}
		m1[ar[i]] = c
		c = 0
	}
	for _, v := range m1 {
		r := v / 2
		if r != 0 {
			c = c + r
		}
	}
	return c
}
