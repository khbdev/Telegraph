package numbergeneration

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)


func AddRandomNumberPrefix(title string) string {
	rand.Seed(time.Now().UnixNano())


	slug := strings.ToLower(title)


	slug = strings.ReplaceAll(slug, " ", "-")


	num := rand.Intn(9000) + 1000

	return fmt.Sprintf("%s-%d", slug, num)
}