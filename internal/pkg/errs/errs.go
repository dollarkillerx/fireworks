package errs

type Error struct {
	HttpCode int    `json:"http_code"`
	Code     string `json:"code"`
	Message  string `json:"message"`
}

var (
	BadRequest = Error{
		Code:    "40001",
		Message: "400 错误的表单",
	}
	PleaseSignIn = Error{
		HttpCode: 401,
		Code:     "401",
		Message:  "401 请重新登录",
	}
	LoginFailed = Error{
		Code:    "4001",
		Message: "登录失败 密码错误 或 用户不存在",
	}
	SqlSystemError = Error{
		Code:    "5002",
		Message: "系统升级中请稍候再试",
	}
	ReptileError = Error{
		Code:    "5002",
		Message: "以记录您的犯罪应为，我司将以非法侵入计算机信息系统 将你起诉",
	}
	SystemError = Error{
		Code:    "5001",
		Message: "系统升级中请稍候再试",
	}
)

func NewError(code string, message string) Error {
	return Error{
		Code:    code,
		Message: message,
	}
}
