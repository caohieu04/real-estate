package component

import (
	"github.com/caohieu04/real-estate/component/upcloud"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() upcloud.UploadProvider
}

type appCtx struct {
	db             *gorm.DB
	uploadProvider upcloud.UploadProvider
}

func NewAppContext(db *gorm.DB, uploadProvider upcloud.UploadProvider) *appCtx {
	return &appCtx{
		db:             db,
		uploadProvider: uploadProvider,
	}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) UploadProvider() upcloud.UploadProvider {
	return ctx.uploadProvider
}
