package main

import (
	"log"
	"os"

	flag "github.com/spf13/pflag"
	"github.com/tupyy/aws-lua/internal/aws"
	"github.com/tupyy/aws-lua/internal/lua"
	"github.com/tupyy/aws-lua/internal/twt"
	glua "github.com/yuin/gopher-lua"
)

var (
	AwsAccessKey string
	AwsSecretkey string
	AwsRegion    string
)

func main() {

	// flags
	luaFile := flag.StringP("filename", "f", "", "lua file")
	flag.StringVar(&AwsAccessKey, "aws-access-key", "", "AWS access key")
	flag.StringVar(&AwsSecretkey, "aws-secret-key", "", "AWS secret key")
	flag.StringVar(&AwsRegion, "aws-region", "", "AWS region")
	flag.Parse()

	if flag.NFlag() != 4 {
		flag.Usage()
		os.Exit(0)
	}

	f, err := os.OpenFile(*luaFile, os.O_RDONLY, 0755)
	if err != nil {
		log.Panic("cannot open lua file")
	}
	f.Close()

	L := glua.NewState()
	defer L.Close()

	awsProvider := aws.New(aws.ClientConfiguration{
		AccessKey: AwsAccessKey,
		SecretKey: AwsSecretkey,
		Region:    AwsRegion,
	})

	twtProvider := twt.New("secret")

	L.PreloadModule("aws", lua.NewAwsModule(awsProvider).Loader)
	L.PreloadModule("twt", lua.NewTwtModule(twtProvider).Loader)

	if err := L.DoFile(*luaFile); err != nil {
		panic(err)
	}

}
