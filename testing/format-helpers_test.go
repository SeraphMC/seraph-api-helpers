package testing

import (
	"github.com/SeraphMC/seraph-api-helpers/src/utils"
	"testing"
	"time"
)

func TestTimeFormatter(t *testing.T) {

	testTimeOne := time.Date(2019, 05, 12, 1, 23, 0, 0, time.UTC)
	gotOne := utils.FormatTimestamp(testTimeOne)
	t.Logf("Result: %s\n", gotOne)

}
