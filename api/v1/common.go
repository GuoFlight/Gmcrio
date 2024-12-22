package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	json "github.com/json-iterator/go"
	"strconv"
)

func handleArgStr(c *gin.Context, argName, argValue string, must bool, defaultValue ...string) (string, bool) {
	// 初始化默认值
	var defaultV string
	if len(defaultValue) > 0 {
		defaultV = defaultValue[0]
	}
	// 若未获取到值
	if argValue == "" {
		if must {
			c.JSON(400, Res{
				Code: 400,
				Msg:  fmt.Sprintf("missing parameter：" + argName),
			})
			return defaultV, false
		} else {
			return defaultV, true
		}
	}
	// 若获取到了值
	return argValue, true
}
func handleArgInt64(c *gin.Context, argName, argValue string, must bool, defaultValue ...int64) (int64, bool) {
	// 初始化默认值
	var defaultV int64
	if len(defaultValue) > 0 {
		defaultV = defaultValue[0]
	}
	// 若未获取到值
	if argValue == "" {
		if must {
			c.JSON(400, Res{
				Code: 400,
				Msg:  fmt.Sprintf("missing parameter：" + argName),
			})
			return -1, false
		} else {
			return defaultV, true
		}
	}
	// 若获取到了值，判断是否为数字
	argInt64, err := strconv.ParseInt(argValue, 10, 64)
	if err != nil {
		c.JSON(400, Res{
			Code: 400,
			Msg:  fmt.Sprintf("invalid arg：%s.It must be an integer.", argName),
		})
		return -1, false
	}
	return argInt64, true
}
func getQueryInt64(c *gin.Context, argName string, must bool, defaultValue ...int64) (int64, bool) {
	argStr := c.DefaultQuery(argName, "")
	return handleArgInt64(c, argName, argStr, must, defaultValue...)
}
func getPostInt64(c *gin.Context, argName string, must bool, defaultValue ...int64) (int64, bool) {
	argStr := c.DefaultPostForm(argName, "")
	return handleArgInt64(c, argName, argStr, must, defaultValue...)
}
func getQueryInt(c *gin.Context, argName string, must bool, defaultValue ...int) (int, bool) {
	// 初始化默认值
	var defaultV int64
	if len(defaultValue) > 0 {
		defaultV = int64(defaultValue[0])
	}
	// 获取参数
	argInt64, ok := getQueryInt64(c, argName, must, defaultV)
	return int(argInt64), ok
}
func getPostInt(c *gin.Context, argName string, must bool, defaultValue ...int) (int, bool) {
	// 初始化默认值
	var defaultV int64
	if len(defaultValue) > 0 {
		defaultV = int64(defaultValue[0])
	}
	// 获取参数
	argInt64, ok := getPostInt64(c, argName, must, defaultV)
	return int(argInt64), ok
}
func getQueryStr(c *gin.Context, argName string, must bool, defaultValue ...string) (string, bool) {
	arg := c.DefaultQuery(argName, "")
	return handleArgStr(c, argName, arg, must, defaultValue...)
}
func getPostStr(c *gin.Context, argName string, must bool, defaultValue ...string) (string, bool) {
	arg := c.DefaultPostForm(argName, "")
	return handleArgStr(c, argName, arg, must, defaultValue...)
}
func getQueryStrings(c *gin.Context, argName string, must bool, defaultValue ...[]string) ([]string, bool) {
	// 初始化默认值
	var defaultV []string
	if len(defaultValue) > 0 {
		defaultV = defaultValue[0]
	}
	// 获取参数值
	arg := c.DefaultQuery(argName, "")
	if arg == "" {
		if must {
			c.JSON(400, Res{
				Code: 400,
				Msg:  fmt.Sprintf("missing parameter：" + argName),
			})
			return defaultV, false
		} else {
			return defaultV, true
		}
	}
	// 解析参数
	var ret []string
	err := json.Unmarshal([]byte(arg), &ret)
	if err != nil {
		if must {
			c.JSON(400, Res{
				Code: 400,
				Msg:  fmt.Sprintf("missing parameter：" + argName),
			})
			return defaultV, false
		} else {
			return defaultV, true
		}
	}
	return ret, true
}
func getPostBool(c *gin.Context, argName string, must bool, defaultValue ...bool) (bool, bool) {
	// 初始化默认值
	var defaultV bool
	if len(defaultValue) > 0 {
		defaultV = defaultValue[0]
	}
	// 获取参数值
	argStr := c.DefaultPostForm(argName, "")
	// 若获取失败
	if argStr == "" {
		if must {
			c.JSON(400, Res{
				Code: 400,
				Msg:  fmt.Sprintf("missing parameter：" + argName),
			})
			return defaultV, false
		} else {
			return defaultV, true
		}
	}
	// 若获取成功，判断是否为bool
	argBool, err := strconv.ParseBool(argStr)
	if err != nil {
		c.JSON(400, Res{
			Code: 400,
			Msg:  fmt.Sprintf("invalid arg：%s.It must be an bool.", argName),
		})
		return defaultV, false
	}
	return argBool, true
}
