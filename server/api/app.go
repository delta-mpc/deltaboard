package api

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	v1 "deltaboard-server/api/v1"
	"deltaboard-server/api/v1/middleware"
	"deltaboard-server/api/v1/response"
	"deltaboard-server/config"
	"deltaboard-server/config/db"
	"deltaboard-server/config/log"
	"deltaboard-server/models"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/loopfz/gadgeto/tonic"
	"github.com/wI2L/fizz"
	"github.com/wI2L/fizz/openapi"
	"gorm.io/gorm"
)

func CheckNCreateAdmin() error {
	fmt.Println("CheckNCreateAdmin")
	user := &models.User{
		Name: "admin",
	}
	if err := db.GetDB().First(user, user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			user = models.NewUser("admin", "admin", "", "")
			user.Role = models.ROLE_ADMIN
			user.ApprvStatus = models.USER_APPROV_STATUS_APPROVED
			user.Create(user, db.GetDB())
			fmt.Println("Creating admin User")
			return nil
		}
	}
	if user.ApprvStatus != models.USER_APPROV_STATUS_APPROVED {
		user.ApprvStatus = models.USER_APPROV_STATUS_APPROVED
		db.GetDB().Save(user)
	}
	return nil
}

func GetHttpApplication(appConfig *config.AppConfig) *gin.Engine {

	// gin.SetMode(appConfig.Environment)
	gin.SetMode("debug")
	engine := gin.New()
	engine.Use(middleware.SetResponseHeader())
	engine.Use(middleware.Cors())
	engine.Use(gin.LoggerWithWriter(log.StandardLogger().Out))
	engine.Use(gin.RecoveryWithWriter(log.StandardLogger().Out))
	engine.Use(APIVersion())

	// Serve static files under static folder
	// for OpenAPI documentations
	engine.Use(static.Serve("/static", static.LocalFile("./static", false)))

	fizzEngine := fizz.NewFromEngine(engine)

	// Do not include package name in component names
	fizzEngine.Generator().UseFullSchemaNames(false)

	// Initialize our own handlers
	tonic.SetErrorHook(TonicResponseErrorHook)
	tonic.SetRenderHook(TonicRenderHook, "")
	tonic.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	// v1 api
	v1.InitRoutes(fizzEngine)
	// Serve OpenAPI specifications
	infos := &openapi.Info{
		Title:       "Go service",
		Description: "A template for Golang API server",
		Version:     "1.0.0",
	}

	fizzEngine.GET("/openapi.json", nil, fizzEngine.OpenAPI(infos, "json"))
	fizzEngine.GET("/openapi.yml", nil, fizzEngine.OpenAPI(infos, "yaml"))

	if len(fizzEngine.Errors()) != 0 {

		for _, err := range fizzEngine.Errors() {
			log.Error(err)
		}

		panic("fizz initialization error")
	}
	CheckNCreateAdmin()
	return engine
}

func APIVersion() gin.HandlerFunc {
	return func(c *gin.Context) {

		path := c.FullPath()

		re := regexp.MustCompile(`^/v([0-9]+)/`)
		matches := re.FindStringSubmatch(path)

		if len(matches) > 1 {
			c.Set("api_version", matches[1])
		}

		c.Next()
	}
}

// Distribute binding & error handling & render handling to implementations in different API versions
func TonicResponseErrorHook(ctx *gin.Context, err error) (int, interface{}) {
	apiVersion := ctx.GetString("api_version")
	switch apiVersion {
	case "1":
		return response.TonicErrorResponse(ctx, err)
	default:
		return tonic.DefaultErrorHook(ctx, err)
	}
}

func TonicRenderHook(ctx *gin.Context, statusCode int, payload interface{}) {
	apiVersion := ctx.GetString("api_version")
	switch apiVersion {
	case "1":
		response.TonicRenderResponse(ctx, statusCode, payload)
	default:
		tonic.DefaultRenderHook(ctx, statusCode, payload)
	}
}
