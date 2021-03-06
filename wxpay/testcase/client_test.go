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

package testcase

import (
	"gitee.com/xiaochengtech/wechat/wxpay"
	"os"
)

var (
	testAppId    = os.Getenv("AppID")
	testSubAppId = os.Getenv("SubAppID")
	testMchId    = os.Getenv("MchID")
	testSubMchId = os.Getenv("SubMchID")
	testApiKey   = os.Getenv("ApiKey")
	testCertPath = os.Getenv("CertFilepath")
)

var testClient = wxpay.NewClient(false, wxpay.ServiceTypeFacilitatorDomestic, testApiKey, testCertPath, wxpay.Config{
	AppId:    testAppId,
	SubAppId: testSubAppId,
	MchId:    testMchId,
	SubMchId: testSubMchId,
})
