package str
func Multiply(num1 string, num2 string) string {
	ans := ""
	for i:=0;i<len(num2);i++ {
		ans = add(ans,mul(num1, int(num2[i]-'0'),len(num2)-i-1))
	}
	return ans
}

func mul(num1 string, num2 int, tail int) string {
	if num2 == 0 {
		return "0"
	}
	var ans []byte
	carry := 0
	for i:=len(num1)-1;i>=0;i-- {
		x:=int(num1[i] - '0')
		y := x*num2 + carry
		ans = append(ans, byte(y%10 + '0'))
		carry = y/10
	}
	if carry>0 {
		ans = append(ans, byte(carry+'0'))
	}
	reverse(ans)
	for i:=0;i<tail;i++ {
		ans = append(ans, '0')
	}
	return string(ans)
}

func add (num1, num2 string) string {
	var ans []byte
	carry := 0
	for i:=0;i<len(num1) || i<len(num2) || carry !=0 ;i++{
		x,y := 0,0
		if i<len(num1) {
			x = int(num1[len(num1)-i-1]-'0')
		}
		if i<len(num2) {
			y = int(num2[len(num2)-i-1] - '0')
		}
		ans = append(ans, byte((x+y+carry)%10 + '0'))
		carry = (x+y+carry)/10
	}
	reverse(ans)
	return string(ans)
}

func reverse(bytes []byte) {
	for i,j:=0,len(bytes)-1;i<j;i,j=i+1,j-1 {
		bytes[i],bytes[j] = bytes[j],bytes[i]
	}
}
