package main

import (
	"fmt"
	"log"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("sms/aliyun_sms/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	client, err := dysmsapi.NewClientWithAccessKey(
		viper.GetString("sms.cn-hangzhou"),
		viper.GetString("aliyun.accessKeyId"),
		viper.GetString("aliyun.accessSecret"))

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = "15205084501" // 多个手机号以"," 分隔
	request.SignName = "极物联"             // 短信签名名称
	request.TemplateCode = "22"          // 短信模板ID
	request.TemplateParam = "11"         // 短信模板变量对应的实际值，JSON格式, 例: {"code":"1111"}
	request.OutId = "11"

	response, err := client.SendSms(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}
