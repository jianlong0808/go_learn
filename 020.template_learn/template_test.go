package template_learn

/*
提供条件判断，数组或map遍历；参数赋值，函数或方法调用；自定义函数扩展，模板嵌套及重用等功能。
基于该工具，可以轻松实现复杂场景的文本渲染。如Helm Template基于此实现了功能强大的Kubernetes配置文件渲染工作。
*/

import (
	"go_learn/017.utils/logging"
	jsonLearn "go_learn/019.json_learn"
	"html/template"
	"os"
	"testing"
	"time"
)

var Log, _ = logging.NewLoggerWithRotate()

//CreatedAt字段值作为daysAgo函数的参数(更多用法参见: https://pkg.go.dev/text/template)
var tempTest = `
{{ .TotalCount }} issues:
{{range .Items}}----------------------------------------
Url: {{.Url}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{ end }}
`

//用于映射template模板中的daysAge
func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func TestTemplate(t *testing.T) {
	ret, err := jsonLearn.GetGitHubIssues([]string{"repo:golang/go", "is:open", "json", "decoder"})
	if err != nil {
		if err != nil {
			Log.Errorf("%#v\n", err)
		}
	}
	var report = template.Must(template.New("测试模板").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(tempTest))
	report.Execute(os.Stdout, ret)

}
