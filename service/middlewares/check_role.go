package middlewares

import (
	"strconv"
	"strings"

	"github.com/DeniesKresna/brigunaduty/service/extensions/terror"
	"github.com/DeniesKresna/brigunaduty/service/modules/user/usecase"
	"github.com/DeniesKresna/brigunaduty/types/constants"
	"github.com/DeniesKresna/gohelper/utstring"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CheckRole(userCase usecase.UserUsecase, roles []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var terr terror.ErrInterface
		// Get the authorization header from the request
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			terr = terror.ErrInvalidRule("this operation is not allowed")
			responseJson(ctx, terr)
			ctx.Abort()
			return
		}

		// Check if the authorization header starts with "Bearer "
		if !strings.HasPrefix(authHeader, "Bearer ") {
			terr = terror.ErrInvalidRule("this operation is not allowed")
			responseJson(ctx, terr)
			ctx.Abort()
			return
		}

		// Extract the token from the authorization header
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse the token
		token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			// Here you would verify that the signing method is correct
			// and that the key used to sign the token is valid
			// For simplicity, we'll just return an arbitrary key here
			return []byte(utstring.GetEnv(constants.ENV_APP_SECRET, "")), nil
		})

		// Check if there was an error parsing the token
		if err != nil {
			terr = terror.ErrInvalidRule("this operation is not allowed")
			responseJson(ctx, terr)
			ctx.Abort()
			return
		}

		// Check if the token is valid
		if !token.Valid {
			terr = terror.ErrInvalidRule("this operation is not allowed")
			responseJson(ctx, terr)
			ctx.Abort()
			return
		}

		// Get the userId claim from the token
		claims, ok := token.Claims.(*jwt.StandardClaims)
		if !ok {
			terr = terror.ErrInvalidRule("this operation is not allowed")
			responseJson(ctx, terr)
			ctx.Abort()
			return
		}

		// Add the userId claim to the request context
		var userID int64
		userID, err = strconv.ParseInt(claims.Subject, 10, 64)
		if err != nil {
			terr = terror.ErrInvalidRule("this operation is not allowed")
			responseJson(ctx, terr)
			ctx.Abort()
			return
		}

		ctx.Set("user_id", userID)

		userRole, _ := userCase.AuthGetFromContext(ctx)

		isAllowed := false
		for _, role := range roles {
			if role == "" {
				isAllowed = true
				break
			}
			if role == userRole.RoleName {
				isAllowed = true
			}
		}

		if !isAllowed {
			terr = terror.ErrInvalidRule("this operation is not allowed")
			responseJson(ctx, terr)
			ctx.Abort()
			return
		}

		ctx.Set("role_name", userRole.RoleName)

		ctx.Next()
	}
}
