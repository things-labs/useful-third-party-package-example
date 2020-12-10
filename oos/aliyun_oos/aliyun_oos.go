package main

import (
	"log"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("oos/aliyun_oos/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	client, err := oss.New(
		viper.GetString("oos.endpoint"),
		viper.GetString("oos.accessKeyId"),
		viper.GetString("oos.accessKeySecret"),
	)
	if err != nil {
		log.Fatal(err)
	}
	bucketName := viper.GetString("oos.bucket")

	bucket, _ := client.Bucket(bucketName)
	err = bucket.PutObjectFromFile("oos/aliyun_oos/aliyun_oos.go", "oos/aliyun_oos/aliyun_oos.go")
	if err != nil {
		log.Fatal(err)
	}
}
