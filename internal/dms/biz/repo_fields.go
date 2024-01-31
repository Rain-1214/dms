// Code generated by gencli for actiontech dms. DO NOT EDIT
package biz

type BasicConfigField string

type CloudbeaverConnectionCacheField string

type CloudbeaverUserCacheField string

type ClusterLeaderField string

type ClusterNodeInfoField string

type CompanyNoticeField string

type DBServiceField string

type DMSConfigField string

type DataExportTaskField string

type DataExportTaskRecordField string

type DatabaseSourceServiceField string

type IMConfigurationField string

type LDAPConfigurationField string

type MemberField string

type MemberGroupField string

type MemberGroupRoleOpRangeField string

type MemberRoleOpRangeField string

type ModelField string

type Oauth2ConfigurationField string

type OpPermissionField string

type PluginField string

type ProjectField string

type ProxyTargetField string

type RoleField string

type SMTPConfigurationField string

type UserField string

type UserGroupField string

type WeChatConfigurationField string

type WebHookConfigurationField string

type WorkflowField string

type WorkflowRecordField string

type WorkflowStepField string

const (
	BasicConfigFieldUID   BasicConfigField = "uid"
	BasicConfigFieldLogo  BasicConfigField = "logo"
	BasicConfigFieldTitle BasicConfigField = "title"
)

const (
	CloudbeaverConnectionCacheFieldDMSDBServiceID          CloudbeaverConnectionCacheField = "dms_db_service_id"
	CloudbeaverConnectionCacheFieldDMSUserID               CloudbeaverConnectionCacheField = "dms_user_id"
	CloudbeaverConnectionCacheFieldDMSDBServiceFingerprint CloudbeaverConnectionCacheField = "dms_db_service_fingerprint"
	CloudbeaverConnectionCacheFieldCloudbeaverConnectionID CloudbeaverConnectionCacheField = "cloudbeaver_connection_id"
)

const (
	CloudbeaverUserCacheFieldDMSUserID         CloudbeaverUserCacheField = "dms_user_id"
	CloudbeaverUserCacheFieldDMSFingerprint    CloudbeaverUserCacheField = "dms_fingerprint"
	CloudbeaverUserCacheFieldCloudbeaverUserID CloudbeaverUserCacheField = "cloudbeaver_user_id"
)

const (
	ClusterLeaderFieldAnchor       ClusterLeaderField = "anchor"
	ClusterLeaderFieldServerId     ClusterLeaderField = "serverid"
	ClusterLeaderFieldLastSeenTime ClusterLeaderField = "lastseentime"
)

const (
	ClusterNodeInfoFieldServerId     ClusterNodeInfoField = "serverid"
	ClusterNodeInfoFieldHardwareSign ClusterNodeInfoField = "hardwaresign"
	ClusterNodeInfoFieldCreatedAt    ClusterNodeInfoField = "createdat"
)

const (
	CompanyNoticeFieldUID         CompanyNoticeField = "uid"
	CompanyNoticeFieldNoticeStr   CompanyNoticeField = "noticestr"
	CompanyNoticeFieldReadUserIds CompanyNoticeField = "readuserids"
)

const (
	DBServiceFieldUID               DBServiceField = "uid"
	DBServiceFieldName              DBServiceField = "name"
	DBServiceFieldDBType            DBServiceField = "db_type"
	DBServiceFieldHost              DBServiceField = "db_host"
	DBServiceFieldPort              DBServiceField = "db_port"
	DBServiceFieldUser              DBServiceField = "db_user"
	DBServiceFieldPassword          DBServiceField = "db_password"
	DBServiceFieldDesc              DBServiceField = "desc"
	DBServiceFieldBusiness          DBServiceField = "business"
	DBServiceFieldAdditionalParams  DBServiceField = "additionalparams"
	DBServiceFieldSource            DBServiceField = "source"
	DBServiceFieldProjectUID        DBServiceField = "project_uid"
	DBServiceFieldMaintenancePeriod DBServiceField = "maintenanceperiod"
	DBServiceFieldExtraParameters   DBServiceField = "extraparameters"
)

const (
	DMSConfigFieldUID                            DMSConfigField = "uid"
	DMSConfigFieldNeedInitOpPermissions          DMSConfigField = "need_init_op_permissions"
	DMSConfigFieldNeedInitUsers                  DMSConfigField = "need_init_users"
	DMSConfigFieldNeedInitRoles                  DMSConfigField = "need_init_roles"
	DMSConfigFieldNeedInitProjects               DMSConfigField = "need_init_projects"
	DMSConfigFieldEnableSQLResultSetsDataMasking DMSConfigField = "enable_sql_result_sets_data_masking"
)

