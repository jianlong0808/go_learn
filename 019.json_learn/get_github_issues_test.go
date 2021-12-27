package json_learn

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetIssues(t *testing.T) {
	var err error
	querySlice := []string{"repo:golang/go", "is:open", "json", "decoder"}
	ret, err := GetGitHubIssues(querySlice)
	if err != nil {
		Log.Errorf("%#v", err)
		return
	}
	fmt.Printf("%#v\n", ret)
	data, err := json.MarshalIndent(ret, "", "    ")
	if err != nil {
		Log.Errorf("%#v", err)
		return
	}
	fmt.Println(string(data))

}
