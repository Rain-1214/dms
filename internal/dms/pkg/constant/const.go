package constant

import (
	"fmt"

	dmsCommonV1 "github.com/actiontech/dms/pkg/dms-common/api/dms/v1"
)

// internel build-in uid
const (
	UIDOfOpPermissionCreateProject        = "700001"
	UIDOfOpPermissionProjectAdmin         = "700002"
	UIDOfOpPermissionCreateWorkflow       = "700003"
	UIDOfOpPermissionAuditWorkflow        = "700004"
	UIDOfOpPermissionAuthDBServiceData    = "700005"
	UIDOfOpPermissionExecuteWorkflow      = "700006"
	UIDOfOpPermissionViewOthersWorkflow   = "700007"
	UIDOfOpPermissionViewOthersAuditPlan  = "700008"
	UIDOfOpPermissionSaveAuditPlan        = "700009"
	UIDOfOpPermissionSQLQuery             = "700010"
	UIDOfOpPermissionExportApprovalReject = "700011"
	UIDOfOpPermissionExportCreate         = "700012"

	UIDOfDMSConfig = "700100"

	UIDOfUserAdmin = "700200"
	UIDOfUserSys   = "700201"

	UIDOfProjectDefault = "700300"

	UIDOfRoleProjectAdmin   = "700400"
	UIDOfRoleSQLEAdmin      = "700401"
	UIDOfRoleProvisionAdmin = "700402"
)

func ConvertPermissionIdToType(opPermissionUid string) (apiOpPermissionTyp dmsCommonV1.OpPermissionType, err error) {
	switch opPermissionUid {
	case UIDOfOpPermissionCreateWorkflow:
		apiOpPermissionTyp = dmsCommonV1.OpPermissionTypeCreateWorkflow
	case UIDOfOpPermissionAuditWorkflow:
		apiOpPermissionTyp = dmsCommonV1.OpPermissionTypeAuditWorkflow
	case UIDOfOpPermissionAuthDBServiceData:
		apiOpPermissionTyp = dmsCommonV1.OpPermissionTypeAuthDBServiceData
	case UIDOfOpPermissionProjectAdmin:
		apiOpPermissionTyp = dmsCommonV1.OpPermissionTypeProjectAdmin
	case UIDOfOpPermissionCreateProject:
		apiOpPermissionTyp = dmsCommonV1.OpPermissionTypeCreateProject
	case UIDOfOpPermissionExecuteWorkflow:
		apiOpPermissionTyp = dmsCommonV1.OpPermissionTypeExecuteWorkflow
	case UIDOfOpPermissionViewOthersWorkflow:
		apiOpPermissionTyp = dmsCommonV1.OpPermissionTypeViewOthersWorkflow
	case UIDOfOpPermissionSaveAuditPlan:
		apiOpPermissionTyp = dmsCommonV1.OpPermissionTypeSaveAuditPlan
	case UIDOfOpPermissionViewOthersAuditPlan:
		apiOpPermissionTyp = dmsCommonV1.OpPermissionTypeViewOtherAuditPlan
	case UIDOfOpPermissionSQLQuery:
		apiOpPermissionTyp = dmsCommonV1.OpPermissionTypeSQLQuery
	case UIDOfOpPermissionExportApprovalReject:
		apiOpPermissionTyp = dmsCommonV1.OpPermissionTypeAuditExportWorkflow
	case UIDOfOpPermissionExportCreate:
		apiOpPermissionTyp = dmsCommonV1.OpPermissionTypeExportCreate
	default:
		return dmsCommonV1.OpPermissionTypeUnknown, fmt.Errorf("get user op permission type error: invalid op permission uid: %v", opPermissionUid)

	}
	return apiOpPermissionTyp, nil
}

