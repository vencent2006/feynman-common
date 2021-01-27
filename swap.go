/**
 * @Author: vincent
 * @Description:
 * @File:  swap
 * @Version: 1.0.0
 * @Date: 2021/1/15 07:52
 */

package common

import "encoding/json"

// 通过json tag 进行结构体赋值
func SwapTo(request, response interface{}) (err error) {
	dataByte, err := json.Marshal(request)
	if err != nil {
		return
	}
	return json.Unmarshal(dataByte, response)
}
