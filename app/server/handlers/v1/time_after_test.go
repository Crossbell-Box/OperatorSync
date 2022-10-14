package v1

import (
	"testing"
	"time"
)

func TestTimeAfter(t *testing.T) {
	t1 := time.Unix(0, 0)
	t.Log(t1.After(t1)) // false
}
