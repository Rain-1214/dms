package v1

import (
	"fmt"

	base "github.com/actiontech/dms/pkg/dms-common/api/base/v1"
)

// swagger:parameters GetUser
type GetUserReq struct {
	// user uid
	// in:path
	UserUid string `param:"user_uid" json:"user_uid" validate:"required"`
}

// A dms user
type GetUser struct {
	// user uid
	UserUid string `json:"uid"`
	// user name
	Name string `json:"name"`
	// user email
	Email string `json:"email"`
	// user phone
	Phone string `json:"phone"`
	// user wxid
	WxID string `json:"wxid"`
	// user stat
	Stat Stat `json:"stat"`
	// user authentication type
	AuthenticationType UserAuthenticationType `json:"authentication_type"`
	// user groups
	UserGroups []UidWithName `json:"user_groups"`
	// user operation permissions
	OpPermissions []UidWithName `json:"op_permissions"`
	// is admin
	IsAdmin bool `json:"is_admin"`
	// user bind name space
	UserBindNamespaces []UserBindNamespace `json:"user_bind_namespaces"`
}

type UserBindNamespace struct {
	NamespaceID   string `json:"namespace_id"`
	NamespaceName string `json:"namespace_name"`
	IsManager     bool   `json:"is_manager"`
}

// swagger:enum Stat
type Stat string

const (
	StatOK      Stat = "正常"
	StatDisable Stat = "被禁用"
	StatUnknown Stat = "未知"
)

type UidWithName struct {
	Uid  string `json:"uid"`
	Name string `json:"name"`
}

// swagger:enum UserAuthenticationType
type UserAuthenticationType string

const (
	UserAuthenticationTypeLDAP    UserAuthenticationType = "ldap"   // user verify through ldap
	UserAuthenticationTypeDMS     UserAuthenticationType = "dms"    // user verify through dms
	UserAuthenticationTypeOAUTH2  UserAuthenticationType = "oauth2" // user verify through oauth2
	UserAuthenticationTypeUnknown UserAuthenticationType = "unknown"
)

// swagger:model GetUserReply
type GetUserReply struct {
	// Get user reply
	Payload struct {
		User *GetUser `json:"user"`
	} `json:"payload"`

	// Generic reply
	base.GenericResp
}

// swagger:parameters GetUserOpPermission
type GetUserOpPermissionReq struct {
	// user uid
	// in:path
	UserUid string `param:"user_uid" json:"user_uid" validate:"required"`
	// user op permission info
	// in:body
	UserOpPermission *UserOpPermission `json:"user_op_permission" validate:"required"`
}

type UserOpPermission struct {
	// uesr namespace uid
	NamespaceUid string `json:"namespace_uid" validate:"required"`
}

// swagger:model GetUserOpPermissionReply
type GetUserOpPermissionReply struct {
	// user op permission reply
	Payload struct {
		// is user admin, admin has all permissions
		IsAdmin bool `json:"is_admin"`
		// user op permissions
		OpPermissionList []OpPermissionItem `json:"op_permission_list"`
	} `json:"payload"`

	// Generic reply
	base.GenericResp
}

type OpPermissionItem struct {
	// object uids, object type is defined by RangeType
	RangeUids []string `json:"range_uids"`
	// object type of RangeUids
	RangeType OpRangeType `json:"range_type"`
	// op permission type
	OpPermissionType OpPermissionType `json:"op_permission_type"`
}

// swagger:enum OpRangeType
type OpRangeType string

const (
	OpRangeTypeUnknown OpRangeType = "unknown"
	// 全局权限: 该权限只能被用户使用
	OpRangeTypeGlobal OpRangeType = "global"
	// 空间权限: 该权限只能被成员使用
	OpRangeTypeNamespace OpRangeType = "namespace"
	// 空间内的数据源权限: 该权限只能被成员使用
	OpRangeTypeDBService OpRangeType = "db_service"
)

func ParseOpRangeType(typ string) (OpRangeType, error) {
	switch typ {
	case string(OpRangeTypeDBService):
		return OpRangeTypeDBService, nil
	case string(OpRangeTypeNamespace):
		return OpRangeTypeNamespace, nil
	case string(OpRangeTypeGlobal):
		return OpRangeTypeGlobal, nil
	default:
		return "", fmt.Errorf("invalid op range type: %s", typ)
	}
}

// swagger:enum OpPermissionType
type OpPermissionType string

