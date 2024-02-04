package main

// import "fmt"

// func main() {
// 	arr := []int{1, 0, 0, 0, 1}
// 	fmt.Println(canPlaceFlowers(arr, 1)) // expected: true
// }

func canPlaceFlowers(flowerbed []int, n int) bool {
	if len(flowerbed) == 1 && flowerbed[0] == 0 && n < 2 {
		flowerbed[0] = n
		return flowerbed[0] >= 0
	}

	for i, v := range flowerbed {
		if i == 0 && v == 0 && flowerbed[i+1] == 0 {
			n -= 1
			flowerbed[i] = 1
		} else if i == len(flowerbed)-1 && v == 0 && flowerbed[len(flowerbed)-2] == 0 {
			n -= 1
			flowerbed[i] = 1
		} else {
			if v == 0 && i > 0 && i < len(flowerbed)-1 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
				n -= 1
				flowerbed[i] = 1
			}
		}
	}

	return n <= 0
}

/*
n = 2
flowerbed[]
for flowerbed
if i's neighbors are 0s --> we replace i with 1 and n-=1
return false if n>0

flowerbed = [1,0,0,0,1], n = 1
for loop
i=1
check 1 -> false
i=0 -> false
i=0 -> true -> flowerbed[i]=1 and n-=1
i....
*/
