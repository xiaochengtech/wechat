package wechat

import (
	"fmt"
	"testing"
)

// 测试撤销订单
func TestReverse(t *testing.T) {
	fmt.Println("----------撤销订单----------")
	// 初始化参数
	body := ReverseBody{}
	body.OutTradeNo = "YgENQFTovdeJdFouNyy3nFVOhGD6ZvPH"
	// 请求撤销订单
	wxRsp, err := testClient.Reverse(body)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("返回值: %+v\n", wxRsp)
}
