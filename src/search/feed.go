package search

import (
	"encoding/json"
	"os"
)

const dataFile = "data/data.json"

/*什么是feed流
最近常听朋友说中了抖音的毒，一有时间就刷抖音，根本停不下来。

刷朋友圈、逛微博，以及现在很火的短视频，我们每天有大量时间消耗在“feed流”中，并且刷的不亦乐乎。

以下是维基百科中关于“web feed”的定义：

> a web feed (or news feed) is a data format used for providing users with frequently updated content. 
> Content distributors syndicatea web feed, thereby allowing users to subscribe a channel to it

feed是一种给用户持续提供、频繁更新内容的数据形式。文章提供者发布feed，因此使得用户可订阅该频道。

feed由多个内容提供源组成的资源聚合器，由用户主动订阅消息源并且向用户提供内容。

feed 概念类比：
以微信订阅号为例，“订阅号”为聚合器（聚合了多个公众号），“公众号”是聚合器里面的数据格式，有公众号更新了，推送给你，就是feed流。
朋友圈里，feed 是以“朋友圈”为聚合器(聚合了多位朋友)，feed的数据格式是“朋友”，朋友发朋友圈，小红点提醒了，这也是feed流。

以上举例的feed是按照时间形式来推送，即最新更新了会推送给你。也有按照主题形式来推送的，比如微博的热点话题。

下面我们用json文件作为一个“聚合器”，存放了若干需要访问的网址

*/


// 定义源的结构体类型；源是我们自己准备好的一个json文件
// 用来存放要搜索的网址的信息，包含Name，URI，Type
// 代码需要解读我们的json文件，来做转换string类型，因此需要准备解读格式的“容器”struct变量。
type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}


//检索feeds，并反序列化feed数据文档；
//无参数输入，会返回一个类型是Feed指针的切片,即有多个*Feed组成。
//该切片装了的内容是Json已解码后的内容。
//把我们准备好的json文件进行解读
func RetrieveFeeds() ([]*Feed, error) {
	// 打开准备好的文件
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	// 当函数结束后，文件关闭
	defer file.Close()

	//把文件（json）解码到我们定义好的Feed切片指针类型
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)
	//复合语句，等同于 
	// a := json.NewDecoder(file);  
	// err = a.Decode(&fedds)

	// 不需要检查错误，返回err,后面其它函数会处理
	return feeds, err
}
