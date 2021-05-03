package easy

import "math"

/**

7. Reverse Integer
Given a signed 32-bit integer x, return x with its digits reversed.
If reversing x causes the value to go outside the signed 32-bit integer range [-2^31, 2^31 - 1], then return 0.
Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-integer
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func reverse(x int) (r int) {

	for x != 0 {
		r = r * 10 + x % 10
		x /= 10
	}

	if r > math.MaxInt32 || r < math.MinInt32 {
		return 0
	}
	return r
}