package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		//close可以关闭一个channel
		close(c)
	}()

	for {
		//ok如果为true表示channel未关闭，如果为false表示channel已关闭
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}

	fmt.Println("Main Finished...")
}
