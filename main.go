package main

import (
	"bufio"
	"flag"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/service/s3"
)

var region string
var bucket string
var asset string
var path string

func init() {
	flag.StringVar(&region, "region", "", "aws region")
	flag.StringVar(&bucket, "bucket", "", "bucket name")
	flag.StringVar(&asset, "asset", "", "asset key")
	flag.StringVar(&path, "path", "", "path to write asset - defaults to stdout")
}

func main() {
	flag.Parse()
	if bucket == "" {
		println("bucket is required.")
		os.Exit(1)
	}
	if asset == "" {
		println("asset is required.")
		os.Exit(1)
	}

	if region == "" {
		meta := ec2metadata.New(&ec2metadata.Config{})
		if meta.Available() {
			r, err := meta.Region()
			if err != nil {
				println(err.Error())
				os.Exit(1)
			}
			region = r
		} else {
			println("no region provided, and metadata service unreachable.")
			os.Exit(1)
		}
	}

	config := &aws.Config{
		Region: &region,
	}

	client := s3.New(config)
	object, err := client.GetObject(&s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &asset,
	})
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	var out *bufio.Writer
	if path == "" {
		out = bufio.NewWriter(os.Stdout)
	} else {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			println(err.Error())
			os.Exit(1)
		}
		file, err := os.Create(path + "/" + asset)
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}
		out = bufio.NewWriter(file)
	}
	reader := bufio.NewReader(object.Body)
	_, err = reader.WriteTo(out)
	defer out.Flush()
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}
