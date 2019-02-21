package alg

//CountingValleys ...
func CountingValleys(n int32, s string) int32 {
	var l int32
	var v int32
	l = 0
	v = 0
	for _, i := range s {
		if i == 'U' {
			l++
		} else {
			if l == 0 {
				v++
			}
			l--
		}
	}
	return v
}
