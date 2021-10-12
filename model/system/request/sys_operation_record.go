package request

import (
	"devops-api/model/common/request"
	"devops-api/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
