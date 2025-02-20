/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

// https://developers.e.qq.com/docs/api/business_assets/video/videos_add

package main

import (
	"encoding/json"
	"fmt"
	"github.com/antihax/optional"
	"gopkg.in/resty.v1"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/tencentad/marketing-api-go-sdk/pkg/ads"
	"github.com/tencentad/marketing-api-go-sdk/pkg/api"
	"github.com/tencentad/marketing-api-go-sdk/pkg/config"
	"github.com/tencentad/marketing-api-go-sdk/pkg/errors"
	"github.com/tencentad/marketing-api-go-sdk/pkg/model"
)

type VideosAddExample struct {
	TAds          *ads.SDKClient
	AccessToken   string
	AccountId     int64    `json:"account_id"` // 推广帐号 id，有操作权限的帐号 id，包括代理商和广告主帐号 id
	VideoFile     *os.File `json:"video_file"` // 被上传的视频文件，视频二进制流，支持上传的视频文件类型为：mp4、mov、avi, 最大支持 100M 视频上传
	Signature     string   `json:"signature"`  // 视频文件签名  字段长度为 32 字节
	VideosAddOpts *api.VideosAddOpts
}

func (e *VideosAddExample) Init() {
	e.AccessToken = "77ae53631c5aa47bbf97f709fe920aa1"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.AccountId = 30492333
	videoUrl := "https://ark-oss.bettagames.com/2023-03/6211dece8c857d797b571c0012eb77c0.mp4"
	fileBytes, err := getFileBytes(videoUrl)
	if err != nil {
		fmt.Println("getFileBytes videoUrl err:", err)
	}

	//reader := bytes.NewReader(fileBytes)
	file, err := bytesToFile(fileBytes)
	if err != nil {
		fmt.Println("bytesToFile err", err)
		return
	}

	//file, err := os.Open("YOUR FILE PATH")
	//if err != nil {
	e.VideoFile = file
	//}

	e.Signature = "6211dece8c857d797b571c0012eb77c0"
	e.VideosAddOpts = &api.VideosAddOpts{
		Description:          optional.NewString("ddd"),
		AdcreativeTemplateId: optional.Int64{},
	}
}

func (e *VideosAddExample) RunExample() (model.VideosAddResponseData, http.Header, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.Videos().Add(ctx, e.AccountId, e.VideoFile, e.Signature, e.VideosAddOpts)
}

func main() {
	e := &VideosAddExample{}
	e.Init()
	response, headers, err := e.RunExample()
	if err != nil {
		if resErr, ok := err.(errors.ResponseError); ok {
			errStr, _ := json.Marshal(resErr)
			fmt.Println("Response error:", string(errStr))
		} else {
			fmt.Println("Error:", err)
		}
	}
	fmt.Println("Response data:", response)
	fmt.Println("Headers:", headers)
}

func getFileBytes(netUrl string) ([]byte, error) {
	resp, err := resty.New().R().Get(netUrl)
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}

func bytesToFile(b []byte) (*os.File, error) {
	// 创建临时文件
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		return nil, err
	}

	// 将 []byte 内容写入到临时文件中
	_, err = tmpfile.Write(b)
	if err != nil {
		return nil, err
	}

	// 将文件指针重置到文件开始处
	_, err = tmpfile.Seek(0, 0)
	if err != nil {
		return nil, err
	}

	return tmpfile, nil
}
