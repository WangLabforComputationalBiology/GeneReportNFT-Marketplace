package tools

import (
	"GeneReport_platform/api/dto"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strings"
)

var ID string = "szjsbiolab"
var SECRET string = "cgioHbVHp6xp2rZzPthkkp6BRNfsDOr3"

var BASEURL string = "https://api.wegene.com"

var forRisk []int = []int{
	//risk的profile_id
	38, 39, 40, 42, 44, 46, 47, 48, 49, 50, 51, 52, 53, 54, 56, 57, 58,
	59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 79,
	80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100,
	101, 102, 103, 104, 106, 107, 108, 109, 113, 114, 116, 117,
	119, 120, 121, 123, 124, 125, 126, 127, 128, 129, 130, 131, 132, 133, 134,
	135, 136, 137, 138, 139, 141, 216, 242, 243, 1483, 1485, 1521, 1545, 1546, 1547, 1548, 1549,
	1550, 1551, 1552, 1553, 1554, 1555, 1585, 1586, 1587, 1645, 1646, 1647, 1648, 1649, 1650,
}

var forHealthyTraits []int = []int{
	///health/traits
	1, 2, 3, 12, 17, 31, 34, 43, 170, 244, 162, 348, 1440,
}

var forHealthyCarrier []int = []int{
	//healthy_carrier的profile_id
	171, 172, 173, 174, 175, 176, 177, 178, 180, 181, 182, 183, 184, 185, 186, 187, 188, 189, 190, 191, 192, 193,
	194, 195, 196, 197, 198, 199, 200, 201, 202, 203, 205, 206, 207, 208, 209, 210, 211, 212, 214, 223, 232, 1530,
	1531, 1532, 1533, 1534, 1535, 1536, 1537, 1538, 1539, 1540, 1541, 1542, 1543, 1570, 1571, 1572, 1573, 1574,
	1575, 1576, 1577, 1578, 1579, 1580, 1581, 1582, 1583, 1584, 1588,
}
var forHealthyDrug []int = []int{
	///health/drug的report_id
	1461, 1462, 1463, 1464, 1465, 1469, 1470, 1471, 1472, 1474, 1477, 1479, 1480, 1481,
}

var forHealthyMetabolism []int = []int{
	//healthy_metabolism的report_id
	5, 22, 225, 247, 249, 250, 251, 253, 256,
}

var forAthletigen []int = []int{
	///athletigen的report_id
	1486, 1487, 1488, 1489, 1490, 1491, 1492, 1493, 1494, 1495, 1496, 1497, 1498, 1499, 1500, 1501, 1502,
}

var forSkin []int = []int{
	///skin的report_id
	1522, 1523, 1524, 1525, 1526, 1527, 1528, 1529, 1556, 1565, 1566, 1567, 1568, 1569,
}

var forPsychology []int = []int{
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
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
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

	}
}

//health/drug和genotypes还不知道怎么使用
//athletigen、risk、skin、health/carrier、health/metabolism、health/tratis、psychology、这些接口都是只要基因报告id（从user接口获取都是接在url后面）、reportid和token
//demographics/基因报告id、/haplogroups/result/基因报告id、/web_auth/基因报告id \ancestry/profileId这些只要token

// 用到上面的report_id数组的请求头和请求体都差不多，唯一不同的就是响应绑定的对象不同,可以使用泛型来做
func getDataFromWegene[T any](id []int, profileId, url, token string) {
	URL := url + "/" + profileId
	method := "POST"
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
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
		}
		//测试
		fmt.Println(string(body))
		var t T                    //泛型
		tType := reflect.TypeOf(t) // 获取 T 类型的反射类型
		//responseData := reflect.New(tType).Elem() // 使用反射创建 T 类型的实例但是在json.Unmarshal会得到空的数据
		var responseData T
		//获取这个结构体的名称
		name := tType.Name()
		// 解析JSON数据到responseData结构体中
		err = json.Unmarshal(body, &responseData)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
		}
		file, err := os.OpenFile(name+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		//将结构体转换为JSON格式追加到文件中
		jsonData, err := json.Marshal(responseData)
		fmt.Println("jsonData:", string(jsonData)+"\n") //直接打印会出现sonData: [123 34 100 101 115 99 114需要显示用string转换
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
		}
		_, err = file.Write((append(jsonData, '\n')))

		//fixme 这里响应体里面的中文是base64编码，需要转化成中文
	}
}

func SaveAllData(token, profileId string) {

	//athletigen、risk、skin、health/carrier、health/metabolism、health/tratis、psychology
	//health/drug-----X
	/*getDataFromWegene[dto.Risk](forRisk, profileId, BASEURL+"/risk", token)
	getDataFromWegene[dto.HealthyThree](forHealthyTraits, profileId, BASEURL+"/health/traits", token)
	getDataFromWegene[dto.HealthyThree](forHealthyCarrier, profileId, BASEURL+"/health/carrier", token)
	getDataFromWegene[dto.HealthyDrug](forHealthyDrug, profileId, BASEURL+"/health/drug", token)
	getDataFromWegene[dto.HealthyThree](forHealthyMetabolism, profileId, BASEURL+"/health/metabolism", token)
	getDataFromWegene[dto.Athletigen](forAthletigen, profileId, BASEURL+"/athletigen", token)
	getDataFromWegene[dto.Skin](forSkin, profileId, BASEURL+"/skin", token)*/
	getDataFromWegene[dto.Psychology](forPsychology, profileId, BASEURL+"/psychology", token)

}