const (
	OpPermissionTypeUnknown OpPermissionType = "unknown"
	// 创建空间；创建空间的用户自动拥有该空间管理权限
	OpPermissionTypeCreateNamespace OpPermissionType = "create_namespace"
	// 空间管理；拥有该权限的用户可以管理空间下的所有资源
	OpPermissionTypeNamespaceAdmin OpPermissionType = "namespace_admin"
	// 创建/编辑工单；拥有该权限的用户可以创建/编辑工单
	OpPermissionTypeCreateWorkflow OpPermissionType = "create_workflow"
	// 审核/驳回工单；拥有该权限的用户可以审核/驳回工单
	OpPermissionTypeAuditWorkflow OpPermissionType = "audit_workflow"
	// 授权数据源数据权限；拥有该权限的用户可以授权数据源数据权限
	OpPermissionTypeAuthDBServiceData OpPermissionType = "auth_db_service_data"
	// 查看其他工单权限
	OpPermissionTypeViewOthersWorkflow OpPermissionType = "view_others_workflow"
	// 上线工单；拥有该权限的用户可以上线工单
	OpPermissionTypeExecuteWorkflow OpPermissionType = "execute_workflow"
	// 查看其他扫描任务权限
	OpPermissionTypeViewOtherAuditPlan OpPermissionType = "view_other_audit_plan"
	// 创建扫描任务权限；拥有该权限的用户可以创建/更新扫描任务
	OpPermissionTypeSaveAuditPlan OpPermissionType = "save_audit_plan"
	//SQL查询；SQL查询权限
	OpPermissionTypeSQLQuery OpPermissionType = "sql_query"
)

func ParseOpPermissionType(typ string) (OpPermissionType, error) {
	switch typ {
	case string(OpPermissionTypeCreateNamespace):
		return OpPermissionTypeCreateNamespace, nil
	case string(OpPermissionTypeNamespaceAdmin):
		return OpPermissionTypeNamespaceAdmin, nil
	case string(OpPermissionTypeCreateWorkflow):
		return OpPermissionTypeCreateWorkflow, nil
	case string(OpPermissionTypeAuditWorkflow):
		return OpPermissionTypeAuditWorkflow, nil
	case string(OpPermissionTypeAuthDBServiceData):
		return OpPermissionTypeAuthDBServiceData, nil
	case string(OpPermissionTypeViewOthersWorkflow):
		return OpPermissionTypeViewOthersWorkflow, nil
	case string(OpPermissionTypeExecuteWorkflow):
		return OpPermissionTypeExecuteWorkflow, nil
	case string(OpPermissionTypeViewOtherAuditPlan):
		return OpPermissionTypeViewOtherAuditPlan, nil
	case string(OpPermissionTypeSaveAuditPlan):
		return OpPermissionTypeSaveAuditPlan, nil
	case string(OpPermissionTypeSQLQuery):
		return OpPermissionTypeSQLQuery, nil
	default:
		return "", fmt.Errorf("invalid op permission type: %s", typ)
	}
}

func GetOperationTypeDesc(opType OpPermissionType) string {
	switch opType {
	case OpPermissionTypeUnknown:
		return "未知操作类型"
	case OpPermissionTypeCreateNamespace:
		return "创建空间"
	case OpPermissionTypeNamespaceAdmin:
		return "空间管理"
	case OpPermissionTypeCreateWorkflow:
		return "创建/编辑工单"
	case OpPermissionTypeAuditWorkflow:
		return "审核/驳回工单；拥有该权限的用户可以审核/驳回工单"
	case OpPermissionTypeAuthDBServiceData:
		return "授权数据源数据权限"
	case OpPermissionTypeViewOthersWorkflow:
		return "查看其他工单权限"
	case OpPermissionTypeExecuteWorkflow:
		return "上线工单"
	case OpPermissionTypeViewOtherAuditPlan:
		return "查看其他扫描任务权限"
	case OpPermissionTypeSaveAuditPlan:
		return "创建扫描任务权限"
	case OpPermissionTypeSQLQuery:
		return "SQL查询"
	default:
		return "不支持的操作类型"
	}
}

// swagger:parameters ListUsers
type ListUserReq struct {
	// the maximum count of user to be returned
	// in:query
	// Required: true
	PageSize uint32 `query:"page_size" json:"page_size" validate:"required"`
	// the offset of users to be returned, default is 0
	// in:query
	PageIndex uint32 `query:"page_index" json:"page_index"`
	// Multiple of ["name"], default is ["name"]
	// in:query
	OrderBy UserOrderByField `query:"order_by" json:"order_by"`
	// filter the user name
	// in:query
	FilterByName string `query:"filter_by_name" json:"filter_by_name"`
	// filter the user uids
	// in:query
	FilterByUids string `query:"filter_by_uids" json:"filter_by_uids"`
	// filter deleted user to be return ,default is false
	// in:query
	FilterDeletedUser bool `query:"filter_del_user" json:"filter_del_user"`
}

// swagger:enum UserOrderByField
type UserOrderByField string

const (
	UserOrderByName UserOrderByField = "name"
)

// A dms user
type ListUser struct {
	// user uid
	UserUid string `json:"uid"`
	// user name
	Name string `json:"name"`
	// user email
	Email string `json:"email"`
	// user phone
	Phone string `json:"phone"`
	// user wxid
	WxID string `json:"wxid"`
	// user stat
	Stat Stat `json:"stat"`
	// user authentication type
	AuthenticationType UserAuthenticationType `json:"authentication_type"`
	// user groups
	UserGroups []UidWithName `json:"user_groups"`
	// user operation permissions
	OpPermissions []UidWithName `json:"op_permissions"`
	// user is deleted
	IsDeleted bool `json:"is_deleted"`
}

// swagger:model ListUserReply
type ListUserReply struct {
	// List user reply
	Payload struct {
		Users []*ListUser `json:"users"`
		Total int64       `json:"total"`
	} `json:"payload"`

	// Generic reply
	base.GenericResp
}
