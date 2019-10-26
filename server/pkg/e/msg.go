package e

var MsgFlags = map[int]string{
	SUCCESS:                         "ok",
	ERROR:                           "fail",
	INVALID_PARAMS:                  "请求参数错误",
	ACCOUNT_EXIST:                   "账号已存在",
	REGISTRE_FAIL:                   "注册失败，请稍后再试",
	LOGIN_SUCCESS:                   "登录成功",
	LOGIN_FAIL:                      "登陆失败，请稍后再试",
	REGISTRE_SUCCESS:                "注册成功",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}