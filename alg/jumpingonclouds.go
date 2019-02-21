package alg

//JumpingOnClouds ...
func JumpingOnClouds(c []int32) int32 {
	var j int32
	var n = len(c)
	var i = 0
	for i < n-1 {
		if i+2 < n && c[i+2] == 0 {
			i = i + 2
			j++
		} else {
			i++
			j++
		}
	}
	return j
}
