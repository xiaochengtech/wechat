package wechat

import (
	"fmt"
	"os"
	"testing"
)

func TestBasicAccessToken(t *testing.T) {
	fmt.Println("----------获取访问凭证----------")
	// 请求接口
	appId := "wx80adf00e00fecc80"
	appSecret := "fa1c98a5449e909129d08b10c1bbb415"
	token, err := GetBasicAccessToken(appId, appSecret)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("返回值: %+v\n", token)
}

func TestBasicUserInfo(t *testing.T) {
	fmt.Println("----------获取用户基本信息----------")
	// 请求接口
	token := ""
	openId := ""
	user, err := GetBasicUserInfo(token, openId, "")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("返回值: %+v\n", user)
}

func TestAuthAccessToken(t *testing.T) {
	fmt.Println("----------获取授权的access_token----------")
	// 请求接口
	appId := os.Getenv("AppID")
	appSecret := os.Getenv("ApiSecret")
	code := ""
	token, err := GetAuthAccessToken(appId, appSecret, code)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("返回值: %+v\n", token)
}

func TestAuthUserInfo(t *testing.T) {
	fmt.Println("----------获取用户基本信息(授权机制)----------")
	// 请求接口
	token := ""
	openId := ""
	user, err := GetAuthUserInfo(token, openId, "")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("返回值: %+v\n", user)
}

func TestRefreshAuthAccessToken(t *testing.T) {
	fmt.Println("----------刷新授权的access_token----------")
	// 请求接口
	appId := os.Getenv("AppID")
	refresh := ""
	token, err := RefreshAuthAccessToken(appId, refresh)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("返回值: %+v\n", token)
}

func TestCheckAuthAccessToken(t *testing.T) {
	fmt.Println("----------校验授权的access_token----------")
	// 请求接口
	token := ""
	openId := ""
	rst, err := CheckAuthAccessToken(token, openId)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("返回值: %+v\n", rst)
}
