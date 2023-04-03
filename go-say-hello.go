package PackageGosayhello

import (
	"math/rand"
)

func GoSayhello() (string, int) {
	AngkaRandom := rand.Int()
	return "helloo world !! ", AngkaRandom

}
