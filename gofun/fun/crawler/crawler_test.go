package crawler

import (
	"testing"
)

func TestCrawl(t *testing.T) {
	BreadthFirst(Crawl, []string{"http://www.gopl.io"})
}
