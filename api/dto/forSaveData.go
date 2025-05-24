package dto

import (
	"gorm.io/gorm"
	"reflect"
	"time"
)

type Head struct {
	ProfileId string `gorm:"column:profile_id"`
	ReportId  string `gorm:"column:report_id"`
}

// Genotype 表示单个基因型信息
type Genotype struct {
	gorm.Model
	ForKey      uint   `gorm:"column:for_key"`
	Genotype    string `json:"genotype" gorm:"column:genotype"`
	Summary     string `json:"summary" gorm:"column:summary"`
	TSummary    string `json:"tsummary" gorm:"column:tsummary"`
	Score       string `json:"score" gorm:"column:score"`
	Rsid        string `json:"rsid" gorm:"column:rsid"`
	Gene        string `json:"gene" gorm:"column:gene"`
	OrValue     string `json:"or_value,omitempty" gorm:"column:or_value"`
	Orientation string `json:"orientation,omitempty" gorm:"column:orientation"`
}

type Psychology struct {
	gorm.Model
	Head
	Description string     `json:"description"  gorm:"column:description"`
	Score       float64    `json:"score"  gorm:"column:score"`
	Rank        string     `json:"rank"  gorm:"column:rank"`
	Caseid      string     `json:"caseid"  gorm:"column:caseid"`
	Genotypes   []Genotype `json:"genotypes" gorm:"-"`
}

//`Genotype` 结构体中的 `PsychologyID` 字段将作为外键，关联到 `Psychology` 结构体的主键 `ID`
//todo 记得做迁移，算了使用逻辑外键吧！

type Skin struct {
	gorm.Model
	Head
	Description string     `json:"description" gorm:"column:description"`
	Score       float64    `json:"score" gorm:"column:score"`
	Rank        string     `json:"rank" gorm:"column:rank"`
	CaseID      string     `json:"caseid" gorm:"column:caseid"`
	Genotypes   []Genotype `json:"genotypes" gorm:"-"`
}

type Athletigen struct {
	gorm.Model
	Head
	Description string     `json:"description" gorm:"column:description"`
	Score       float64    `json:"score" gorm:"column:score"`
	Rank        string     `json:"rank" gorm:"column:rank"`
	CaseID      int        `json:"caseid" gorm:"column:caseid"`
	Genotypes   []Genotype `json:"genotypes" gorm:"-"`
}

type HealthyDrug struct {
}

// Result 存入数据库前的中间状态
type HealthResultDto struct {
	Genotypes []Genotype `json:"genotypes" `
	Mag       float64    `json:"mag"`
	Odds      float64    `json:"odds"`
	Summary   []string   `json:"summary"`
	SummaryEn []string   `json:"summary_en" `
	Advise    []string   `json:"advise" `
	AdviseEn  []string   `json:"advise_en"`
}

// Result 表示结果信息
type HealthResult struct {
	gorm.Model
	ForKey    uint       `gorm:"column:for_key"`
	Genotypes []Genotype `json:"genotypes"  gorm:"-"`
	Mag       float64    `json:"mag" gorm:"column:mag"`
	Odds      float64    `json:"odds" gorm:"column:odds"`
	/*//从他返回的数据来看，只是[]里面写字符串，并没有什么复杂的，所以先用string存放
	Summary   []string   `json:"summary" gorm:"column:summary"`
	SummaryEn []string   `json:"summary_en" gorm:"column:summary_en"`
	Advise    []string   `json:"advise" gorm:"column:advise"`
	AdviseEn  []string   `json:"advise_en" gorm:"column:advise_en"`*/
	Summary   string `json:"summary" gorm:"column:summary,type:json"` // 修改为 JSON 字段
	SummaryEn string `json:"summary_en" gorm:"column:summary_en,type:json"`
	Advise    string `json:"advise" gorm:"column:advise,type:json"`
	AdviseEn  string `json:"advise_en" gorm:"column:advise_en,type:json"`
}

