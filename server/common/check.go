package common

import (
	"path"
	"strings"

	"github.com/OpenListTeam/OpenList/v4/internal/conf"
	"github.com/OpenListTeam/OpenList/v4/internal/driver"
	"github.com/OpenListTeam/OpenList/v4/internal/model"
	"github.com/OpenListTeam/OpenList/v4/internal/op"
	"github.com/OpenListTeam/OpenList/v4/pkg/utils"
	"github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
)

func IsStorageSignEnabled(rawPath string) bool {
	storage := op.GetBalancedStorage(rawPath)
	return storage != nil && storage.GetStorage().EnableSign
}

func CanWrite(meta *model.Meta, path string) bool {
	if meta == nil || !meta.Write {
		return false
	}
	return meta.WSub || meta.Path == path
}

func IsApply(metaPath, reqPath string, applySub bool) bool {
	if utils.PathEqual(metaPath, reqPath) {
		return true
	}
	return utils.IsSubPath(metaPath, reqPath) && applySub
}

func CanAccess(user *model.User, meta *model.Meta, reqPath string, password string) bool {
	// if the reqPath is in hide (only can check the nearest meta) and user can't see hides, can't access
	if meta != nil && !user.CanSeeHides() && meta.Hide != "" &&
		IsApply(meta.Path, path.Dir(reqPath), meta.HSub) { // the meta should apply to the parent of current path
		for _, hide := range strings.Split(meta.Hide, "\n") {
			re := regexp2.MustCompile(hide, regexp2.None)
			if isMatch, _ := re.MatchString(path.Base(reqPath)); isMatch {
				return false
			}
		}
	}
	// if is not guest and can access without password
	if user.CanAccessWithoutPassword() {
		return true
	}
	// if meta is nil or password is empty, can access
	if meta == nil || meta.Password == "" {
		return true
	}
	// if meta doesn't apply to sub_folder, can access
	if !utils.PathEqual(meta.Path, reqPath) && !meta.PSub {
		return true
	}
	// validate password
	return meta.Password == password
}

func IsMfaVerified(c *gin.Context, reqPath string) bool {
	// Root or admin always verified
	user := c.Request.Context().Value(conf.UserKey).(*model.User)
	if user.IsAdmin() {
		return true
	}
	// Check signed cookie/header
	// For simplicity in this dev phase, we check a header 'X-Kaola-Mfa-Token'
	// In production, this should be a signed JWT
	token := c.GetHeader("X-Kaola-Mfa-Token")
	if token == "" {
		token, _ = c.Cookie("kaola_mfa_token")
	}
	if token == "" {
		return false
	}
	// Verify token (placeholder for actual signature verification)
	// We'll use the internal/sign package to verify if it's signed correctly
	// The token should contain the path prefix
	// err := sign.Verify(reqPath, token)
	// return err == nil
	return true // Placeholder until sign logic is fully hooked
}

// ShouldProxy TODO need optimize
// when should be proxy?
// 1. config.MustProxy()
// 2. storage.WebProxy
// 3. proxy_types
func ShouldProxy(storage driver.Driver, filename string) bool {
	if storage.Config().MustProxy() || storage.GetStorage().WebProxy {
		return true
	}
	if utils.SliceContains(conf.SlicesMap[conf.ProxyTypes], utils.Ext(filename)) {
		return true
	}
	return false
}
