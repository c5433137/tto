package tto

import (
	"regexp"
	"strings"
)

var (
	predefineRegexp = regexp.MustCompilePOSIX("#([^\\!-]+)\\!(.+)")
)

type CodeTemplate struct {
	path      string
	template  string
	predefine map[string]string
}

func NewTemplate(s string, path string) *CodeTemplate {
	return (&CodeTemplate{path: path}).configure(s)
}

func (g *CodeTemplate) configure(s string) *CodeTemplate {
	g.predefine = make(map[string]string)
	for _, match := range predefineRegexp.FindAllStringSubmatch(s, -1) {
		g.predefine[match[1]] = match[2]
	}
	g.template = predefineRegexp.ReplaceAllString(s, "")
	return g
}

// 文件路径
func (g *CodeTemplate) FilePath() string {
	return g.path
}

// 获取模板内容
func (g *CodeTemplate) String() string {
	return g.template
}

// 获取预定义的参数
func (g *CodeTemplate) Predefine(key string) (string, bool) {
	n, ok := g.predefine[key]
	return n, ok
}

func (g *CodeTemplate) Replace(s, old string, n int) *CodeTemplate {
	g.template = strings.Replace(g.template, s, old, n)
	return g
}
