package cache

import "github.com/charlie-bit/utils/third_party/go-redis"

const (
	maxSeq = "MAX_SEQ:" // conversation max seq
)

func SetConvIDSeq(
	rdb redis.UniversalClient,
	conversationID string,
	seq int64,
) (int64, error) {
	return rdb.IncrBy(getMaxSeqKey(conversationID), seq).Result()
}

func GetConvIDSeq(rdb redis.UniversalClient, conversationID string) (int64, error) {
	return rdb.Get(getMaxSeqKey(conversationID)).Int64()
}

func getMaxSeqKey(conversationID string) string {
	return maxSeq + conversationID
}
