package router

import (
	"attendance/apis"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func InitRouter(port string) {
	//port := "8080"

	//修改模式
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
	gin.DisableConsoleColor()
	// 记录到文件。
	str, _ := os.Getwd()
	fName := str + "/gin.log"
	f, _ := os.Create(fName)
	gin.DefaultWriter = io.MultiWriter(f)

	//设定请求url不存在的返回值
	r.NoRoute(apis.NoResponse)
	r.Use(cors.Default())

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "a")
	})

	user := r.Group("user")
	{
		user.POST("/login", apis.LoginApi)
		user.POST("/updatePassword", apis.UpdatePsdApi)
	}

	academy := r.Group("academy")
	{
		academy.POST("/", apis.AddOrUpdateAcademyApi)
		academy.GET("/:academyId", apis.FindOneAcademyApi)
		academy.DELETE("/:academyId", apis.DeleteAcademyApi)
		//分页查询
		academy.POST("/search/:page/:limit", apis.SearchAcademyApi)
		//academy.POST("/test", apis.TestApi)
	}

	//teacher := r.Group("teacher")
	//{
	//	//teacher.POST("/", apis.UpdateTeacherApi)
	//	teacher.GET("/", )
	//	teacher.DELETE("/:teacherId", )
	//	//分页查询
	//	teacher.POST("/search/:page/:limit", )
	//}
	//
	//student := r.Group("student")
	//{
	//	//增加或修改
	//	//student.POST("/", apis.UpdateStudentApi)
	//	student.GET("/", )
	//	student.DELETE("/:studentId", )
	//	//分页查询
	//	student.POST("/search/:page/:limit", )
	//}

	attendance := r.Group("attendance")
	{
		attendance.POST("/updateAttendance", apis.UpdateClockState)
		attendance.GET("/queryAttendanceCountAndRateOfSection/:sectionId", apis.SearchAttendanceCountAndRateOfSection)
		attendance.POST("/queryAttendancesOfStudent/:page/:limit", apis.SearchAttendancesOfStudent)
	}

	r.Run(port)
}
