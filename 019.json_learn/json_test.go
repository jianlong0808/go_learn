package json_learn

import (
	"encoding/json"
	"fmt"
	"testing"
)

/*
1. 用于marshal和unmarshal的结构体变量名要为可导出的(首字母大写)
2. json.Marshal可以方便的将结构体对象转换成json字符串
3. 将json字符串转换成对象(json.Unmarshal)的时候可以只取需要的字段, 所取的字段在json中不存在则为默认零值
4. 自带的json包用的是反射技术, 效率很低, 推荐使用第三方高效包进程处理
*/

//测试结构体, 用于marshal
type Movie struct {
	Title  string
	Year   int  `json:"released"`        //tag, Marshal成json串的时候用来显示字段信息
	Color  bool `json:"color,omitempty"` //omitempty: 该字段为默认零值的时候则不显示在json串中
	Actors []string
}

func TestJsonMarshal(t *testing.T) {
	movies := []Movie{
		{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true, Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}

	//ret, err := json.MarshalIndent(movies, "", "    ")
	ret, err := json.Marshal(movies)
	if err != nil {
		Log.Errorf("%#v", err)
		return
	}
	fmt.Println(string(ret))

	/*
		[
		    {
		        "Title": "Casablanca",
		        "released": 1942,
		        "Actors": [
		            "Humphrey Bogart",
		            "Ingrid Bergman"
		        ]
		    },
		    {
		        "Title": "Cool Hand Luke",
		        "released": 1967,
		        "color": true,
		        "Actors": [
		            "Paul Newman"
		        ]
		    },
		    {
		        "Title": "Bullitt",
		        "released": 1968,
		        "color": true,
		        "Actors": [
		            "Steve McQueen",
		            "Jacqueline Bisset"
		        ]
		    }
		]
	*/
}

//测试结构体, 用于unmarshal
type Movies struct {
	Title    string `json:"Title"`
	Released int    `json:"released"`
	Color    bool   `json:"color"`
	Test     string `json:"test"` //对于不存在的字段输出为默认零值
}

func TestJsonUnMarshal(t *testing.T) {
	var movies []Movies
	data := []byte("[{\"Title\":\"Casablanca\",\"released\":1942,\"Actors\":[\"Humphrey Bogart\",\"Ingrid Bergman\"]},{\"Title\":\"Cool Hand Luke\",\"released\":1967,\"color\":true,\"Actors\":[\"Paul Newman\"]},{\"Title\":\"Bullitt\",\"released\":1968,\"color\":true,\"Actors\":[\"Steve McQueen\",\"Jacqueline Bisset\"]}]\n")
	err := json.Unmarshal(data, &movies)
	if err != nil {
		fmt.Printf("%#v\n", err)
		return
	}
	fmt.Printf("%#v\n", movies)
	//[]json_learn.Movies{json_learn.Movies{Title:"Casablanca", Released:1942, Color:false, Test:""}, json_learn.Movies{Title:"Cool Hand Luke", Released:1967, Color:true, Test:""}, json_learn.Movies{Title:"Bullitt", Released:1968, Color:true, Test:""}}
	fmt.Println(movies)
	//[{Casablanca 1942 false } {Cool Hand Luke 1967 true } {Bullitt 1968 true }]
}
