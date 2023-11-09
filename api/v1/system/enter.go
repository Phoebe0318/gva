package system

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	DBApi
	JwtApi
	BaseApi
	SystemApi
	CasbinApi
	AutoCodeApi
	SystemApiApi
	AuthorityApi
	DictionaryApi
	AuthorityMenuApi
	OperationRecordApi
	AutoCodeHistoryApi
	DictionaryDetailApi
	AuthorityBtnApi
	ChatGptApi
}

var (
	apiService              = service.AppServices.SystemServiceGroup.ApiService
	jwtService              = service.AppServices.SystemServiceGroup.JwtService
	menuService             = service.AppServices.SystemServiceGroup.MenuService
	userService             = service.AppServices.SystemServiceGroup.UserService
	initDBService           = service.AppServices.SystemServiceGroup.InitDBService
	casbinService           = service.AppServices.SystemServiceGroup.CasbinService
	autoCodeService         = service.AppServices.SystemServiceGroup.AutoCodeService
	baseMenuService         = service.AppServices.SystemServiceGroup.BaseMenuService
	authorityService        = service.AppServices.SystemServiceGroup.AuthorityService
	dictionaryService       = service.AppServices.SystemServiceGroup.DictionaryService
	systemConfigService     = service.AppServices.SystemServiceGroup.SystemConfigService
	operationRecordService  = service.AppServices.SystemServiceGroup.OperationRecordService
	autoCodeHistoryService  = service.AppServices.SystemServiceGroup.AutoCodeHistoryService
	dictionaryDetailService = service.AppServices.SystemServiceGroup.DictionaryDetailService
	authorityBtnService     = service.AppServices.SystemServiceGroup.AuthorityBtnService
	chatGptService          = service.AppServices.SystemServiceGroup.ChatGptService
)
