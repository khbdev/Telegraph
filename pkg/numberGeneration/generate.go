package numbergeneration

import (
	"fmt"
	"math/rand"
	"time"
)


func AddRandomNumberPrefix(title string) string {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(9000) + 1000 
	return fmt.Sprintf("%s-%d", title, num)
}