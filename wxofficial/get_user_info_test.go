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

package wxofficial

import (
	"fmt"
	"testing"
)

func TestGetUserInfo(t *testing.T) {
	fmt.Println("----------获取用户基本信息----------")
	// 请求接口
	token := ""
	openId := ""
	user, err := GetUserInfo(token, openId, "")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("返回值: %+v\n", user)
}
