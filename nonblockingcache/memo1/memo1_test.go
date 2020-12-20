package memo

import (
	"testing"

	"github.com/TomLan42/snippet/nonblockingcache/util"
)

var httpGetBody = util.HTTPGetBody

func Test(t *testing.T) {
	m := New(httpGetBody)
	util.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := New(httpGetBody)
	util.Concurrent(t, m)
}
