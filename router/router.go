package router

import  (
    "net/http"
    "github.com/gin-gonic/gin"
    "min-services/controller"
    "min-services/middleware/jwt"
)

func InitRouter(router *gin.Engine)  {
    router.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    
    // form-data
    router.POST("/form_post", func(c *gin.Context) {
        message := c.PostForm("message")
        nick := c.DefaultPostForm("nick", "anonymous")

        c.JSON(200, gin.H{
            "status":  "posted",
            "message": message,
            "nick":    nick,
        })
    })

    // This handler will match /user/john but will not match /user/ or /user
    router.GET("/user/:name", func(c *gin.Context) {
        name := c.Param("name")
        c.String(http.StatusOK, "Hello %s", name)
    })

    router.POST("/loginJSON", controller.LoginJson)
    router.POST("/loginXML", controller.LoginXML)
    router.POST("/loginForm", controller.LoginForm)


    apiv1 := router.Group("/api/v1")
    apiv1.Use(jwt.JWT())
    {
        //获取标签列表
        apiv1.GET("/tags", controller.GetTags)
        //新建标签
        apiv1.POST("/tags", controller.AddTag)
        //更新指定标签
        apiv1.PUT("/tags/:id", controller.EditTag)
        //删除指定标签
        apiv1.DELETE("/tags/:id", controller.DeleteTag)
    }
}
