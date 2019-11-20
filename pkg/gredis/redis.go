package gredis

import (
    "fmt"
    "github.com/go-redis/redis"
    "iam/pkg/logging"
    "iam/pkg/settings"
    "log"
    "time"
)

var redisCli *redis.Client

func Setup() {
    redisCli = redis.NewClient(&redis.Options{
        Addr: settings.RedisSetting.Host,
        Password: settings.RedisSetting.Password,
        DB: 0,  // 使用默认数据库
        IdleTimeout: settings.RedisSetting.IdleTimeout,
        PoolSize: settings.RedisSetting.MaxIdle,  // 连接池大小， 可以不设置，默认跟cpu数量有关
    })

    // 通过 cient.Ping() 来检查是否成功连接到了 redis 服务器
    _, err := redisCli.Ping().Result()
    if err != nil {
        log.Fatal(err)
    }
    logging.Info("Redis server is connected")
}

// 存入Redis，expiration 为0 的时候表示不超时
func Set(key string, data interface{}, expiration time.Duration) error {
    err := redisCli.Set(key, data, expiration).Err()
    if err != nil {
        return err
    }
    return nil
}

// 根据key从redis获取值
func Get(key string) (string, error) {
    val, err := redisCli.Get(key).Result()
    if err == redis.Nil {
        logging.Warn(fmt.Sprintf("Redis未获取到 key 为：%s数据", key))
        return "", nil
    } else if err != nil {
        return "", err
    } else {
        return val, nil
    }
}

func Del(key string) error {
    err := redisCli.Del(key).Err()
    if err != nil {
        return err
    }
    return nil
}

func Like(key string) (interface{}, error) {
    data, err := redisCli.Do("KEYS", "*" + key + "*").Result()
    if err != nil {
        return nil, err
    }
    return data, nil
}

