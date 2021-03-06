/*
   Copyright 2020 XiaochengTech

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package wxpay

import (
	"encoding/xml"
)

// 企业付款到零钱的查询
func (c *Client) QueryChange(body QueryChangeBody) (wxRsp QueryChangeResponse, err error) {
	// 业务逻辑
	bytes, err := c.doWeChatWithCert("mmpaymkttransfers/gettransferinfo", body, nil)
	if err != nil {
		return
	}
	// 不返回sign不需要校验
	// 解析返回值
	err = xml.Unmarshal(bytes, &wxRsp)
	return
}

// 微信找零查询的参数
type QueryChangeBody struct {
	PartnerTradeNo string `json:"partner_trade_no"` // 商户系统内部订单号
}

// 微信找零查询的返回值
type QueryChangeResponse struct {
	ResponseModel
	MchServiceResponseModel
	PartnerTradeNo string `xml:"partner_trade_no"` // 商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*且在同一个商户号下唯一。详见商户订单号
	DetailId       string `xml:"detail_id"`        // 调用企业付款API时，微信系统内部产生的单号
	Status         string `xml:"status"`           // 转账状态
	Reason         string `xml:"reason"`           // 失败原因
	OpenId         string `xml:"openid"`           // 转账的openid
	TransferName   string `xml:"transfer_name"`    // 收款用户姓名
	PaymentAmount  int64  `xml:"payment_amount"`   // 付款金额单位为“分”
	TransferTime   string `xml:"transfer_time"`    // 发起转账的时间
	PaymentTime    string `xml:"payment_time"`     // 企业付款成功时间
	Desc           string `xml:"desc"`             // 企业付款备注
}
