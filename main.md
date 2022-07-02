package main

import (
	"log"
	"os"
	"time"

	"fmt"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/joho/godotenv"
	"heroku.com/controller"
	"heroku.com/model"
)

type login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var identityKey = "id"

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	db := c.MustGet("db").(*gorm.DB)
	// user, _ := c.Get(identityKey)
	var Pengguna model.User
	db.Where("email = ?", claims[identityKey]).Find(&Pengguna)
	c.JSON(200, gin.H{
		"user": claims[identityKey],
		// "userName": user.(*User).UserName,
		"level": Pengguna.Level,
	})
}

func user(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var login login
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(400, gin.H{
			"status":  "Error",
			"message": "Request harus dalam bentuk JSON.",
		})
		return
	}
	var Pengguna model.User
	db.Where("email = ?", login.Email).Find(&Pengguna)
	fmt.Println(Pengguna)
	c.JSON(200, gin.H{
		"user":  Pengguna.Email,
		"level": Pengguna.Level,
	})
}

// User demo
type User struct {
	UserName  string
	FirstName string
	LastName  string
}

func main() {
	r := gin.Default()
	db := model.SetupModels()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	var userLogin model.User

	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "toriq afanudin zone",
		Key:         []byte("jadgdkabcgakayuwahyuni"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				UserName: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.Email
			password := loginVals.Password
			db.Where("email = ?", userID).Where("password = ?", password).Find(&userLogin)

			if userID == userLogin.Email && password == userLogin.Password {
				return &User{
					UserName:  userID,
					LastName:  "Bo-Yi",
					FirstName: "Wu",
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*User); ok && v.UserName == userLogin.Email {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	r.POST("/login", authMiddleware.LoginHandler)
	r.POST("/login2", controller.Login)

	admin := r.Group("/admin")
	admin.GET("/data_pasien", controller.DataPasien)

	dokter := r.Group("/dokter")
	dokter.PUT("/akun_dokter_update/:id", controller.Edit_akun_dokter_by_id)
	dokter.GET("/akun_dokter_lihat", controller.Lihat_akun_dokter)

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	r.GET("/", controller.Utama)
	auth := r.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", helloHandler)

		//DATA PASIEN
		auth.GET("/data_pasien", controller.DataPasien)
		auth.POST("/tambah_data_pasien", controller.Tambah_data_pasien)
		auth.PUT("/edit_data_pasien/:id", controller.Edit_data_pasien)
		auth.DELETE("hapus_data_pasien/:id", controller.Hapus_data_pasien)

		//MANAGE ACCOUNT
		auth.POST("/akun_tambah", controller.TambahAkun)
		auth.GET("/akun_tampil", controller.Akun_tampil)

		//RAWAT JALAN
		auth.GET("/rawat_jalan", controller.Rawat_jalan_lihat)
	}

	godotenv.Load()
	port := os.Getenv("PORT")
	var dns = fmt.Sprintf("127.0.0.1:%s", port)

	r.Run(dns)
}