const (
	DataExportTaskFieldUID                   DataExportTaskField = "uid"
	DataExportTaskFieldDBServiceUid          DataExportTaskField = "dbserviceuid"
	DataExportTaskFieldDatabaseName          DataExportTaskField = "databasename"
	DataExportTaskFieldWorkFlowRecordUid     DataExportTaskField = "workflowrecorduid"
	DataExportTaskFieldExportType            DataExportTaskField = "exporttype"
	DataExportTaskFieldExportFileType        DataExportTaskField = "exportfiletype"
	DataExportTaskFieldExportFileName        DataExportTaskField = "export_file_name"
	DataExportTaskFieldExportStatus          DataExportTaskField = "exportstatus"
	DataExportTaskFieldExportStartTime       DataExportTaskField = "export_start_time"
	DataExportTaskFieldExportEndTime         DataExportTaskField = "export_end_time"
	DataExportTaskFieldCreateUserUID         DataExportTaskField = "create_user_uid"
	DataExportTaskFieldAuditLevel            DataExportTaskField = "auditlevel"
	DataExportTaskFieldDataExportTaskRecords DataExportTaskField = "dataexporttaskrecords"
)

const (
	DataExportTaskRecordFieldNumber           DataExportTaskRecordField = "number"
	DataExportTaskRecordFieldDataExportTaskId DataExportTaskRecordField = "data_export_task_id"
	DataExportTaskRecordFieldExportSQL        DataExportTaskRecordField = "exportsql"
	DataExportTaskRecordFieldExportSQLType    DataExportTaskRecordField = "export_sql_type"
	DataExportTaskRecordFieldExportStatus     DataExportTaskRecordField = "exportstatus"
	DataExportTaskRecordFieldAuditResults     DataExportTaskRecordField = "auditresults"
)

const (
	DatabaseSourceServiceFieldUID                 DatabaseSourceServiceField = "uid"
	DatabaseSourceServiceFieldName                DatabaseSourceServiceField = "name"
	DatabaseSourceServiceFieldSource              DatabaseSourceServiceField = "source"
	DatabaseSourceServiceFieldVersion             DatabaseSourceServiceField = "version"
	DatabaseSourceServiceFieldURL                 DatabaseSourceServiceField = "url"
	DatabaseSourceServiceFieldDbType              DatabaseSourceServiceField = "dbtype"
	DatabaseSourceServiceFieldProjectUID          DatabaseSourceServiceField = "project_uid"
	DatabaseSourceServiceFieldCronExpress         DatabaseSourceServiceField = "cron_express"
	DatabaseSourceServiceFieldLastSyncErr         DatabaseSourceServiceField = "last_sync_err"
	DatabaseSourceServiceFieldLastSyncSuccessTime DatabaseSourceServiceField = "last_sync_success_time"
	DatabaseSourceServiceFieldExtraParameters     DatabaseSourceServiceField = "extraparameters"
)

const (
	IMConfigurationFieldUID         IMConfigurationField = "uid"
	IMConfigurationFieldAppKey      IMConfigurationField = "app_key"
	IMConfigurationFieldAppSecret   IMConfigurationField = "app_secret"
	IMConfigurationFieldIsEnable    IMConfigurationField = "is_enable"
	IMConfigurationFieldProcessCode IMConfigurationField = "process_code"
	IMConfigurationFieldType        IMConfigurationField = "type"
)

const (
	LDAPConfigurationFieldUID                   LDAPConfigurationField = "uid"
	LDAPConfigurationFieldEnable                LDAPConfigurationField = "enable"
	LDAPConfigurationFieldEnableSSL             LDAPConfigurationField = "enablessl"
	LDAPConfigurationFieldHost                  LDAPConfigurationField = "host"
	LDAPConfigurationFieldPort                  LDAPConfigurationField = "port"
	LDAPConfigurationFieldConnectDn             LDAPConfigurationField = "connectdn"
	LDAPConfigurationFieldConnectSecretPassword LDAPConfigurationField = "connectsecretpassword"
	LDAPConfigurationFieldBaseDn                LDAPConfigurationField = "basedn"
	LDAPConfigurationFieldUserNameRdnKey        LDAPConfigurationField = "usernamerdnkey"
	LDAPConfigurationFieldUserEmailRdnKey       LDAPConfigurationField = "useremailrdnkey"
)

const (
	MemberFieldUID              MemberField = "uid"
	MemberFieldUserUID          MemberField = "user_uid"
	MemberFieldProjectUID       MemberField = "project_uid"
	MemberFieldRoleWithOpRanges MemberField = "rolewithopranges"
)

