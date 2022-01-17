package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"time"

	log "github.com/pion/ion-log"
)

func LoadFile(path string) (tojson string, err error) {
	// Open our jsonFile
	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	return string(byteValue), nil
}

func Str2int(str string) int {
	v, _ := strconv.Atoi(str)
	return v
}

func Str2int64(str string) int64 {
	v, _ := strconv.Atoi(str)
	return int64(v)
}

func Int2str(num int) string {
	return strconv.Itoa(num)
}

func GetTimeNow() time.Time {
	return time.Now()
}

func GetTimeNowStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetJsonStrByStruct(req interface{}) string {
	byteArray, err := json.MarshalIndent(req, " ", " ")
	if err != nil {
		log.Errorf(err.Error())
	}

	jsonstr := string(byteArray)

	return jsonstr
}

// 用于查询多条记录的SQL语句中，数组转化为('id1','id2',...,'idn')格式
func SelectArr2Str(arr []string) (str string) {
	str = "("
	for _, v := range arr {
		str = str + "'" + v + "',"
	}
	str = string([]byte(str)[:len(str)-1])
	str = str + ")"

	return str
}

// 用于更新一条记录多个字段的SQL语句中，数组转化为('field1'=Vue1,'field2'=Vue2,,...,'fieldn'=Vuen,)格式
func UpdateMap2Str(mapVue map[string]string) string {
	str := ""
	if len(mapVue) == 0 {
		str = "1=1" //加此条件为了sql条件where不报错
		return str
	}

	for field := range mapVue {
		str = str + "`" + field + "`="
		str = str + "'" + mapVue[field] + "',"
	}

	str = string([]byte(str)[:len(str)-1])
	return str
}

// 用于组按条件模糊查询的SQL语句中
func ListMap2Str(mapVue map[string]string) string {
	str := ""
	if len(mapVue) == 0 {
		str = "1=1" //加此条件为了sql条件where不报错
		return str
	}

	for field := range mapVue {
		if field == "id" || field == "dept_id" || field == "roomId" || field == "station_no" || field == "category" {
			str = str + "`" + field + "` = "
			str = str + "'" + mapVue[field] + "' and "
		} else {
			str = str + "`" + field + "` like "
			str = str + "'%" + mapVue[field] + "%' and "
		}
	}

	str = string([]byte(str)[:len(str)-4])
	return str
}

// 获取需要查询的字段与字段值,并存入map
func GetFindFields(info interface{}) (map[string]string, error) {
	var fieldsMap = make(map[string]string)
	keys := reflect.TypeOf(info)
	vues := reflect.ValueOf(info)

	for k := 0; k < keys.NumField(); k++ {
		if keys.Field(k).Name != "Id" && vues.Field(k).Interface().(string) != "" {
			fieldsMap[keys.Field(k).Name] = vues.Field(k).Interface().(string)
		}
	}

	return fieldsMap, nil
}

func Map2Array(m interface{}) []map[string]interface{} {
	var list []map[string]interface{}
	if reflect.TypeOf(m).Kind() == reflect.Slice {
		s := reflect.ValueOf(m)
		for i := 0; i < s.Len(); i++ {
			ele := s.Index(i)
			list = append(list, ele.Interface().(map[string]interface{}))
		}
	}
	return list
}
