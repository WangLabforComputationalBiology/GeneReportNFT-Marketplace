package tools

import (
	"GeneReport_platform/api/dto"
	"GeneReport_platform/configs"
	"GeneReport_platform/internal/models"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
	"sync"
	"time"
)

var BASEURL = "https://api.wegene.com"

var forRisk = []int{
	//risk的profile_id
	38, 39, 40, 42, 44, 46, 47, 48, 49, 50, 51, 52, 53, 54, 56, 57, 58,
	59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 79,
	80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100,
	101, 102, 103, 104, 106, 107, 108, 109, 113, 114, 116, 117,
	119, 120, 121, 123, 124, 125, 126, 127, 128, 129, 130, 131, 132, 133, 134,
	135, 136, 137, 138, 139, 141, 216, 242, 243, 1483, 1485, 1521, 1545, 1546, 1547, 1548, 1549,
	1550, 1551, 1552, 1553, 1554, 1555, 1585, 1586, 1587, 1645, 1646, 1647, 1648, 1649, 1650,
}

var forHealthyTraits = []int{
	///health/traits
	1, 2, 3, 12, 17, 31, 34, 43, 170, 244, 162, 348, 1440,
}

var forHealthyCarrier = []int{
	//healthy_carrier的profile_id
	171, 172, 173, 174, 175, 176, 177, 178, 180, 181, 182, 183, 184, 185, 186, 187, 188, 189, 190, 191, 192, 193,
	194, 195, 196, 197, 198, 199, 200, 201, 202, 203, 205, 206, 207, 208, 209, 210, 211, 212, 214, 223, 232, 1530,
	1531, 1532, 1533, 1534, 1535, 1536, 1537, 1538, 1539, 1540, 1541, 1542, 1543, 1570, 1571, 1572, 1573, 1574,
	1575, 1576, 1577, 1578, 1579, 1580, 1581, 1582, 1583, 1584, 1588,
}
var forHealthyDrug = []int{
	///health/drug的report_id
	1461, 1462, 1463, 1464, 1465, 1469, 1470, 1471, 1472, 1474, 1477, 1479, 1480, 1481,
}

var forHealthyMetabolism = []int{
	//healthy_metabolism的report_id
	5, 22, 225, 247, 249, 250, 251, 253, 256,
}

var forAthletigen = []int{
	///athletigen的report_id
	1486, 1487, 1488, 1489, 1490, 1491, 1492, 1493, 1494, 1495, 1496, 1497, 1498, 1499, 1500, 1501, 1502,
}

var forSkin = []int{
	///skin的report_id
	1522, 1523, 1524, 1525, 1526, 1527, 1528, 1529, 1556, 1565, 1566, 1567, 1568, 1569,
}

var forPsychology = []int{
	///psychology的report_id
	1557, 1558, 1559, 1560, 1561, 1562, 1563, 1564,
}

// 测试微基因的接口是否能拿到数据
func SaveDataTest(token string) {
	url := BASEURL + "/psychology" + "/c5eda99872dba7959fce2406bce3b237"
	method := "POST"
	fmt.Println(url)
	for _, v := range forPsychology {

		payload := strings.NewReader(fmt.Sprintf("report_id=%d", v))

		client := &http.Client{}
		req, err := http.NewRequest(method, url, payload)

		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Add("Authorization", "Bearer "+token)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		//fmt.Println(string(body))//实际上这里打印出来的中文是base64编码
		// 创建一个ResponseData实例来存储解析后的数据
		var responseData dto.Psychology
		// 解析JSON数据到responseData结构体中
		err = json.Unmarshal(body, &responseData)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			return
		}
		//打印解析的结构体
		fmt.Println("解析后的结构体：", responseData)
		//将body写入到当前文件夹（就是项目根目录）下的file.txt中,追加
		//ioutil.WriteFile("file.txt", body, 0644)//这个会覆盖
		// 将内容追加到文件中
		file, err := os.OpenFile("file.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		/*_, err = file.WriteString(string(body) + "\n") // 添加换行符以便区分不同响应
		if err != nil {
			fmt.Println(err)
			return
		}*/
		//将结构体转换为JSON格式追加到文件中
		jsonData, err := json.Marshal(responseData)
		fmt.Println("jsonData:", string(jsonData)) //直接打印会出现sonData: [123 34 100 101 115 99 114需要显示用string转换
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			return
		}
		//_, err = file.Write(jsonData)
		_, err = file.Write(append(jsonData, '\n'))

		//fixme 这里响应体里面的中文是base64编码，需要转化成中文
		_ = res.Body.Close()
	}
}