const (
	MemberGroupFieldUID              MemberGroupField = "uid"
	MemberGroupFieldName             MemberGroupField = "name"
	MemberGroupFieldProjectUID       MemberGroupField = "project_uid"
	MemberGroupFieldUsers            MemberGroupField = "users"
	MemberGroupFieldRoleWithOpRanges MemberGroupField = "rolewithopranges"
)

const (
	MemberGroupRoleOpRangeFieldMemberGroupUID MemberGroupRoleOpRangeField = "member_group_uid"
	MemberGroupRoleOpRangeFieldRoleUID        MemberGroupRoleOpRangeField = "role_uid"
	MemberGroupRoleOpRangeFieldOpRangeType    MemberGroupRoleOpRangeField = "op_range_type"
	MemberGroupRoleOpRangeFieldRangeUIDs      MemberGroupRoleOpRangeField = "range_uids"
)

const (
	MemberRoleOpRangeFieldMemberUID   MemberRoleOpRangeField = "member_uid"
	MemberRoleOpRangeFieldRoleUID     MemberRoleOpRangeField = "role_uid"
	MemberRoleOpRangeFieldOpRangeType MemberRoleOpRangeField = "op_range_type"
	MemberRoleOpRangeFieldRangeUIDs   MemberRoleOpRangeField = "range_uids"
)

const (
	ModelFieldUID       ModelField = "uid"
	ModelFieldCreatedAt ModelField = "created_at"
)

const (
	Oauth2ConfigurationFieldUID             Oauth2ConfigurationField = "uid"
	Oauth2ConfigurationFieldEnableOauth2    Oauth2ConfigurationField = "enable_oauth2"
	Oauth2ConfigurationFieldClientID        Oauth2ConfigurationField = "client_id"
	Oauth2ConfigurationFieldClientKey       Oauth2ConfigurationField = "clientkey"
	Oauth2ConfigurationFieldClientSecret    Oauth2ConfigurationField = "clientsecret"
	Oauth2ConfigurationFieldClientHost      Oauth2ConfigurationField = "client_host"
	Oauth2ConfigurationFieldServerAuthUrl   Oauth2ConfigurationField = "server_auth_url"
	Oauth2ConfigurationFieldServerTokenUrl  Oauth2ConfigurationField = "server_token_url"
	Oauth2ConfigurationFieldServerUserIdUrl Oauth2ConfigurationField = "server_user_id_url"
	Oauth2ConfigurationFieldScopes          Oauth2ConfigurationField = "scopes"
	Oauth2ConfigurationFieldAccessTokenTag  Oauth2ConfigurationField = "access_token_tag"
	Oauth2ConfigurationFieldUserIdTag       Oauth2ConfigurationField = "user_id_tag"
	Oauth2ConfigurationFieldUserWeChatTag   Oauth2ConfigurationField = "user_wechat_tag"
	Oauth2ConfigurationFieldUserEmailTag    Oauth2ConfigurationField = "user_email_tag"
	Oauth2ConfigurationFieldLoginTip        Oauth2ConfigurationField = "login_tip"
)

const (
	OpPermissionFieldUID       OpPermissionField = "uid"
	OpPermissionFieldName      OpPermissionField = "name"
	OpPermissionFieldDesc      OpPermissionField = "description"
	OpPermissionFieldRangeType OpPermissionField = "range_type"
)

const (
	PluginFieldName                         PluginField = "name"
	PluginFieldAddDBServicePreCheckUrl      PluginField = "add_db_service_pre_check_url"
	PluginFieldDelDBServicePreCheckUrl      PluginField = "del_db_service_pre_check_url"
	PluginFieldDelUserPreCheckUrl           PluginField = "del_user_pre_check_url"
	PluginFieldDelUserGroupPreCheckUrl      PluginField = "del_user_group_pre_check_url"
	PluginFieldOperateDataResourceHandleUrl PluginField = "operate_data_resource_handle_url"
)

const (
	ProjectFieldUID           ProjectField = "uid"
	ProjectFieldName          ProjectField = "name"
	ProjectFieldDesc          ProjectField = "desc"
	ProjectFieldCreateUserUID ProjectField = "create_user_uid"
	ProjectFieldStatus        ProjectField = "status"
)

const (
	ProxyTargetFieldName            ProxyTargetField = "name"
	ProxyTargetFieldUrl             ProxyTargetField = "url"
	ProxyTargetFieldVersion         ProxyTargetField = "version"
	ProxyTargetFieldProxyUrlPrefixs ProxyTargetField = "proxy_url_prefixs"
)

const (
	RoleFieldUID           RoleField = "uid"
	RoleFieldName          RoleField = "name"
	RoleFieldDesc          RoleField = "description"
	RoleFieldStat          RoleField = "stat"
	RoleFieldOpPermissions RoleField = "oppermissions"
)

