package eutm

type (
	UTMSource   string //  表示流量来源
	UTMMedium   string //  表示流量的媒介
	UTMCampaign string //  表示营销活动名称
	UTMContent  string //  表示内容类型
)

type UTM struct {
	// 流量来源（Source），用来标记流量是从哪个平台或渠道来的(可选)
	UTMSource *UTMSource
	// 流量的媒介（Medium），用来标记流量通过什么方式传播的(可选)
	UTMMedium *UTMMedium
	// 具体的营销活动（Campaign），用来区分不同的营销推广活动
	UTMCampaign *UTMCampaign
	// 邀请码
	UTMReferralCode *string
	// 用于区分相同广告或链接的不同版本(可选)
	UTMContent *UTMContent
	// 用于标记搜索词，通常在付费搜索广告中使用(可选)
	UTMTerm *string
	// 附加信息,传入序列化好的json字符串
	ExtraInfo *string
}
