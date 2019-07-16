package wechat

// 场景信息模型
type SceneInfoModel struct {
	ID       string `json:"id"`        // 门店唯一标识
	Name     string `json:"name"`      // 门店名称
	AreaCode string `json:"area_code"` // 门店所在地行政区划码，详细见《最新县及县以上行政区划代码》
	Address  string `json:"address"`   // 门店详细地址
}

// 返回结果的通信标识
type ResponseModel struct {
	ReturnCode string `xml:"return_code"` // SUCCESS/FAIL 此字段是通信标识，非交易标识，交易是否成功需要查看result_code来判断
	ReturnMsg  string `xml:"return_msg"`  // 返回信息，如非空，为错误原因：签名失败/参数格式校验错误
	RetMsg     string `xml:"retmsg"`      // 沙盒时返回的错误信息
}

// 业务返回结果的错误信息
type ServiceResponseModel struct {
	ResultCode string `xml:"result_code"`  // SUCCESS/FAIL
	ErrCode    string `xml:"err_code"`     // 详细参见第6节错误列表
	ErrCodeDes string `xml:"err_code_des"` // 错误返回的信息描述
}
