package function

import (
	"fmt"
	"testing"
)

func TestAllocateCoins(t *testing.T) {
	users := []string{"MatthEw", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth"}
	total := 50
	allocation, res := AllocateCoins(users, total)
	fmt.Printf("allocation:%v res:%v\n", allocation, res)
}
