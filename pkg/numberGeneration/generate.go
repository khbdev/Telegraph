package numbergeneration

import (
	"fmt"
	"math/rand"
	"time"
)


func AddRandomNumberPrefix(title string) string {
	rand.Seed(time.Now().UnixNano())

	// 1. lower case
	slug := strings.ToLower(title)

	// 2. bo‘sh joylarni "-" ga almashtirish
	slug = strings.ReplaceAll(slug, " ", "-")

	// 3. random number
	num := rand.Intn(9000) + 1000

	return fmt.Sprintf("/%s-%d", slug, num)
}