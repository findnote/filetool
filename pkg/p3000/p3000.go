package p3000

// var cfg config.Config
// var orderFilePath string

// var _lastPostOrderArray []string

// type OrdersArray struct {
// 	Data []interface{} `json:"data"`
// }

// func LoadConfig() {
// 	// cfg = config.GetConfig()
// 	// log.Infof("p3000 LoadConfig() cfg=%v", cfg)
// }

// func SetOrderFilesPath(path string) {
// 	orderFilePath = path
// }

// func LoopAndHandlerOrderFiles(path string) {
// 	orderFilePath = path

// 	jsondat := &OrdersArray{}
// 	jsondat.Data = make([]interface{}, 0)

// 	_lastPostOrderArray = []string{}

// 	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
// 		if err != nil {
// 			log.Infof(">>filepath.Walk err: %v", err)
// 			return err
// 		}

// 		if info.IsDir() {
// 			return nil
// 		}

// 		// filter backup dir
// 		if strings.Contains(path, "backup") {
// 			return nil
// 		}

// 		log.Infof("orderFile: %v", path)

// 		jsonstr, errload := utils.LoadFile(path)
// 		if errload != nil {
// 			log.Infof("Failed to load file: %v", path)
// 			return errload
// 		}

// 		orjson := gjson.Get(jsonstr, "data.0").String()
// 		if orjson == "" {
// 			log.Infof("Invalid order json file!!!!")
// 			return nil
// 		}

// 		id := gjson.Get(orjson, "id").String()
// 		_lastPostOrderArray = append(_lastPostOrderArray, id)

// 		var data interface{}
// 		json.Unmarshal([]byte(orjson), &data)

// 		jsondat.Data = append(jsondat.Data, data)

// 		return nil
// 	})

// 	OrdersPostToServer(jsondat.Data)
// }

// OrdersPostToServer 转发调令至P3000
// func OrdersPostToServer(orjson []interface{}) {
// 	LoadConfig()

// 	//当前时间的字符串，2006-01-02 15:04:05据说是golang的诞生时间，固定写法
// 	timeStr := time.Now().Format("2006-01-02 15:04:05")

// 	var req OrdersPostReq
// 	req.ReceiveAt = timeStr
// 	req.OrderType = "Order"

// 	log.Infof(">>orjson: %v, %v", reflect.TypeOf(orjson))
// 	req.Data = orjson

// 	log.Infof(">>req: %v, %s", req, req)

// 	pclient := NewP3000Client()
// 	resp, errcode, _ := pclient.PostOrders(req)

// 	// log.Infof("errcode=%v, err=%v, resp=%v", errcode, err, resp)

// 	// 接口请求失败，直接返回
// 	if errcode == -1 {
// 		return
// 	}

// 	handlerOrderFiles(resp)
// }

// handlerOrderFiles 根据返回的errID，对调令文件进行处理
// func handlerOrderFiles(resp string) error {
// 	log.Infof(">>handlerOrderFile failedIdArray: %v", resp)

// 	if orderFilePath == "" {
// 		return nil
// 	}

// 	// Create a new map.
// 	m := cmap.New()

// 	err := filepath.Walk(orderFilePath, func(path string, info os.FileInfo, err error) error {
// 		if err != nil {
// 			log.Infof(">>filepath.Walk err: %v", err)
// 			return err
// 		}

// 		if info.IsDir() {
// 			return nil
// 		}

// 		// filter backup dir
// 		if strings.Contains(path, "backup") {
// 			return nil
// 		}

// 		jsonstr, errload := utils.LoadFile(path)
// 		if errload != nil {
// 			log.Infof("Failed to load file: %v", path)
// 			return errload
// 		}

// 		datastr := gjson.Get(jsonstr, "data.0").String()
// 		if datastr == "" {
// 			log.Infof("Invalid order json file!!!!")
// 			return nil
// 		}

// 		id := gjson.Get(datastr, "id").String()

// 		_find := FindOrder(id, _lastPostOrderArray)
// 		if _find {
// 			log.Infof(">>found post order, id=%v, path=%v", id, path)
// 			backupPostOrders(path)
// 		}
// 		// log.Infof(">>id: %v, orderFile: %v", id, path)

// 		m.Set(id, path)

// 		return nil
// 	})

// 	errIDArray := gjson.Get(resp, "errID").Array()
// 	for _, id := range errIDArray {
// 		// 从m中获取指定键值.
// 		if tmp, ok := m.Get(id.String()); ok {
// 			path := tmp.(string)
// 			log.Infof(">> id: %v, path: %v", id, path)
// 			backupPostOrders(path)
// 		}
// 	}

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // 处理已提交过的调令（重复提交）
// func backupPostOrders(path string) {
// 	// log.Infof(">>backupOrderFile path: %v", path)

// 	base := filepath.Base(path)
// 	dir := filepath.Dir(path)

// 	// log.Infof(">>base: %v, dir: %v", base, dir)

// 	backupDir := dir + "/backup"
// 	if !Exists(backupDir) {
// 		os.Mkdir(backupDir, os.ModePerm)
// 	}

// 	backupFile := backupDir + "/" + base

// 	err := os.Rename(path, backupFile)
// 	if err != nil {
// 		log.Infof(">>Failed to rename file, err: %v", err)
// 		return
// 	}

// 	// log.Infof(">>backupFile: %v", backupFile)
// }

// // 判断所给路径文件/文件夹是否存在3  直接用os.IsNotExist(err)
// func Exists(path string) bool {
// 	_, err := os.Stat(path) //os.Stat获取文件信息
// 	if err != nil {
// 		return !os.IsNotExist(err)
// 	}

// 	return true
// }

// // 二分法查找相应调令是否已上传
// func FindOrder(target string, str_array []string) bool {
// 	sort.Strings(str_array)

// 	index := sort.SearchStrings(str_array, target)

// 	if index < len(str_array) && str_array[index] == target {

// 		return true

// 	}

// 	return false

// }
