package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"math/rand"
	"strconv"
	"time"
)

// 声明一个全局的redisDb变量
var redisDb *redis.Client

// 根据redis配置初始化一个客户端
func initClient() (err error) {

	redisDb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // redis地址
		Password: "",               // redis密码，没有则留空
		DB:       1,                // 默认数据库，默认是0
	})

	//通过 *redis.Client.Ping() 来检查是否成功连接到了redis服务器
	_, err = redisDb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func saveHash(n string) {

	for i := 0; i < 100000; i++ {
		title := n + "user:" + strconv.Itoa(rand.Intn(100000000000000))
		_ = redisDb.HSet(title, "username", "pwd").Err()
		_ = redisDb.HSet(title, "password", "abc123").Err()
		_ = redisDb.HSet(title, "create_time", time.Now()).Err()
		//time.Sleep(100 * time.Millisecond)
	}
}
func main() {
	start := time.Now() // 获取当前时间
	err := initClient()
	if err != nil {
		//redis连接错误
		panic(err)
	}
	//go saveHash("1")
	//go saveHash("2")
	saveHash("3")
	cost := time.Since(start) // 计算此时与start的时间差
	fmt.Println(cost)
	// err = redisDb.HSet("user_1", "username", "admin").Err()
	// if err != nil {
	// 	panic(err)
	// }
	// _ = redisDb.HSet("user_1", "password", "abc123").Err()

	// // HGetAll 一次性返回key=user_1的所有hash字段和值
	// data, err := redisDb.HGetAll("user_1").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// // data是一个map类型，这里使用使用循环迭代输出
	// for field, val := range data {
	// 	fmt.Println(field, val)
	// }

	// // 第三个参数代表key的过期时间，0代表不会过期。
	// err = redisDb.Set("name1", "zhangsan", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }
	// var val string
	// // Result函数返回两个值，第一个是key的值，第二个是错误信息
	// val, err = redisDb.Get("name1").Result()
	// // 判断查询是否出错
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("name1的值：", val) //name1的值：zhangsan

	// //仅当列表存在的时候才插入数据,此时列表不存在，无法插入
	// redisDb.LPushX("studentList", "tom")

	// //此时列表不存在，依然可以插入
	// redisDb.LPush("studentList", "jack")

	// //此时列表存在的时候才能插入数据
	// redisDb.LPushX("studentList", "tom")

	// // LPush支持一次插入任意个数据
	// err = redisDb.LPush("studentList", "lily", "lilei", "zhangsan", "lisi").Err()
	// if err != nil {
	// 	panic(err)
	// }
	// // 返回从0开始到-1位置之间的数据，意思就是返回全部数据
	// vals, err := redisDb.LRange("studentList", 0, -1).Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(vals) //注意列表是有序的，输出结果是[lisi zhangsan lilei lily tom jack]

	// /**/

	// // LPush支持一次插入任意个数据
	// err = redisDb.LPush("studentList", "lily", "lilei", "zhangsan", "lisi").Err()
	// if err != nil {
	// 	panic(err)
	// }
	// // 返回从[0,2]位置之间的数据，意思就是返回3个数据
	// vals2, err := redisDb.LRange("studentList", 0, 2).Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(vals2) //注意列表是有序的，输出结果是[lisi zhangsan lilei]

	// // 返回从0开始到-1位置之间的数据，意思就是返回全部数据
	// vals2, err = redisDb.LRange("studentList", 0, -1).Result()
	// if err != nil {
	// 	panic(err)
	// }
	// //返回list集合中的长度
	// studentLen, err := redisDb.LLen("studentList").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("student集合的长度为:", studentLen)

	// //并集&交集&差集
	// redisDb.SAdd("blacklist", "Obama")     // 向 blacklist 中添加元素
	// redisDb.SAdd("blacklist", "Hillary")   // 再次添加
	// redisDb.SAdd("blacklist", "the Elder") // 添加新元素

	// redisDb.SAdd("whitelist", "the Elder") // 向 whitelist 添加元素

	// // 求交集, 即既在黑名单中, 又在白名单中的元素
	// names, err := redisDb.SInter("blacklist", "whitelist").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// // 获取到的元素是 "the Elder"
	// fmt.Println("交集结果是: ", names) // [the Elder]

	// //求交集并将交集保存到 destSet 的集合
	// res, err := redisDb.SInterStore("destSet", "blacklist", "whitelist").Result()
	// fmt.Println(res)
	// //获取交集的值[the Elder]
	// destStr, _ := redisDb.SMembers("destSet").Result()
	// fmt.Println(destStr) //[the Elder]

	// // 求差集
	// diffStr, err := redisDb.SDiff("blacklist", "whitelist").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("差集结果是: ", diffStr) //[Hillary Obama]

}
