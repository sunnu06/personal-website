package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mywebsiteProject/lib/ini"
	"mywebsiteProject/controllers/project"
	"net/http"
)

//[Global var]
var url string


func Main()  {
	r := gin.Default()
	r.LoadHTMLGlob("views/*")
	r.Static("/public","./public")

	//[ini檔配置]
	url = ini.Main()

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
			fmt.Println(view_name)
			switch view_name {
			case "index.html":
				ctx.HTML(http.StatusOK, ctx.Param("view_name"), gin.H{
					"title" : "Jing Ru's Website",
					"nav" : "index",
				})
			default:
				ctx.HTML(http.StatusOK, ctx.Param("view_name"), gin.H{
					"title" : "Jing Ru's Website",
					"nav" : "not index",
				})
			}

		})
	}

	//http://localhost:8080/sideproject/portfolio_details.html?projectID=1&project=Manufacturing%20Execution%20System%20(MES)


	sideprojectGroup := r.Group("/sideproject")
	{
		sideprojectGroup.GET(":view_name", project.Show)
	}
	r.Run()
}

