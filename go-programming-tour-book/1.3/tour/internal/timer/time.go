package timer

import "time"

func GetNowTime() time.Time {
	// 用于返回当前本地时间的 Time 对象，方便后续对 Time 对象做进一步处理
	return time.Now()
}

func GetCalculateTime(currentTimer time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d)  // 在字符串中解析出 duration，例如 "300ms"、"-1.5h"
	if err != nil {
		return time.Time{}, err
	}
	return currentTimer.Add(duration), nil  //  当前 timer 时间加上 duration 后所得到的最终时间
}
