package _gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// http框架gin的使用 https://github.com/gin-gonic/gin


func SampleMainFirst() {

	// 创建默认的engine：注入了默认的中间件
	r := gin.Default()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	r.MaxMultipartMemory = 8 << 20  // 8 MiB

	// get请求
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// post请求
	r.POST("/ping-post", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong-post",
		})
	})

	// 通过url路径传参
	r.GET("/user/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK, "Hello %s", name)
	})

	// 将会匹配 /user/dictx/ 或者 /user/dictx/send
	r.GET("/user/:name/*action", func(context *gin.Context) {
		name := context.Param("name")
		action := context.Param("action")  // action is /send which starts with '/'
		message := name + " is " + action
		context.String(http.StatusOK, message)
	})

	// 通过url问号传参
	r.GET("/welcome", func(context *gin.Context) {
		firstname := context.DefaultQuery("firstname", "Guest")
		lastname := context.Query("lastname")

		context.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	// Multipart/Urlencoded Form
	//  curl -X POST -H"Content-Type:application/x-www-form-urlencoded" -d"message=%E5%A4%A9%E5%AE%89%E9%97%A8&nick=liudehua" http://localhost:8080/form-post
	r.POST("/form-post", func(context *gin.Context) {
		message := context.PostForm("message")
		nick := context.DefaultPostForm("nick", "anonymous")

		context.JSON(http.StatusOK, gin.H{
			"status": "posted",
			"message": message,
			"nick": nick,
		})
	})

	// 混合方式获取参数
	// curl -X POST -H"Content-Type:application/x-www-form-urlencoded" -d"message=hello&name=ldh" http://localhost:8080/form-post-multi?id=12
	r.POST("form-post-multi", func(context *gin.Context) {
		id := context.Query("id")
		page := context.DefaultQuery("page", "1")
		name := context.PostForm("name")
		message := context.PostForm("message")

		context.String(http.StatusOK, "id:" + id + ",page:" + page + ",name:" + name + ",message:" + message)
	})

	// 获取map方式的参数
	// curl -X POST -H"Content-Type:application/x-www-form-urlencoded" -d"names[first]=wang&names[last]=dictx" "http://localhost:8080/post-map?ids\[a\]=12&ids\[b\]=23"
	r.POST("post-map", func(context *gin.Context) {
		ids := context.QueryMap("ids")
		names := context.PostFormMap("names")

		context.JSON(http.StatusOK, gin.H{
			"ids": ids,
			"names": names,
		})  // {"ids":{"a":"12","b":"23"},"names":{"first":"wang","last":"dictx"}}
	})

	type User struct {
		Name string `form:"name" json:"name" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}

	// 获取request-body，并且绑定到结构体
	// curl -X POST -H"Content-Type:application/x-www-form-urlencoded" -d'{"name":"wangqiang","password":"1234"}' "http://localhost:8080/post-payload"
	r.POST("post-payload", func(context *gin.Context) {
		var user User
		/** 更多数据绑定的来源
		context.ShouldBindHeader()
		context.ShouldBindUri()
		context.ShouldBindQuery()
		 */
		if err := context.ShouldBindJSON(&user); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		context.JSON(http.StatusOK, gin.H{
			"message": "hello",
			"name": user.Name,
		})
	})

	// 上传单个文件
	// curl -H"Content-Type:multipart/form-data" -F "Filename=@E:/data_transfer/index.html" "http://localhost:8080/upload"
	r.POST("/upload", func(context *gin.Context) {
		file,_ := context.FormFile("Filename")
		context.SaveUploadedFile(file, "E:/data_transfer/xx-upload.html")

		context.String(http.StatusOK, "%s uploaded", file.Filename)
	})

	// 上传多个文件
	r.POST("/upload-files", func(context *gin.Context) {
		form,_ := context.MultipartForm()
		files := form.File["Filename"]

		for _, file := range files {
			context.SaveUploadedFile(file, "E:/data_transfer/uploaded_" + file.Filename)
		}

		context.String(http.StatusOK, "%d files uploaded", len(files))
	})

	v2 := r.Group("/v2")
	{
		v2.GET("/login", func(context *gin.Context) {
			context.String(http.StatusOK, "v2 login")
		})

		v2.GET("/logout", func(context *gin.Context) {
			context.String(http.StatusOK, "v2 logout")
		})
	}

	v1 := r.Group("/v1")
	// 使用自定义中间件
	v1.Use(func(context *gin.Context) {
		fullpath := context.FullPath()
		context.ClientIP()
		fmt.Println(">>>> use custom middleware", fullpath)
	})
	{
		v1.GET("/login", func(context *gin.Context) {
			context.String(http.StatusOK, "v1 login")
		})

		v1.GET("/logout", func(context *gin.Context) {
			context.String(http.StatusOK, "v1 logout")
		})
	}

	r.Run()  // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func log(context *gin.Context) {
	start := time.Now()

	context.Next()
	time.Sleep(1234 * time.Millisecond)

	end := time.Now()
	cost := end.Sub(start)

	fullpath := context.FullPath()
	// 打印请求路径和耗时
	fmt.Println("req=", fullpath, ",cost=", cost.Milliseconds(), "ms")
}

func SampleMainSecond() {

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(log)

	router.GET("ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	// redirect跳转
	router.GET("redirect", func(context *gin.Context) {
		context.Redirect(http.StatusFound, "http://baidu.com")
	})

	router.Run("localhost:8080")
}