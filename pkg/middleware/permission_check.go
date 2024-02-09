package middleware

import (
	"fmt"
	"fullstack_api_test/model"
	"fullstack_api_test/pkg/config"
	pkgerror "fullstack_api_test/pkg/error"
	"fullstack_api_test/pkg/util/responseutil"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"golang.org/x/exp/slices"
)

// Path and permissions mapping
var mapping = map[string][]string{}

func withAppName(names ...string) []string {
	permissions := []string{}
	for _, n := range names {
		permissions = append(permissions, config.Data.AppCode+":"+n)
	}
	return permissions
}

func validateUserPermission(method, path string, claims *model.JwtClaims) pkgerror.CustomError {
	key := method + path
	permissions, keyExists := mapping[key]
	if !keyExists || len(permissions) == 0 {
		// If the mapping does not contain path or the permissions is empty,
		// the request does not need any specific permission
		return pkgerror.NoError
	}
	prefixedPermissions := withAppName(permissions...)
	for _, role := range claims.User.Roles {
		for _, p := range prefixedPermissions {
			if slices.Contains(role.Permissions, p) {
				return pkgerror.NoError
			}
		}
	}
	log.Errorf("None of the required permissions %v found in the JWT", permissions)
	return pkgerror.ErrForbiddenRequest.WithError(fmt.Errorf("missing permissions: %v", permissions))
}

func PermissionCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		c := ctx.Get("jwt_claims")
		if c == nil {
			return next(ctx)
		}
		claims := c.(*model.JwtClaims)
		if claims.User.Superadmin {
			return next(ctx)
		}
		if err := validateUserPermission(ctx.Request().Method, ctx.Path(), claims); !err.IsNoError() {
			return responseutil.SendErrorResponse(ctx, err)
		}
		return next(ctx)
	}
}
