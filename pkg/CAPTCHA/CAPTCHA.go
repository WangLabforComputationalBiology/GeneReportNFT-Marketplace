package CAPTCHA

import (
	"GeneReport_platform/configs"
	"GeneReport_platform/pkg/custom_errors"
	"context"
	"errors"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/wenlng/go-captcha-assets/resources/images_v2"
	"github.com/wenlng/go-captcha/v2/base/option"
	"github.com/wenlng/go-captcha/v2/rotate"
)

var rotateCAPTCHA rotate.Captcha

func init() {
	builder := rotate.NewBuilder(rotate.WithRangeAnglePos([]option.RangeVal{
		{Min: 20, Max: 330}, //生成随机角度范围
	}))

	//加载图片资源
	imgs, err := images.GetImages()
	if err != nil {
		log.Fatalln(err)
	}

	//设置图片资源
	builder.SetResources(
		rotate.WithImages(imgs[1:]),
	)
	//构建
	rotateCAPTCHA = builder.Make()
}

// GetRotateCAPTCHA 返回旋转式验证的主图&缩略图
func GetRotateCAPTCHA(address string) (string, string, error) {
	captData, err := rotateCAPTCHA.Generate()
	if err != nil {
		log.Fatalln(err)
	}
	//获取旋转验证元数据
	metaData := captData.GetData()
	if metaData == nil {
		return "", "", custom_errors.New(503, "请稍后再试", errors.New("元数据为空"))
	}
	//存入redis
	_, err = configs.RedisClient.SetEX(context.Background(), "rotate_captcha:"+address, "rotate_angle:"+strconv.Itoa(metaData.Angle), time.Hour*24).Result()
	if err != nil {
		return "", "", custom_errors.New(503, "请稍后再试", err)
	}

	var mBase64, tBase64 string
	mBase64, err = captData.GetMasterImage().ToBase64()
	if err != nil {
		return "", "", custom_errors.New(503, "请稍后再试", err)
	}

	tBase64, err = captData.GetThumbImage().ToBase64()
	if err != nil {
		return "", "", custom_errors.New(503, "请稍后再试", err)
	}

	return mBase64, tBase64, nil
}

// CheckRotateCAPTCHA 执行旋转验证
func CheckRotateCAPTCHA(address string, clientAngle int) (bool, error) {
	target, err := configs.RedisClient.Get(context.Background(), "rotate_captcha:"+address).Result() //获取旋转验证元数据
	if err != nil {
		return false, custom_errors.New(503, "请稍后再试", err)
	}
	correctAngle, _ := strconv.Atoi(strings.Split(target, ":")[1])
	return rotate.CheckAngle(int64(clientAngle), int64(correctAngle), 5), nil
}
