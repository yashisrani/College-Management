package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/yashisrani/Go-Backend/model"
	"github.com/yashisrani/Go-Backend/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Middleware function for token authentication
func (api Apiroutes) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader(model.Authorization)
		if tokenString == "" {
			utils.Log(model.LogLevelInfo, model.ApiPackage,
				model.AuthMiddleware, "token string empty", nil)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization token"})
			c.Abort()
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return model.SecretKey, nil
		})
		if err != nil || !token.Valid {
			utils.Log(model.LogLevelError, model.ApiPackage,
				model.AuthMiddleware, "token value is not valid", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
			c.Abort()
			return
		}
		utils.Log(model.LogLevelInfo, model.ApiPackage,
			model.AuthMiddleware, "token value parsed", token)

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok {
			email := claims["email"].(string)
			password := claims["password"].(string)
			utils.Log(model.LogLevelInfo, model.ApiPackage,
				model.AuthMiddleware, "token value for email and password", email+" password = "+password)
			db, err := gorm.Open(postgres.Open(model.DNS), &gorm.Config{})
			defer func() {
				sqldb, err := db.DB()
				if err != nil {
					utils.Log(model.LogLevelError, model.ApiPackage,
						model.AuthMiddleware, "error in geting sql object for middleware", nil)
				}
				sqldb.Close()
				utils.Log(model.LogLevelInfo, model.ApiPackage,
					model.AuthMiddleware, "middleware db connection closed", nil)
			}()
			if err != nil {
				utils.Log(model.LogLevelError, model.ApiPackage, model.NewStore,
					"error while connecting database", err)
				panic(err)
			}
			var user model.User
			utils.Log(model.LogLevelInfo, model.ApiPackage, model.AuthMiddleware,
				"reading user data from db based on email", email+" pass = "+password)
			resp := db.Where("email = ? AND password = ?", email, password).First(&user)
			if resp.Error != nil {
				utils.Log(model.LogLevelError, model.ApiPackage, model.AuthMiddleware,
					"error while reading user data", resp.Error)
				c.JSON(http.StatusUnauthorized, gin.H{"error": "error in geting user data from Database"})
				c.Abort()
				return
			}
			if !validateUserCredability(user, c.Request.URL.Path, c.Request.Method) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "user is not valid for request"})
				c.Abort()
				return
			}
			utils.Log(model.LogLevelInfo, model.ApiPackage, model.AuthMiddleware,
				"returning user data", user)
		}
		c.Next()
	}
}

func validateUserCredability(user model.User, url, protocal string) bool {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.AuthMiddleware,
		"validating user credability", user.Email+" Password="+user.Password+" url="+url+" protocal="+protocal)
	if user.Type == model.SuperAdminUser {
		return true
	} else if user.Type == model.AdminUser {

	} else if user.Type == model.NormalUser && protocal == "get" {

	} else {

	}
	return true
}
