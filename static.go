package notihub

const (
	DatabaseName  = "notihub"
	RedisDatabase = 2

	RedisNewsKey              = "index:project:news"
	RedisKeywordKey           = "index:keyword"
	RedisKeywordIndexComplete = "index:keyword:complete"
	RedisNotifyKey            = "notify:project:complete"
)
