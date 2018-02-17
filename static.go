package notihub

const (
	DatabaseName  = "notihub"
	RedisDatabase = 2

	RedisNewsKey                 = "index:project:news"
	RedisKeywordKey              = "index:keyword"
	RedisKeywordIndexCompleteKey = "index:keyword:complete"
	RedisNotifyKey               = "notify:project:complete"

	RedisNewKeywordChannel = "index.keyword"
)
