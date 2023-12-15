package model

type Link struct {
	Link      string   `json:"link"`               // 长链接
	Region    []string `json:"region,optional"`    // 地区数字组合
	UserAgent string   `json:"userAgent,optional"` // 用户代理正则编码
	TimeType  int64    `json:"timeType"`

	StartHour int64 `json:"startHour,optional"` // 生效小时
	EndHour   int64 `json:"endHour,optional"`   // 失效小时
	StartMin  int64 `json:"startMin,optional"`  // 生效分钟
	EndMin    int64 `json:"endMin,optional"`    // 失效分钟

	StartDay   int64 `json:"startDay,optional"`
	EndDay     int64 `json:"endDay,optional"`
	StartMonth int64 `json:"startMonth,optional"`
	EndMonth   int64 `json:"endMonth,optional"`

	Hours  []int64  `json:"hours"`
	Days   []int64  `json:"days,optional"`  // 生效日期
	Weeks  []string `json:"week,optional"`  // 生效星期 0-周一 --- 6-周天
	Months []string `json:"month,optional"` // 生效月份
}
