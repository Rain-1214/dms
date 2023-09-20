// Code generated by gencli for actiontech dms. DO NOT EDIT
package biz

type CloudbeaverConnectionCacheField string

type CloudbeaverUserCacheField string

type DBServiceField string

type DMSConfigField string

type DatabaseSourceServiceField string

type IMConfigurationField string

type LDAPConfigurationField string

type MemberField string

type MemberGroupField string

type MemberGroupRoleOpRangeField string

type MemberRoleOpRangeField string

type ModelField string

type ProjectField string

type Oauth2ConfigurationField string

type OpPermissionField string

type PluginField string

type ProxyTargetField string

type RoleField string

type SMTPConfigurationField string

type UserField string

type UserGroupField string

type WeChatConfigurationField string

type WebHookConfigurationField string

const (
	CloudbeaverConnectionCacheFieldDMSDBServiceID          CloudbeaverConnectionCacheField = "dms_db_service_id"
	CloudbeaverConnectionCacheFieldDMSDBServiceFingerprint CloudbeaverConnectionCacheField = "dms_db_service_fingerprint"
	CloudbeaverConnectionCacheFieldCloudbeaverConnectionID CloudbeaverConnectionCacheField = "cloudbeaver_connection_id"
)

const (
	CloudbeaverUserCacheFieldDMSUserID         CloudbeaverUserCacheField = "dms_user_id"
	CloudbeaverUserCacheFieldDMSFingerprint    CloudbeaverUserCacheField = "dms_fingerprint"
	CloudbeaverUserCacheFieldCloudbeaverUserID CloudbeaverUserCacheField = "cloudbeaver_user_id"
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
	DMSConfigFieldUID                   DMSConfigField = "uid"
	DMSConfigFieldNeedInitOpPermissions DMSConfigField = "need_init_op_permissions"
	DMSConfigFieldNeedInitUsers         DMSConfigField = "need_init_users"
	DMSConfigFieldNeedInitRoles         DMSConfigField = "need_init_roles"
	DMSConfigFieldNeedInitProjects      DMSConfigField = "need_init_projects"
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
	ModelFieldUID ModelField = "uid"
)

const (
	ProjectFieldUID           ProjectField = "uid"
	ProjectFieldName          ProjectField = "name"
	ProjectFieldDesc          ProjectField = "desc"
	ProjectFieldCreateUserUID ProjectField = "create_user_uid"
	ProjectFieldStatus        ProjectField = "status"
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
	ProxyTargetFieldName            ProxyTargetField = "name"
	ProxyTargetFieldUrl             ProxyTargetField = "url"
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
	UserFieldMemberGroup            UserField = "membergroup"
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
)

const (
	WebHookConfigurationFieldUID                  WebHookConfigurationField = "uid"
	WebHookConfigurationFieldEnable               WebHookConfigurationField = "enable"
	WebHookConfigurationFieldMaxRetryTimes        WebHookConfigurationField = "maxretrytimes"
	WebHookConfigurationFieldRetryIntervalSeconds WebHookConfigurationField = "retryintervalseconds"
	WebHookConfigurationFieldEncryptedToken       WebHookConfigurationField = "encryptedtoken"
	WebHookConfigurationFieldURL                  WebHookConfigurationField = "url"
)
