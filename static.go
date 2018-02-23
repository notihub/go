package notihub

const (
	DatabaseName  = "notihub"
	RedisDatabase = 2

	// Indexing keys
	RedisNewsKey                 = "index:project:news"
	RedisKeywordKey              = "index:keyword"
	RedisKeywordIndexCompleteKey = "index:keyword:complete"
	// Notify keys
	RedisNotifyKey               = "notify:project:complete"
	// Project link keys
	RedisProjectLinkKey          = "project:link"
	RedisProjectLinkRequestKey   = "project:link:request"

	RedisNewKeywordChannel = "index.keyword"
)