func ConvertPermissionTypeToId(opPermissionType dmsCommonV1.OpPermissionType) (permissionId string, err error) {
	switch opPermissionType {
	case dmsCommonV1.OpPermissionTypeCreateWorkflow:
		permissionId = UIDOfOpPermissionCreateWorkflow
	case dmsCommonV1.OpPermissionTypeAuditWorkflow:
		permissionId = UIDOfOpPermissionAuditWorkflow
	case dmsCommonV1.OpPermissionTypeAuthDBServiceData:
		permissionId = UIDOfOpPermissionAuthDBServiceData
	case dmsCommonV1.OpPermissionTypeProjectAdmin:
		permissionId = UIDOfOpPermissionProjectAdmin
	case dmsCommonV1.OpPermissionTypeCreateProject:
		permissionId = UIDOfOpPermissionCreateProject
	case dmsCommonV1.OpPermissionTypeExecuteWorkflow:
		permissionId = UIDOfOpPermissionExecuteWorkflow
	case dmsCommonV1.OpPermissionTypeViewOthersWorkflow:
		permissionId = UIDOfOpPermissionViewOthersWorkflow
	case dmsCommonV1.OpPermissionTypeSaveAuditPlan:
		permissionId = UIDOfOpPermissionSaveAuditPlan
	case dmsCommonV1.OpPermissionTypeViewOtherAuditPlan:
		permissionId = UIDOfOpPermissionViewOthersAuditPlan
	case dmsCommonV1.OpPermissionTypeSQLQuery:
		permissionId = UIDOfOpPermissionSQLQuery
	case dmsCommonV1.OpPermissionTypeAuditExportWorkflow:
		permissionId = UIDOfOpPermissionExportApprovalReject
	case dmsCommonV1.OpPermissionTypeExportCreate:
		permissionId = UIDOfOpPermissionExportCreate
	default:
		return "", fmt.Errorf("get user op permission id error: invalid op permission type: %v", opPermissionType)
	}

	return permissionId, nil
}

type DBType string

func (d DBType) String() string {
	return string(d)
}

func ParseDBType(s string) (DBType, error) {
	switch s {
	case "MySQL":
		return DBTypeMySQL, nil
	case "TDSQL For InnoDB":
		return DBTypeTDSQLForInnoDB, nil
	case "TiDB":
		return DBTypeTiDB, nil
	case "PostgreSQL":
		return DBTypePostgreSQL, nil
	case "Oracle":
		return DBTypeOracle, nil
	case "DB2":
		return DBTypeDB2, nil
	case "SQL Server":
		return DBTypeSQLServer, nil
	case "OceanBase For MySQL":
		return DBTypeOceanBaseMySQL, nil
	case "GoldenDB":
		return DBTypeGoldenDB, nil
	default:
		return "", fmt.Errorf("invalid db type: %s", s)
	}
}

const (
	DBTypeMySQL          DBType = "MySQL"
	DBTypePostgreSQL     DBType = "PostgreSQL"
	DBTypeTiDB           DBType = "TiDB"
	DBTypeSQLServer      DBType = "SQL Server"
	DBTypeOracle         DBType = "Oracle"
	DBTypeDB2            DBType = "DB2"
	DBTypeOceanBaseMySQL DBType = "OceanBase For MySQL"
	DBTypeTDSQLForInnoDB DBType = "TDSQL For InnoDB"
	DBTypeGoldenDB       DBType = "GoldenDB"
)

type FilterCondition struct {
	// Filter For Preload Table
	Table         string
	KeywordSearch bool
	Field         string
	Operator      FilterOperator
	Value         interface{}
}

type FilterOperator string

const (
	FilterOperatorEqual              FilterOperator = "="
	FilterOperatorIsNull             FilterOperator = "isNull"
	FilterOperatorNotEqual           FilterOperator = "<>"
	FilterOperatorContains           FilterOperator = "like"
	FilterOperatorNotContains        FilterOperator = "not like"
	FilterOperatorGreaterThanOrEqual FilterOperator = ">="
	FilterOperatorLessThanOrEqual    FilterOperator = "<="
	FilterOperatorIn                 FilterOperator = "in"
)

type DBServiceSourceName string

const (
	DBServiceSourceNameDMP  DBServiceSourceName = "Actiontech DMP"
	DBServiceSourceNameDMS  DBServiceSourceName = "Actiontech DMS"
	DBServiceSourceNameSQLE DBServiceSourceName = "SQLE"
)

func ParseDBServiceSource(s string) (DBServiceSourceName, error) {
	switch s {
	case string(DBServiceSourceNameDMP):
		return DBServiceSourceNameDMP, nil
	case string(DBServiceSourceNameDMS):
		return DBServiceSourceNameDMS, nil
	case string(DBServiceSourceNameSQLE):
		return DBServiceSourceNameSQLE, nil
	default:
		return "", fmt.Errorf("invalid data object source name: %s", s)
	}
}

const (
	DMSToken = "dms-token"
)
