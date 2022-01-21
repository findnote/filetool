package utils

import "errors"

type Code string

var (
	//  字段已存在
	ErrFieldExisted = errors.New("field Existed")

	//  请求错误
	ErrBadRequest = errors.New("bad request")

	// 无效数据
	ErrInvalidData = errors.New("invalid data")

	//  记录不存在
	ErrRecordNotFound = errors.New("record not found")
)

const (
	/*
		Canceled         = grpc.Canceled
		Unauthenticated  = grpc.Unauthenticated
		Unavailable      = grpc.Unavailable
		Unimplemented    = grpc.Unimplemented
		Internal         = grpc.Internal
		InvalidArgument  = grpc.InvalidArgument
		PermissionDenied = grpc.PermissionDenied
	*/

	Ok                     Code = "200"
	BadRequest             Code = "400"
	Forbidden              Code = "403" // 禁止的
	NotFound               Code = "404"
	RequestTimeout         Code = "408"
	UnsupportedMediaType   Code = "415" // 不支持的类型
	BusyHere               Code = "486"
	TemporarilyUnavailable Code = "480" // 暂时不可用
	InternalError          Code = "500" // 内部错误
	NotImplemented         Code = "501" // 未实施
	ServiceUnavailable     Code = "503" // 服务不可用
	InvalidToken           Code = "601" // token无效
	InvalidVerification    Code = "602" // 验证无效或出错
	DeniedPermission       Code = "603" // 没有权限
	ExternalServerError    Code = "604" // 调用对外的第三方的api结果错误
	TaskingNotOnly         Code = "605" // 进行中的任务不唯一
	WrongOperated          Code = "606" // 问题操作，可用户主观避免
	RecordNotFound         Code = "607" // 未找到记录
	InvalidData            Code = "608" // 无效数据

)

//
// GetErrMessage
//  @Description: 根据错误码获取错误消息
//
func GetErrMessage(code Code) string {
	switch code {
	case BadRequest:
		return "请求错误"
	case Forbidden:
		return "禁止的"
	case UnsupportedMediaType:
		return "不支持的类型"
	case TemporarilyUnavailable:
		return "暂时不可用"
	case InternalError:
		return "内部错误"
	case NotImplemented:
		return "未实施"
	case ServiceUnavailable:
		return "服务不可用"
	case InvalidToken:
		return "token无效"
	case InvalidVerification:
		return "验证无效或出错"
	case DeniedPermission:
		return "没有权限"
	case ExternalServerError:
		return "调用对外的第三方的api结果出错"
	case TaskingNotOnly:
		return "进行中的任务不唯一"
	case RecordNotFound:
		return "未找到记录"
	case InvalidData:
		return "无效数据"
	//case DisableDelete:
	//	return "存在关联，不可被删除"
	default:
		return ""
	}
}

//
// Get606ErrMsg
//  @Description: 根据错误原因获取错误Msg返回值
//
func Get606ErrMsg(reason string) string {
	switch reason {
	case "assetNo":
		return "资产编号已存在"
	case "equipNo":
		return "设备编号已存在"
	case "equipRFId":
		return "设备RFID已存在"
	case "equipQRCode":
		return "二维码已存在"
	case "cabinetNo":
		return "机柜编号已存在"
	case "dictnary":
		return "字典值已存在"
	case "model":
		return "该型号已存在"
	case "category":
		return "该工器具类型已存在"
	case "role_name":
		return "角色昵称已存在"
	case "role_key":
		return "角色名已存在"
	case "wrongUserPwd":
		return "账号或密码错误"
	case "wrongOldPwd":
		return "原密码错误"
	case "roomNo":
		return "库房编号已存在"
	case "roomName":
		return "库房名称已存在"
	default:
		return ""
	}
}
