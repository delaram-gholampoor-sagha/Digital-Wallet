package hri

import (
	"fmt"
	"math/rand"
	"time"
)

var random *rand.Rand
var letters = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	source := rand.NewSource(time.Now().UnixNano())
	random = rand.New(source)
}

func Create(prefix string, length int) string {

	b := make([]byte, length)
	for i := range b {
		b[i] = letters[random.Intn(len(letters))]
	}

	return fmt.Sprintf(prefix+"%s-%s", string(b), time.Now().Format("060102"))
}
