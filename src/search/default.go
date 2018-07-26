package search

// defaultMatcher implements the default matcher.
type defaultMatcher struct{}

// init registers the default matcher with the program.
// 注册 default 名字的匹配器， defaltMatcher 类型
func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

// Search implements the behavior for the default matcher.
// 实现匹配器的方法
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
