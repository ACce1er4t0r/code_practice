package reverse

func reverse(x int) int {
	res := 0
	sign := 1
	if x < 0 {
		sign = -1
		x *= -1
	}
	for x > 0 {
		tmp := x % 10
		res = res*10 + tmp
		x /= 10
	}
	res *= sign
	if res > 1<<31-1 || res < -1<<31 {
		res = 0
	}
	return res
}