//health/drug和genotypes还不知道怎么使用
//athletigen、risk、skin、health/carrier、health/metabolism、health/tratis、psychology、这些接口都是只要基因报告id（从user接口获取都是接在url后面）、reportid和token
//demographics/基因报告id、/haplogroups/result/基因报告id、/web_auth/基因报告id \ancestry/profileId这些只要token

// 用到上面的report_id数组的请求头和请求体都差不多，唯一不同的就是响应绑定的对象不同,可以使用泛型来做
func getDataFromWegene[T any](id []int, profileId, url, token, addressT, formatT, sexT string, wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()
	URL := url + "/" + profileId
	method := "POST"

	var t T                    //泛型
	tType := reflect.TypeOf(t) // 获取 T 类型的反射类型
	//responseData := reflect.New(tType).Elem() // 使用反射创建 T 类型的实例但是在json.Unmarshal会得到空的数据
	//获取这个结构体的名称
	name := tType.Name()
	var allJsonStrToHashBulider []byte
	var exitGenotypr = false
	for _, v := range id {
		payload := strings.NewReader(fmt.Sprintf("report_id=%d", v))
		client := &http.Client{}
		req, err := http.NewRequest(method, URL, payload)
		if err != nil {
			fmt.Println(err)
		}
		req.Header.Add("Authorization", "Bearer "+token)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
		}
		//判断状态码是否为200
		if res.StatusCode != 200 {
			fmt.Println("请求失败-", name, "-", profileId, "-reportId:", v)
			continue
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			continue
		}
		//测试
		fmt.Println(string(body))

		var responseData T
		// 解析JSON数据到responseData结构体中
		err = json.Unmarshal(body, &responseData)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
		}

		/*
			file, err := os.OpenFile(name+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			//将结构体转换为JSON格式追加到文件中
			jsonData, err := json.Marshal(responseData)
			fmt.Println("jsonData:", string(jsonData)+"\n") //直接打印会出现sonData: [123 34 100 101 115 99 114需要显示用string转换
			if err != nil {
				fmt.Println("Error marshaling JSON:", err)
			}
			_, err = file.Write((append(jsonData, '\n')))*/

		//获取成功插入后返回的主键id
		//fixme 因为genetype被多个表共用，这里的id最好随机
		var fekId uint
		// 使用反射获取ID属性
		val := reflect.ValueOf(&responseData) //这里提供指针，下面的嵌套结构体才能修改！！！
		//判断反射得到的是否是指针，是就取出指针指向的内存
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}
		//判断内存上的类型是否是结构体
		if val.Kind() == reflect.Struct {
			// 获取Head嵌套结构体的反射值
			headVal := val.FieldByName("Head")
			// 修改ProfileId字段
			profileIdField := headVal.FieldByName("ProfileId")
			if profileIdField.CanSet() {
				profileIdField.SetString(profileId)
			}
			// 修改ReportId字段
			reportIdField := headVal.FieldByName("ReportId")
			if reportIdField.CanSet() {
				reportIdField.SetString(fmt.Sprintf("%d", v))
			}

			//==============================在这里写入库逻辑================================
			result := configs.DB.Create(&responseData) //已经用泛型声明了
			if result.Error != nil {
				fmt.Println(name, "---->创建主表记录错误:", result.Error)
			}
			//如果成功插入就会返回主键ID！

			idField := val.FieldByName("ID")
			if idField.IsValid() && idField.CanInterface() {
				idValue := idField.Interface()
				fekId = idValue.(uint) //类型断言，尝试把 idValue 转换为 uint 类型。
				fmt.Println("ID:", idValue)
			} else {
				fmt.Println("ID属性不存在或不可访问！")
			}

			//=======================上面的操作不知道body具体是什么类型，所以一直用反射，下面确定具体类型的直接转==================
			//利用反射判断是否存在Genotypes这个属性
			genotypesValue := val.FieldByName("Genotypes")
			if genotypesValue.IsValid() { //存在genotype属性
				exitGenotypr = true
				// 遍历 Genotypes 属性的值
				for i := 0; i < genotypesValue.Len(); i++ {
					//将每个Genotypes的值转换为dto.Genotype类型
					genotype := genotypesValue.Index(i).Interface().(dto.Genotype)
					genotype.ProfileId = profileId //添加profile_id
					genotype.ReportId = v          //添加report_id
					genotype.Type = name           //属于哪个结构体的名称
					genotype.ForKey = fekId        //如果上面成功prk必定有值
					result := configs.DB.Create(&genotype)
					if result.Error != nil {
						fmt.Println("逻辑外键关联的genotype插入失败:", result.Error)
					}
				}
				fmt.Println("Genotypes插入成功")
				//序列化这个数组
				jsonStrToHash, err := json.Marshal(genotypesValue.Interface()) //使用 .Interface() 获取实际值
				if err != nil {
					fmt.Println("序列化genotypesValue（一个切片）成json字符串时出错:", err)
				}
				allJsonStrToHashBulider = append(allJsonStrToHashBulider, jsonStrToHash...)

			} else {
				fmt.Println("没有Genotypes属性")
			}
			//判断val里是否有“Result”这个属性
			resultValue := val.FieldByName("Result")
			if resultValue.IsValid() {
				//将resultValue转换为dto.HealthResult类型
				resultDto := resultValue.Interface().(dto.HealthResultDto)
				summaryStr, _ := json.Marshal(resultDto.Summary)
				summaryEnStr, _ := json.Marshal(resultDto.SummaryEn)
				adviseStr, _ := json.Marshal(resultDto.Advise)
				adviseEnStr, _ := json.Marshal(resultDto.AdviseEn)
				healthResult := dto.HealthResult{
					ForKey:    fekId,
					Mag:       resultDto.Mag,
					Odds:      resultDto.Odds,
					Summary:   string(summaryStr),
					SummaryEn: string(summaryEnStr),
					Advise:    string(adviseStr),
					AdviseEn:  string(adviseEnStr),
				}
				//将HealthyResult存入数据库
				result := configs.DB.Create(&healthResult)
				if result.Error != nil {
					fmt.Println("HealthResult插入失败:", result.Error)
				}
				healthyResultId := healthResult.ID
				//对dto.HealthResultDto里面的[]generytype存如数据库
				for _, v := range resultDto.Genotypes {
					v.ForKey = healthyResultId //todo：这里的id最好使用uuid不要和上面一样，有可能会和其他记录的外键重复
					result := configs.DB.Create(&v)
					if result.Error != nil {
						fmt.Println("HealthyResultDto的Genotype插入失败:", result.Error)
					}
				}

			}
		} else {
			fmt.Println("responseData不是一个结构体")
		}
		//单次循环结束，手动关闭
		_ = res.Body.Close()
	}
	if exitGenotypr {
		//var originMetaData dto.Metadatas
		//configs.DB.Where("profile_id = ? and data_hash=''", profileId).First(&originMetaData)
		//计算jsonStrToHash这个字符串的哈希
		hash := sha256.New()
		hash.Write(allJsonStrToHashBulider)
		hashString := hex.EncodeToString(hash.Sum(nil))
		metadata := models.Metadata{
			DataHash:   "0x" + hashString,
			ProfileID:  profileId,
			Category:   name, //ex： skin、risk……
			Owner:      addressT,
			Format:     formatT,
			Sex:        sexT,
			IsSharable: false,
			IsHidden:   true,
			CreatedAt:  time.Now(),
		}
		res := configs.DB.Table("metadatas").Create(&metadata)
		if res.Error == nil {
			fmt.Println(name, "的genotype哈希保存成功")
		} else {
			fmt.Println(name, "的genotype哈希保存失败", res.Error)
		}
	}
	ch <- 1
}
func getDataFromWegeneSimple[T any](profileId, url, token string, wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()
	url += "/" + profileId
	method := "POST"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	req.Header.Add("Authorization", "Bearer "+token)
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))

	var responseData T
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
	}
	val := reflect.ValueOf(&responseData)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() == reflect.Struct {
		// 获取Head嵌套结构体的反射值
		headVal := val.FieldByName("Head")
		// 修改ProfileId字段
		profileIdField := headVal.FieldByName("ProfileId")
		if profileIdField.CanSet() {
			profileIdField.SetString(profileId)
		}
	}

	result := configs.DB.Create(&responseData) //已经用泛型声明了
	if result.Error != nil {
		var t T
		tType := reflect.TypeOf(t)
		name := tType.Name()
		fmt.Println(name, "---->创建主表记录错误:", result.Error)
	}
	ch <- 1
}

