package PackageGosayhello

import (
	"math/rand"
)

func GoSayhello(nama string) (string, int) {
	AngkaRandom := rand.Int()
	return "helloo world !! " + nama, AngkaRandom

}
