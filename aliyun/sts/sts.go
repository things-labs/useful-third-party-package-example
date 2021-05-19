package main

import (
	"fmt"
	"log"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
)

const (
	accessKeyId  = "accessKeyId"
	accessSecret = "accessSecret"
	roleArn      = "roleArn"
)

func main() {
	//构建一个阿里云客户端, 用于发起请求。
	//构建阿里云客户端时，需要设置AccessKey ID和AccessKey Secret。
	client, err := sts.NewClientWithAccessKey("cn-hangzhou", accessKeyId, accessSecret)
	if err != nil {
		log.Fatal(err)
	}
	//构建请求对象。
	req := sts.CreateAssumeRoleRequest()
	req.Scheme = "https"

	//设置参数。关于参数含义和设置方法，请参见API参考。
	req.RoleArn = roleArn
	req.RoleSessionName = "RoleSessionName"
	req.DurationSeconds = requests.NewInteger(900)

	//发起请求，并得到响应的临时令牌
	rsp, err := client.AssumeRole(req)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("response is %+v\n", rsp)

	// 使用临时令牌的oss上传文件
	// cli, err := oss.New(
	// 	"oss-cn-hangzhou.aliyuncs.com",
	// 	rsp.Credentials.AccessKeyId,
	// 	rsp.Credentials.AccessKeySecret,
	// 	oss.SecurityToken(rsp.Credentials.SecurityToken),
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// bucket, _ := cli.Bucket("zcai-dev")
	// err = bucket.PutObjectFromFile("image/sts.go", "aliyun/sts/sts.go")
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
