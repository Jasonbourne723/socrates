package global

type CustomError struct {
	ErrorCode int
	ErrorMsg  string
}

type CustomErrors struct {
	BusinessError       CustomError
	ValidateError       CustomError
	TokenError          CustomError
	CodeDuplicateError  CustomError
	NameDuplicateError  CustomError
	RecordNotFoundError CustomError
	MobileExistedError  CustomError
}

var Errors = CustomErrors{
	BusinessError:       CustomError{40000, "业务错误"},
	ValidateError:       CustomError{40001, "请求参数错误"},
	TokenError:          CustomError{40002, "登录授权失效"},
	CodeDuplicateError:  CustomError{40003, "编号重复错误"},
	NameDuplicateError:  CustomError{40004, "名称重复错误"},
	RecordNotFoundError: CustomError{40005, "记录未找到错误"},
	MobileExistedError:  CustomError{40006, "手机号已存在"},
}

func (err CustomError) Error() string {
	return err.ErrorMsg
}
