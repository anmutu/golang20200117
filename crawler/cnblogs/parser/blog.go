/*
  author='du'
  date='2020/1/26 12:58'
*/
package parser

import (
	"fmt"
	"golang20200117/crawler/engine"
	"regexp"
)

//<a href="https://home.cnblogs.com/u/volare/">
//以梦为码
//</a>

//<div class="postDesc">posted @
//<span id="post-date">2020-01-26 17:19</span>&nbsp;
//<a href="https://www.cnblogs.com/volare/">以梦为码</a>&nbsp;
//阅读(<span id="post_view_count">41</span>)&nbsp;
//评论(<span id="post_comment_count">0</span>)&nbsp;
//<a href="https://i.cnblogs.com/EditPosts.aspx?postid=12234466" rel="nofollow">编辑</a>&nbsp;
//<a href="javascript:void(0)" onclick="AddToWz(12234466);return false;">收藏</a></div>

const blogDetailRe = `<span id="post-date">([^<])</span>&nbsp;
<a href="https://www.cnblogs.com/([a-zA-Z0-9])/">([^<])</a>&nbsp;
阅读(<span id="post_view_count">([0-9])</span>)&nbsp;
评论(<span id="post_comment_count">([0-9])</span>)&nbsp;`

func ParseBlogDetail(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(blogDetailRe)

	result := engine.ParseResult{}
	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		fmt.Printf("发布时间:%s,英文名:%s,昵称:%s,阅读量:%d,评论量：%d\n", m[1], m[2], m[3], m[4], m[5])
		result.Items = append(result.Items, string(m[1]))
		result.Requests = append(result.Requests, engine.Request{Url: string(m[1]), ParserFunc: engine.NilParser})
	}
	//fmt.Println(len(matches))
	return result
}
