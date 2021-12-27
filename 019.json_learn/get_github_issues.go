package json_learn

import (
	"encoding/json"
	"errors"
	"go_learn/017.utils/logging"
	"net/http"
	"strings"
	"time"
)

var Log, _ = logging.NewLoggerWithRotate()

const IssuesURL = "https://api.github.com/search/issues"

//Result 结构体定义, 注: 实际上gitlab的issues接口返回的json字符串比我定义的结构体返回的东西要多
type Result struct {
	TotalCount        int    `json:"total_count"`
	IncompleteResults bool   `json:"incomplete_results"`
	Items             []Item `json:"items"`
}

type Item struct {
	Url           string    `json:"url"`
	RepositoryUrl string    `json:"repository_url"`
	User          User      `json:"user"`
	Labels        []Label   `json:"labels"`
	Title         string    `json:"title"`
	CreatedAt     time.Time `json:"created_at"`
}

type User struct {
	Login string `json:"login"`
	Id    int    `json:"id"`
}

type Label struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

//q []string 传入查询条件
func GetGitHubIssues(q []string) (*Result, error) {

	var ret Result
	query := strings.Join(q, "&")
	//请求
	resp, err := http.Get(IssuesURL + "?q=" + query)
	defer resp.Body.Close()
	Log.Info(IssuesURL + "?q=" + query)
	if err != nil {
		Log.Errorf("%#v", err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("search query failed: " + resp.Status)
	}

	//从一个输入流解码JSON数据
	err = json.NewDecoder(resp.Body).Decode(&ret)
	if err != nil {
		Log.Errorf("%#v", err)
		return nil, err
	}

	return &ret, nil
}
