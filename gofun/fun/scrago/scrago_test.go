package scrago

import "testing"

func TestNew(t *testing.T) {

	New().SetDefault().AddUrl("https://www.zhihu.com/explore").Start()

}
