package redisson

type RedisKey = string

const (
	Authentication RedisKey = "authentication"
)

func AppendRedisKey(preinstall RedisKey, suffix string) string {
	return preinstall + ":" + suffix
}
