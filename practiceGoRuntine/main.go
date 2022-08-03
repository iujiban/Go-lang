package main

/*
import (
	"fmt"
	"time"
)

func PrintHangul() {
	hanguls := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
	for _, v := range hanguls {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c ", v)

	}
}

func PrintNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d", i)
	}
}

func main() {
	go PrintHangul()
	go PrintNumbers()

	time.Sleep(3 * time.Second)
}

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func SumAtoB(a, b int) {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Printf("%d,,,,,,%d total is %d .\n", a, b, sum)
	wg.Done() //Operation done

}
func main() {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go SumAtoB(1, 10000000)
	}

	wg.Wait()

	fmt.Println("Operation done completed")
}
*/

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

type Account struct {
	Balance int
}

func DepositAndWithdraw(account *Account) {
	mutex.Lock()         //뮤텍스 획득
	defer mutex.Unlock() // defer를 사용한 Unlock()
	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be negative value: %d", account.Balance))
	}

	account.Balance += 1000
	time.Sleep(time.Microsecond)
	account.Balance -= 1000

}

func main() {
	var wg sync.WaitGroup

	account := &Account{0}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				DepositAndWithdraw(account)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
