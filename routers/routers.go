package routers

import (
	"github.com/gin-gonic/gin"
	"mywebsiteProject/lib/ini"
	"mywebsiteProject/controllers/project"
	"net/http"
)

//[Global var]
var url string
var webport int

func Main()  {
	r := gin.Default()
	r.LoadHTMLGlob("views/*")
	r.Static("/public","./public")

	//[ini檔配置]
	url, webport = ini.Main()

	//[Router 配置]
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, url)
	})

	viewsGroup := r.Group("/views")
	{
		viewsGroup.GET("/:view_name", func(ctx *gin.Context) {
			defer func(){
				if err := recover(); err != nil{
					ctx.Redirect(http.StatusTemporaryRedirect, "httpstatus/notfound")
				}
			}()
			view_name := ctx.Param("view_name")
			ctx.HTML(http.StatusOK, view_name, gin.H{
				"title" : "Jing Ru's CV",
			})
		})
	}

	//http://localhost:8080/sideproject/portfolio_details.html?projectID=1&project=Manufacturing%20Execution%20System%20(MES)


	sideprojectGroup := r.Group("/sideproject")
	{
		sideprojectGroup.GET(":view_name", project.Show)
	}
	r.Run()
}

