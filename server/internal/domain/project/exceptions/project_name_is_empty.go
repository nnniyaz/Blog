package exceptions

import "github.com/nnniyaz/blog/pkg/core"

var ErrProjectNameIsEmpty = core.NewI18NError(core.EINVALID, core.TXT_NAME_OF_PROJECT_IS_EMPTY)