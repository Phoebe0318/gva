package example

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	CustomerApi
	FileUploadAndDownloadApi
}

var (
	customerService              = service.AppServices.ExampleServiceGroup.CustomerService
	fileUploadAndDownloadService = service.AppServices.ExampleServiceGroup.FileUploadAndDownloadService
)
