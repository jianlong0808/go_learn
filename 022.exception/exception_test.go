package exception

import (
	"fmt"
	"go_learn/017.utils/logging"
	"net/http"
	"testing"
	"time"
)

/*
常见的异常处理方式:
1. 返回给调用者, 让调用者处理(此时返回的错误要足够详尽, 如访问url错误, 要体现出被访问的url; 读取文件错误, 要体现出是哪个文件)
2. 错误发生时重试, 但是要合理的控制重试次数和重试时间间隔
3. 记录错误, 退出程序(这种方式最好只在main函数中使用, 在其他包中最好用向上传递的方式)
4. 不处理, 只是记录(适用于不需要处理错误, 只需要记录下来就好的场景)
5. 直接忽略错误, 不接收被调用函数返回的错误信息(适用于已知这种错误无关紧要, 不影响逻辑的场景)

注意:
1. 如果确定错误只有一个, 那么一般返回一个bool值作为错误
2. 对于I/O处理的逻辑一定要考虑异常的情况
3. 尽早出错原则, 先处理可能的所有错误, 然后再处理正常的逻辑
4. 成功时的逻辑代码不应放在else语句块中，而应直接放在函数体中

经验:
1. 更早的出错: 将检测错误的程序, 放在正常的处理逻辑之前
2. 少用 if ret, err := func() ; err != nil {.........} 的形式,
而是用:
ret, err := func()
if err != nil{
    //记录日志
    return .....
}
3. 成功时的逻辑不要放到判断err的else语句块中
*/

var Log, _ = logging.NewLoggerWithRotate()

//1. 将错误返回给调用者
func returnTo() (*http.Response, error) {
	url := "https://www.hao123.com"
	resp, err := http.Get(url)
	if err != nil {
		//将错误直接返回给调用者
		return nil, err
	}
	return resp, nil

}

//2. 发生错误的时候重试, 如果重试失败则返回, 重试的要控制好重试的时间间隔和次数
func waitForResponse() error {
	url := "https://www.jianlong.info/asa"
	deadLine := time.Now().Add(GetTimeout)
	for retries := 0; retries < getRetryTimes && time.Now().Before(deadLine); retries++ {
		reps, err := http.Head(url)
		if err == nil {
			Log.Infof("Get %s success, StatusCode is %d.", url, reps.StatusCode)
			return nil
		}
		Log.Errorf("Get %s failed. retrying…", url)
		time.Sleep(time.Second << retries)
		retries++
	}
	Log.Errorf("Get %s failed.", url)
	return fmt.Errorf("server %s failed to respond after %s", url, GetTimeout)
}

//3. 输出错误, 结束程序, 这种最好只在main函数中执行
func printExit() (*http.Response, error) {
	url := "https://baidu.com"
	reps, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return reps, nil

}

func TestReturn(t *testing.T) {
	//测试将错误返回
	reps, err := returnTo()
	if err != nil {
		return
	}
	Log.Info(reps.StatusCode)

	//测试重试
	if err := waitForResponse(); err != nil {
		Log.Error(err)
	}
	//测试遇到错误直接退出
	reps, err = printExit()
	if err != nil {
		//Fatalf会退出程序
		Log.Fatalf("错误: %v", err)
		//Log.Fatalf("错误: %v", err) 等价于下面两行
		//Log.Errorf("错误: %v", err)
		//os.Exit(1)
	}
	Log.Info(reps.StatusCode)

	//输出错误信息, 但是不退出程序
	reps, err = printExit()
	if err != nil {
		Log.Error(err)
	}
	Log.Info(reps.StatusCode)

	//直接忽略错误, 忽略的变量可以用_代替, _类似于linux中的 /dev/null
	reps, _ = printExit()
	Log.Info(reps.StatusCode)
}

/*
输出:
time="2022-01-02T22:57:21+08:00" level=info msg=200
time="2022-01-02T22:57:22+08:00" level=info msg="Get https://www.jianlong.info/asa success, StatusCode is 404."
time="2022-01-02T22:57:22+08:00" level=info msg=200
time="2022-01-02T22:57:23+08:00" level=info msg=200
time="2022-01-02T22:57:23+08:00" level=info msg=200
*/