// 重复性检测
func checkRepeat(profileId string) bool {
	var count int64
	configs.DB.Model(&dto.UniqueProfiles{}).Where("profile_id = ? and status = 1", profileId).Count(&count)
	return count > 0
}

func SaveAllData(Msg string) {
	var wg sync.WaitGroup

	//将msgs[i].Body的数据按照":"分割 token+id
	parts := strings.Split(Msg, ":")
	token := parts[0]
	profileId := parts[1]
	addressT := parts[2]
	formatT := parts[3]
	sexT := parts[4]
	fmt.Printf("token:%s\nprofileId:%s", parts[0], parts[1])
	if checkRepeat(profileId) {
		fmt.Println("重复性检测：", profileId, "已存在")
		//FIXME 这里的return控制是重复新检测不通过要不要存数，用于开发环境！
		return
	} else {
		fmt.Println("重复性检测通过，开始保存数据：", profileId)
	}

	//athletigen、risk、skin、health/carrier、health/metabolism、health/tratis、psychology
	//health/drug-----Xd
	//getDataFromWegene[dto.HealthyDrug](forHealthyDrug, profileId, BASEURL+"/health/drug", token)
	wg.Add(10)
	completedChan := make(chan int, 10)
	go getDataFromWegene[dto.HealthyTraits](forHealthyTraits, profileId, BASEURL+"/health/traits", token, addressT, formatT, sexT, &wg, completedChan)
	go getDataFromWegene[dto.HealthyCarrier](forHealthyCarrier, profileId, BASEURL+"/health/carrier", token, addressT, formatT, sexT, &wg, completedChan)
	go getDataFromWegene[dto.HealthyMetabolism](forHealthyMetabolism, profileId, BASEURL+"/health/metabolism", token, addressT, formatT, sexT, &wg, completedChan)
	go getDataFromWegene[dto.Risk](forRisk, profileId, BASEURL+"/risk", token, addressT, formatT, sexT, &wg, completedChan)
	go getDataFromWegene[dto.Athletigen](forAthletigen, profileId, BASEURL+"/athletigen", token, addressT, formatT, sexT, &wg, completedChan)
	go getDataFromWegene[dto.Skin](forSkin, profileId, BASEURL+"/skin", token, addressT, formatT, sexT, &wg, completedChan)
	go getDataFromWegene[dto.Psychology](forPsychology, profileId, BASEURL+"/psychology", token, addressT, formatT, sexT, &wg, completedChan)

	//单独的接口
	go getDataFromWegeneSimple[dto.Ancestry](profileId, BASEURL+"/ancestry", token, &wg, completedChan)
	go getDataFromWegeneSimple[dto.Haplogroups](profileId, BASEURL+"/haplogroups", token, &wg, completedChan)
	go getDataFromWegeneSimple[dto.Demographics](profileId, BASEURL+"/demographics", token, &wg, completedChan)

	// 处理子任务完成信号协程
	go func() {
		completed := 0
		ctx := context.Background()
		for range 10 {
			// 接收子任务完成信号
			<-completedChan
			completed++
			// 计算进度（每个子任务 10%）
			progress := float64(completed) * 10
			// 更新 Redis
			err := configs.RedisClient.Set(ctx, "task:"+profileId, progress, 20*time.Minute).Err()
			if err != nil {
				log.Println("Redis 存储进度失败:", err)
			}
			// 发布进度到 Redis 频道
			configs.RedisClient.Publish(ctx, "progress:"+profileId, progress)
		}
		// 所有子任务完成后，设置最终状态
		configs.RedisClient.Set(ctx, "task:"+profileId, "completed", 10*time.Minute)
		configs.RedisClient.Publish(ctx, "progress:"+profileId, "completed")
	}()

	wg.Wait()
	//关闭通道
	close(completedChan)
	//更新redis
	ctxRedis := context.Background()
	configs.RedisClient.Set(ctxRedis, "task:"+profileId, "completed", 24*time.Hour)
	configs.RedisClient.Publish(ctxRedis, "progress:"+profileId, "completed")
	//更新数据库
	configs.DB.Model(&dto.UniqueProfiles{}).Where("profile_id = ?", profileId).Update("status", 1)
	log.Println("数据异步保存完成")
	//修改状态为已完成,根据profileId查找记录

}
