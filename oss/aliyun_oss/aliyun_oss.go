package main

import (
	"log"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("oss/aliyun_oss/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	client, err := oss.New(
		viper.GetString("oss.endpoint"),
		viper.GetString("oss.accessKeyId"),
		viper.GetString("oss.accessKeySecret"),
	)
	if err != nil {
		log.Fatal(err)
	}
	bucketName := viper.GetString("oss.bucket")

	bucket, _ := client.Bucket(bucketName)
	err = bucket.PutObjectFromFile("oss/aliyun_oss/aliyun_oss.go", "oss/aliyun_oss/aliyun_oss.go")
	if err != nil {
		log.Fatal(err)
	}
}
