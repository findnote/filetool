package redisapi

import (
	"mwp3000/api/config"
	"mwp3000/pkg/redis"
	"strconv"
	"time"

	log "github.com/pion/ion-log"
	"github.com/tal-tech/go-zero/core/logx"
)

var (
	isNeedReconnect = true
)
var (
	NULL     = " "
	Error    = "0"
	Stockin  = "1"
	Readying = "2"
	Working  = "3"
	Occupy   = "4"
)

var RClient *redis.Redis

/*
 *	获取连接的redis客户端
 */
func GetRedisClient() *redis.Redis {
	// str := RClient.HGet("CYGDW:Hash:Device:1:101", "yx")
	// log.Infof("SSSSSSGetRedisClient>> %v", str)
	return RClient
}

/*
 *	监听p3000实时库中，接地线实时值状态，每5秒同步一次
 */
func InitP3000RedisListen(c config.Config) {
	RClient = redisConnect(c)

	go func() {
		log.Infof(">>startRedisListen")

		for {
			if !isNeedReconnect {
				// ping 实时库
				if !RClient.Ping() {
					log.Infof("------Ping Redis Err， Reconnecting!!-------")
					isNeedReconnect = true
				} // else {
				// 查询数据库，获取需要监听的jdx设备
				// listenP90Redis(client)
				// logx.Infof("P90 redis listening!!")
				// }
			} else {
				log.Infof("------Reconnecting!!-------")
				RClient = redisConnect(c)
			}
			time.Sleep(time.Second * 5)
		}

	}()
}

/*
 *	redis连接
 */
func redisConnect(c config.Config) *redis.Redis {
	host := c.P3000Redis.Host
	port := c.P3000Redis.Port

	url := host + ":" + strconv.Itoa(port)
	var adds = []string{url}
	logx.Infof("adds: %s, pwd: %s, redis: %s")
	var conf = redis.Config{Addrs: adds, Pwd: c.P3000Redis.Pwd, DB: c.P3000Redis.DB}
	var redisClient = redis.NewRedis(conf)
	if redisClient == nil {
		isNeedReconnect = true
	} else {
		log.Infof("------redis Reconnect success!!-------")
		isNeedReconnect = false
	}
	return redisClient
}

// func listenP90Redis(client *redis.Redis) {
// 	mapList := make(map[string]string)
// 	fieldsStr := utils.ListMap2Str(mapList)
// 	list, _, err := models.NewTJDXDeviceInfoModel().List(fieldsStr, 1, 10000)
// 	if err != nil {
// 		log.Infof("listenP90Redis()>>NewTJDXDeviceInfoModel().List() Failed!!")
// 		return
// 	}

// 	for _, index := range *list {
// 		var devStatus models.JDXStatusDetails

// 		strTmp := "CYGDW:Hash:Device:" + utils.Int2str(int(index.StationNum)) + ":" + index.DeviceNum
// 		mapVue := client.HGetAll(strTmp)

// 		// 实时态不需更新
// 		if len(mapVue) == 0 || index.DevStatus == mapVue["yx"] || (mapVue["yx"] == Occupy && index.DevStatus == Readying) || mapVue["yx"] == Error {
// 			continue

// 		} else if mapVue["yx"] == Working {
// 			jsonStr := mapVue["hanged_dev"]
// 			index.DevStatus = mapVue["yx"]
// 			index.JDDNum = gjson.Get(jsonStr, "deviceNum").String()
// 			devStatus.JDDNum = index.JDDNum
// 		} else if mapVue["yx"] == Stockin || mapVue["yx"] == Readying {
// 			index.DevStatus = mapVue["yx"]
// 			index.JDDNum = NULL
// 			devStatus.JDDNum = NULL
// 		} else {
// 			index.DevStatus = Readying
// 			index.JDDNum = NULL
// 			devStatus.JDDNum = NULL
// 		}

// 		// 新增变位明细
// 		devStatus.CreatedAt = utils.GetTimeNow()
// 		devStatus.StationNum = index.StationNum
// 		devStatus.DeviceNum = index.DeviceNum
// 		devStatus.DeviceName = index.DeviceName
// 		devStatus.Voltage = index.Voltage
// 		devStatus.Status = index.DevStatus
// 		devStatus.TimeStamp = mapVue["yx_tc"]
// 		err = models.NewTJDXStatusDetailsModel().Insert(&devStatus)
// 		if err != nil {
// 			log.Infof("listenP90Redis()>>NewTJDXStatusDetailsModel().Insert() Failed!!")
// 			return
// 		}

// 		// 更新状态
// 		index.UpdatedAt = utils.GetTimeNow()
// 		err = models.NewTJDXDeviceInfoModel().Update(&index)
// 		if err != nil {
// 			log.Infof("listenP90Redis()>>NewTJDXDeviceInfoModel().Update() Failed!!")
// 			return
// 		}

// 	}
// }