// HealthyTraits、HealthyCarrier、HealthyMetabolism共用
type HealthyTraits struct {
	gorm.Model
	Head
	Description      string          `json:"description" gorm:"column:description"`
	DescriptionEn    string          `json:"description_en" gorm:"column:description_en"`
	Mag              float64         `json:"mag" gorm:"column:mag"`
	Odds             float64         `json:"odds" gorm:"column:odds"`
	Sex              string          `json:"sex" gorm:"column:sex"`
	Result           HealthResultDto `json:"result" gorm:"-"`
	AddTime          int             `json:"add_time" gorm:"column:add_time"`
	CustomUpdateTime int             `json:"custom_update_time" gorm:"column:custom_update_time"`
	UpdateTime       int             `json:"update_time" gorm:"column:update_time"`
	CategoryChild    string          `json:"category_child" gorm:"column:category_child"`
	CategoryThird    string          `json:"category_third" gorm:"column:category_third"` //这部分暂时不知道是什么，string先占着！！
	Genotypes        []Genotype      `json:"genotypes" gorm:"-"`
	TSummary         string          `json:"tsummary" gorm:"column:tsummary"`
	CaseID           int             `json:"caseid" gorm:"column:caseid"`
	Score            float64         `gorm:"column:score" json:"score"`
	Rank             string          `gorm:"column:rank" json:"rank"`
}
type HealthyCarrier struct {
	gorm.Model
	Head
	Description      string          `json:"description" gorm:"column:description"`
	DescriptionEn    string          `json:"description_en" gorm:"column:description_en"`
	Mag              float64         `json:"mag" gorm:"column:mag"`
	Odds             float64         `json:"odds" gorm:"column:odds"`
	Sex              string          `json:"sex" gorm:"column:sex"`
	Result           HealthResultDto `json:"result" gorm:"-"`
	AddTime          int             `json:"add_time" gorm:"column:add_time"`
	CustomUpdateTime int             `json:"custom_update_time" gorm:"column:custom_update_time"`
	UpdateTime       int             `json:"update_time" gorm:"column:update_time"`
	CategoryChild    string          `json:"category_child" gorm:"column:category_child"`
	CategoryThird    string          `json:"category_third" gorm:"column:category_third"` //这部分暂时不知道是什么，string先占着！！
	Genotypes        []Genotype      `json:"genotypes" gorm:"-"`
	TSummary         string          `json:"tsummary" gorm:"column:tsummary"`
	CaseID           int             `json:"caseid" gorm:"column:caseid"`
	Score            float64         `gorm:"column:score" json:"score"`
	Rank             string          `gorm:"column:rank" json:"rank"`
}
type HealthyMetabolism struct {
	gorm.Model
	Head
	Description      string          `json:"description" gorm:"column:description"`
	DescriptionEn    string          `json:"description_en" gorm:"column:description_en"`
	Mag              float64         `json:"mag" gorm:"column:mag"`
	Odds             float64         `json:"odds" gorm:"column:odds"`
	Sex              string          `json:"sex" gorm:"column:sex"`
	Result           HealthResultDto `json:"result" gorm:"-"`
	AddTime          int             `json:"add_time" gorm:"column:add_time"`
	CustomUpdateTime int             `json:"custom_update_time" gorm:"column:custom_update_time"`
	UpdateTime       int             `json:"update_time" gorm:"column:update_time"`
	CategoryChild    string          `json:"category_child" gorm:"column:category_child"`
	CategoryThird    string          `json:"category_third" gorm:"column:category_third"` //这部分暂时不知道是什么，string先占着！！
	Genotypes        []Genotype      `json:"genotypes" gorm:"-"`
	TSummary         string          `json:"tsummary" gorm:"column:tsummary"`
	CaseID           int             `json:"caseid" gorm:"column:caseid"`
	Score            float64         `gorm:"column:score" json:"score"`
	Rank             string          `gorm:"column:rank" json:"rank"`
}

type Risk struct {
	gorm.Model
	Head
	Risk        float64    `json:"risk" gorm:"column:risk"`
	Description string     `json:"description" gorm:"column:description"`
	CaseID      int        `json:"caseid" gorm:"column:caseid"`
	Percent     string     `json:"percent" gorm:"column:percent"`
	Genotypes   []Genotype `json:"genotypes" gorm:"-"`
}

//==================================上面是需要循环的，下面的请求一次=====================================

type Block struct {
	ChineseNation float64 `json:"chinese_nation" gorm:"column:chinese_nation"`
	NeAsian       float64 `json:"ne_asian" gorm:"column:ne_asian"`
	SeAsian       float64 `json:"se_asian" gorm:"column:se_asian"`
	SouthAsian    float64 `json:"south_asian" gorm:"column:south_asian"`
	CentralAsian  float64 `json:"central_asian" gorm:"column:central_asian"`
	MiddleEastern float64 `json:"middle_eastern" gorm:"column:middle_eastern"`
	African       float64 `json:"african" gorm:"column:african"`
	European      float64 `json:"european" gorm:"column:european"`
	American      float64 `json:"american" gorm:"column:american"`
	Oceanian      float64 `json:"oceanian" gorm:"column:oceanian"`
}

