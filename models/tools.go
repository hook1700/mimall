package models

import (
	"crypto/md5"
	"encoding/hex"
	"path"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	. "github.com/hunterhug/go_image"
)

func UnixToDate(timestamp int) string {

	t := time.Unix(int64(timestamp), 0)

	return t.Format("2006-01-02 15:04:05")
}

//2020-05-02 15:04:05
func DateToUnix(str string) int64 {
	template := "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(template, str, time.Local)
	if err != nil {
		beego.Info(err)
		return 0
	}
	return t.Unix()
}

func GetUnix() int64 {
	return time.Now().Unix()
}

//GetUnixNano 获取纳秒
func GetUnixNano() int64 {
	return time.Now().UnixNano()
}

func GetDate() string {
	template := "2006-01-02 15:04:05"
	return time.Now().Format(template)
}

func Md5(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return string(hex.EncodeToString(m.Sum(nil)))
}

func GetDay() string {
	template := "20060102"
	return time.Now().Format(template)
}


// GetSettingFromColumn("SiteTitle")
func GetSettingFromColumn(columnName string) string {

	//redis file
	setting := Setting{}
	DB.First(&setting)
	//反射来获取
	v := reflect.ValueOf(setting)
	val := v.FieldByName(columnName).String()
	return val
}


func ResizeImage(filename string) {
	extName := path.Ext(filename)
	resizeImage := strings.Split(beego.AppConfig.String("resizeImageSize"), ",")

	for i := 0; i < len(resizeImage); i++ {
		w := resizeImage[i]
		width, _ := strconv.Atoi(w)
		savepath := filename + "_" + w + "x" + w + extName
		err := ThumbnailF2F(filename, savepath, width, width)
		if err != nil {
			beego.Error(err)
		}
	}

}


//格式化图片
func FormatImg(picName string) string {
	ossStatus, err := beego.AppConfig.Bool("ossStatus")
	if err != nil {
		//判断目录前面是否有/
		flag := strings.Contains(picName, "/static")
		if flag {
			return picName
		}
		return "/" + picName
	}
	if ossStatus {
		return beego.AppConfig.String("ossDomain") + "/" + picName
	} else {
		flag := strings.Contains(picName, "/static")
		if flag {
			return picName
		}
		return "/" + picName

	}

}