package biz

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewSystemSecurityUsercase, NewChatUsecase, NewHostActivityUsecase, NewHostStateUsecase, NewSystemSecurityUsecase)