type Area struct {
	Eskimo         float64 `json:"eskimo" gorm:"column:eskimo"`
	Tungus         float64 `json:"tungus" gorm:"column:tungus"`
	Sindhi         float64 `json:"sindhi" gorm:"column:sindhi"`
	Mbuti          float64 `json:"mbuti" gorm:"column:mbuti"`
	French         float64 `json:"french" gorm:"column:french"`
	Papuan         float64 `json:"papuan" gorm:"column:papuan"`
	Sardinian      float64 `json:"sardinian" gorm:"column:sardinian"`
	Cambodian      float64 `json:"cambodian" gorm:"column:cambodian"`
	Japanese       float64 `json:"japanese" gorm:"column:japanese"`
	HanSouthern    float64 `json:"han_southern" gorm:"column:han_southern"`
	HanNorthern    float64 `json:"han_northern" gorm:"column:han_northern"`
	Mayan          float64 `json:"mayan" gorm:"column:mayan"`
	FinnishRussian float64 `json:"finnish_russian" gorm:"column:finnish_russian"`
	Yoruba         float64 `json:"yoruba" gorm:"column:yoruba"`
	Yakut          float64 `json:"yakut" gorm:"column:yakut"`
	Bantusa        float64 `json:"bantusa" gorm:"column:bantusa"`
	Pima           float64 `json:"pima" gorm:"column:pima"`
	Ny             float64 `json:"ny" gorm:"column:ny"`
	Mongolian      float64 `json:"mongolian" gorm:"column:mongolian"`
	Uygur          float64 `json:"uygur" gorm:"column:uygur"`
	Dai            float64 `json:"dai" gorm:"column:dai"`
	Lahu           float64 `json:"lahu" gorm:"column:lahu"`
	She            float64 `json:"she" gorm:"column:she"`
	Somali         float64 `json:"somali" gorm:"column:somali"`
	Hungarian      float64 `json:"hungarian" gorm:"column:hungarian"`
	Iranian        float64 `json:"iranian" gorm:"column:iranian"`
	Saudi          float64 `json:"saudi" gorm:"column:saudi"`
	Balkan         float64 `json:"balkan" gorm:"column:balkan"`
	Egyptian       float64 `json:"egyptian" gorm:"column:egyptian"`
	Uzbek          float64 `json:"uzbek" gorm:"column:uzbek"`
	Gaoshan        float64 `json:"gaoshan" gorm:"column:gaoshan"`
	Korean         float64 `json:"korean" gorm:"column:korean"`
	English        float64 `json:"english" gorm:"column:english"`
	Spanish        float64 `json:"spanish" gorm:"column:spanish"`
	Kinh           float64 `json:"kinh" gorm:"column:kinh"`
	Bengali        float64 `json:"bengali" gorm:"column:bengali"`
	Thai           float64 `json:"thai" gorm:"column:thai"`
	Ashkenazi      float64 `json:"ashkenazi" gorm:"column:ashkenazi"`
	Kyrgyz         float64 `json:"kyrgyz" gorm:"column:kyrgyz"`
	Tibetan        float64 `json:"tibetan" gorm:"column:tibetan"`
	MiaoYao        float64 `json:"miao_yao" gorm:"column:miao_yao"`
	Mala           float64 `json:"mala" gorm:"column:mala"`
}

type Ancestry struct {
	gorm.Model
	Head
	UniqueID   string `json:"unique_id" gorm:"column:unique_id"`
	UpdateTime int64  `json:"update_time" gorm:"column:update_time"`
	Block      Block  `json:"block" gorm:"embedded;embeddedPrefix:block_"`
	Area       Area   `json:"area" gorm:"embedded;embeddedPrefix:area_"`
}

type Haplogroups struct {
	gorm.Model
	Head
	Y  string `json:"y" gorm:"column:y"`
	Mt string `json:"mt" gorm:"column:mt"`
}

type Demographics struct {
	gorm.Model
	Head
	Surname        string `json:"surname" gorm:"column:surname"`
	NativeProvince string `json:"native_province" gorm:"column:native_province"`
	NativeCity     string `json:"native_city" gorm:"column:native_city"`
	Population     string `json:"population" gorm:"column:population"`
}

// 重复性检测的表
type UniqueProfiles struct {
	gorm.Model
	Address    string    `json:"address" gorm:"column:address"`
	ProfileId  string    `json:"profile_id" gorm:"column:profile_id"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
	Describe   string    `json:"describe" gorm:"column:describe"`
	Status     int       `json:"status" gorm:"column:status"`
}

// 数据访问记录
type DataVisitRecord struct {
	gorm.Model
	Address    string    `json:"address" gorm:"column:address"`
	ProfileID  string    `json:"profile_id" gorm:"column:profile_id"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
}

// 将有gorm.Mod属性属性的结构体注册到map然后写一个方法可以获取他们的类型
var typeRegistry = map[string]reflect.Type{
	//"Genotype":          reflect.TypeOf(Genotype{}),
	"Psychology":        reflect.TypeOf(Psychology{}),
	"Skin":              reflect.TypeOf(Skin{}),
	"Athletigen":        reflect.TypeOf(Athletigen{}),
	"HealthResult":      reflect.TypeOf(HealthResultDto{}),
	"HealthyTraits":     reflect.TypeOf(HealthyTraits{}),
	"HealthyCarrier":    reflect.TypeOf(HealthyCarrier{}),
	"HealthyMetabolism": reflect.TypeOf(HealthyMetabolism{}),
	"Risk":              reflect.TypeOf(Risk{}),
	"Ancestry":          reflect.TypeOf(Ancestry{}),
	"Haplogroups":       reflect.TypeOf(Haplogroups{}),
	"Demographics":      reflect.TypeOf(Demographics{}),
	"UniqueProfiles":    reflect.TypeOf(UniqueProfiles{}),
}

func GetStructType(name string) (reflect.Type, bool) {
	t, ok := typeRegistry[name]
	return t, ok
}
