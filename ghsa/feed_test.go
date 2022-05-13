package ghsa

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseRss(t *testing.T) {
	rss, err := ParseRss("../feed/container.xml")
	assert.NoError(t, err)
	spew.Dump(rss)
}
