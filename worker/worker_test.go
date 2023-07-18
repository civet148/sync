package worker

import (
	"fmt"
	"testing"
	"time"
)

func TestWorkers(t *testing.T) {
	var n = 5
	wks, err := NewWorkers(n)
	if err != nil {
		fmt.Printf("error [%s]\n", err.Error())
		return
	}

	for i := 0; i < n+100; i++ {
		wks.Take()
		go func(idx int) {
			defer wks.Give()
			fmt.Printf("this is worker %d\n", idx)
			time.Sleep(2 * time.Second)
		}(i)
	}
	wks.Wait()
}
