package apis

import (
	"attendance/models"
	"github.com/buger/jsonparser"
	"github.com/gin-gonic/gin"
)

func LoginApi(c *gin.Context) {
	data, _ := c.GetRawData()
	userName, err := jsonparser.GetString(data, "username")
	if err != nil || userName == "" {
		errJson(c, 400, "用户名不能为空")
		return
	}
	passWord, err := jsonparser.GetString(data, "password")
	if err != nil || passWord == "" {
		errJson(c, 400, "密码不能为空")
		return
	}

	var admin models.Admin
	var student models.Student
	var teacher models.Teacher

	if userName == "admin" {
		admin.Username = userName
		err := admin.SearchAdmIdAndPsdByUsername()
		if err != nil {
			errJson(c, 400, "账号不存在： "+err.Error())
			return
		}
		if admin.PassWord != passWord {
			errJson(c, 400, "密码错误")
			return
		}
		dataJson(c, "登陆成功", gin.H{
			"authority": "admin",
			"userId":    admin.AdminId,
			"username":  admin.Username,
		})
		return
	} else if len(userName) == 10 {
		student.StudentNum = userName
		err := student.SearchStuIdAndNameAndPsdByStuNum()
		if err != nil {
			errJson(c, 400, "账号不存在： "+err.Error())
			return
		}
		if student.PassWord != passWord {
			errJson(c, 400, "密码错误")
			return
		}
		dataJson(c, "登陆成功", gin.H{
			"authority": "student",
			"userId":    student.StudentId,
			"username":  student.TrueName,
		})
		return
	} else {
		teacher.TeacherNum = userName
		err := teacher.SearchTeaIdAndTrueNameAndPsdByTeaNum()
		if err != nil {
			errJson(c, 400, "账号不存在： "+err.Error())
			return
		}
		if teacher.PassWord != passWord {
			errJson(c, 400, "密码错误")
			return
		}
		dataJson(c, "登陆成功", gin.H{
			"authority": "teacher",
			"userId":    teacher.TeacherId,
			"username":  teacher.TrueName,
		})
		return
	}
}

func UpdatePsdApi(c *gin.Context) {
	data, _ := c.GetRawData()
	userType, err := jsonparser.GetInt(data, "type")
	if err != nil || userType <= 0 || userType > 3 {
		errJson(c, 400, "用户类型参数传递错误")
		return
	}
	userId, err := jsonparser.GetInt(data, "userId")
	if err != nil || userId <= 0 {
		errJson(c, 400, "用户id参数错误")
		return
	}
	oldPsd, err := jsonparser.GetString(data, "oldPassword")
	if err != nil || oldPsd == "" {
		errJson(c, 400, "原密码为空")
		return
	}
	newPsd, err := jsonparser.GetString(data, "newPassword")
	if err != nil || newPsd == "" {
		errJson(c, 400, "新密码为空")
		return
	}

	var admin models.Admin
	var student models.Student
	var teacher models.Teacher

	if userType == 1 {
		admin.AdminId = int(userId)
		err := admin.SearchAdmPasswordById()
		if err != nil {
			errJson(c, 400, "用户不存在： "+err.Error())
			return
		}
		if admin.PassWord != oldPsd {
			errJson(c, 400, "原密码错误")
			return
		}
		admin.PassWord = newPsd
		err = admin.UpdateAdmPasswordById()
		if err != nil {
			errJson(c, 500, "数据库错误： "+err.Error())
			return
		}
		okJson(c, "密码修改成功")
		return
	} else if userType == 2 {
		teacher.TeacherId = int(userId)
		err := teacher.SearchTeaPsdById()
		if err != nil {
			errJson(c, 400, "用户不存在： "+err.Error())
			return
		}
		if teacher.PassWord != oldPsd {
			errJson(c, 400, "原密码错误")
			return
		}
		teacher.PassWord = newPsd
		err = teacher.UpdateTeaPsdById()
		if err != nil {
			errJson(c, 500, "数据库错误： "+err.Error())
			return
		}
		okJson(c, "密码修改成功")
		return
	} else if userType == 3 {
		student.StudentId = userId
		err := student.SearchStuPsdById()
		if err != nil {
			errJson(c, 400, "用户不存在： "+err.Error())
			return
		}
		if student.PassWord != oldPsd {
			errJson(c, 400, "原密码错误")
			return
		}
		student.PassWord = newPsd
		err = student.UpdateStuPsdById()
		if err != nil {
			errJson(c, 500, "数据库错误： "+err.Error())
			return
		}
		okJson(c, "密码修改成功")
		return
	}
}