const (
	SMTPConfigurationFieldUID              SMTPConfigurationField = "uid"
	SMTPConfigurationFieldEnableSMTPNotify SMTPConfigurationField = "enablesmtpnotify"
	SMTPConfigurationFieldHost             SMTPConfigurationField = "smtp_host"
	SMTPConfigurationFieldPort             SMTPConfigurationField = "smtp_port"
	SMTPConfigurationFieldUsername         SMTPConfigurationField = "smtp_username"
	SMTPConfigurationFieldSecretPassword   SMTPConfigurationField = "secret_smtp_password"
	SMTPConfigurationFieldIsSkipVerify     SMTPConfigurationField = "isskipverify"
)

const (
	UserFieldUID                    UserField = "uid"
	UserFieldName                   UserField = "name"
	UserFieldThirdPartyUserID       UserField = "third_party_user_id"
	UserFieldEmail                  UserField = "email"
	UserFieldPhone                  UserField = "phone"
	UserFieldWeChatID               UserField = "wechat_id"
	UserFieldPassword               UserField = "password"
	UserFieldUserAuthenticationType UserField = "user_authentication_type"
	UserFieldStat                   UserField = "stat"
	UserFieldLastLoginAt            UserField = "last_login_at"
	UserFieldDeletedAt              UserField = "delete_at"
	UserFieldUserGroups             UserField = "usergroups"
	UserFieldOpPermissions          UserField = "oppermissions"
)

const (
	UserGroupFieldUID   UserGroupField = "uid"
	UserGroupFieldName  UserGroupField = "name"
	UserGroupFieldDesc  UserGroupField = "description"
	UserGroupFieldStat  UserGroupField = "stat"
	UserGroupFieldUsers UserGroupField = "users"
)

const (
	WeChatConfigurationFieldUID                 WeChatConfigurationField = "uid"
	WeChatConfigurationFieldEnableWeChatNotify  WeChatConfigurationField = "enablewechatnotify"
	WeChatConfigurationFieldCorpID              WeChatConfigurationField = "corpid"
	WeChatConfigurationFieldEncryptedCorpSecret WeChatConfigurationField = "encryptedcorpsecret"
	WeChatConfigurationFieldAgentID             WeChatConfigurationField = "agentid"
	WeChatConfigurationFieldSafeEnabled         WeChatConfigurationField = "safeenabled"
	WeChatConfigurationFieldProxyIP             WeChatConfigurationField = "proxyip"
)

const (
	WebHookConfigurationFieldUID                  WebHookConfigurationField = "uid"
	WebHookConfigurationFieldEnable               WebHookConfigurationField = "enable"
	WebHookConfigurationFieldMaxRetryTimes        WebHookConfigurationField = "maxretrytimes"
	WebHookConfigurationFieldRetryIntervalSeconds WebHookConfigurationField = "retryintervalseconds"
	WebHookConfigurationFieldEncryptedToken       WebHookConfigurationField = "encryptedtoken"
	WebHookConfigurationFieldURL                  WebHookConfigurationField = "url"
)

const (
	WorkflowFieldUID               WorkflowField = "uid"
	WorkflowFieldName              WorkflowField = "name"
	WorkflowFieldProjectUID        WorkflowField = "project_uid"
	WorkflowFieldWorkflowType      WorkflowField = "workflow_type"
	WorkflowFieldDesc              WorkflowField = "desc"
	WorkflowFieldCreateTime        WorkflowField = "create_time"
	WorkflowFieldCreateUserUID     WorkflowField = "create_user_uid"
	WorkflowFieldWorkflowRecordUid WorkflowField = "workflow_record_uid"
	WorkflowFieldWorkflowRecord    WorkflowField = "workflowrecord"
)

const (
	WorkflowRecordFieldUID         WorkflowRecordField = "uid"
	WorkflowRecordFieldWorkflowUid WorkflowRecordField = "workflowuid"
	WorkflowRecordFieldStatus      WorkflowRecordField = "status"
	WorkflowRecordFieldTaskIds     WorkflowRecordField = "taskids"
	WorkflowRecordFieldSteps       WorkflowRecordField = "steps"
)

const (
	WorkflowStepFieldStepId            WorkflowStepField = "stepid"
	WorkflowStepFieldWorkflowRecordUid WorkflowStepField = "workflowrecorduid"
	WorkflowStepFieldOperationUserUid  WorkflowStepField = "operationuseruid"
	WorkflowStepFieldState             WorkflowStepField = "state"
	WorkflowStepFieldReason            WorkflowStepField = "reason"
	WorkflowStepFieldAssignees         WorkflowStepField = "assignees"
)
