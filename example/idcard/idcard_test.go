package idcard

import (
	"testing"

	"github.com/usthooz/gutil"
)

// TestIdCard
func TestIdCard(t *testing.T) {
	citizenNo := []byte("")
	idcard, err := gutil.GetCitizenNoInfo(citizenNo)
	if err != nil {
		t.Errorf("err-> %v", err)
		return
	}
	t.Logf("birthday-> %d", idcard.Birthday)
	t.Logf("sex-> %v", idcard.Sex)
	t.Logf("addrMask-> %d", idcard.AddrMask)
}
