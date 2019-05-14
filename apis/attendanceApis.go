package apis

import (
	"attendance/models"
	"attendance/util"
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/gin-gonic/gin"
	"strconv"
)

//更新签到状态
func UpdateClockState(c *gin.Context) {
	var clockIn models.ClockIn
	data, _ := c.GetRawData()

	//
	secId, err := jsonparser.GetInt(data, "sectionId")
	if err != nil || secId <= 0 {
		errJson(c, 400, "sectionId参数有误")
		return
	}
	clockIn.Section.SectionId = secId

	//
	state, err := jsonparser.GetInt(data, "state")
	if state < 0 || state > 3 {
		errJson(c, 400, "state参数有误")
		return
	}
	clockIn.State = state

	//
	stuId, err := jsonparser.GetInt(data, "student", "studentId")
	if err != nil {
		stuNum, err := jsonparser.GetString(data, "student", "studentNum")
		if err != nil || stuNum == "" {
			errJson(c, 400, "studentNum参数有误")
			return
		}
		clockIn.Student.StudentNum = stuNum
		err = clockIn.Student.SearchStuIdAndNameByStuNum()
		if err != nil {
			errJson(c, 400, "获取学生信息失败: " + err.Error())
			return
		}
	} else {
		if stuId <= 0 {
			errJson(c, 400, "studentId参数有误")
			return
		}
		clockIn.Student.StudentId = stuId
		err = clockIn.Student.SearchStuNumAndNameById()
		if err != nil {
			errJson(c, 500, "获取学生信息失败: " + err.Error())
			return
		}
	}

	err = clockIn.UpdateClockStateBySecIdAndStuId()
	if err != nil {
		errJson(c, 400, "无此学生的签到相关信息: " + err.Error())
		return
	}
	dataJson(c, "修改成功", gin.H{
		"student": clockIn.Student,
		"state": clockIn.State,
	})
	return
}

//查询
func QueryAttendanceNumAndRateOfCourse(c *gin.Context) {
	var clockIn models.ClockIn
	data, _ := c.GetRawData()

	//
	secId, err := jsonparser.GetInt(data, "sectionId")
	if err != nil || secId <= 0 {
		errJson(c, 400, "sectionId参数有误")
		return
	}
	clockIn.Section.SectionId = secId

	//
	state, err := jsonparser.GetInt(data, "state")
	if state < 0 || state > 3 {
		errJson(c, 400, "state参数有误")
		return
	}
	clockIn.State = state

	//
	stuId, err := jsonparser.GetInt(data, "student", "studentId")
	if err != nil {
		stuNum, err := jsonparser.GetString(data, "student", "studentNum")
		if err != nil || stuNum == "" {
			errJson(c, 400, "studentNum参数有误")
			return
		}
		clockIn.Student.StudentNum = stuNum
		err = clockIn.Student.SearchStuIdAndNameByStuNum()
		if err != nil {
			errJson(c, 400, "获取学生信息失败")
			return
		}
	} else {
		if stuId <= 0 {
			errJson(c, 400, "studentId参数有误")
			return
		}
		clockIn.Student.StudentId = stuId
		err = clockIn.Student.SearchStuNumAndNameById()
		if err != nil {
			errJson(c, 500, "获取学生信息失败")
			return
		}
	}

	err = clockIn.UpdateClockStateBySecIdAndStuId()
	if err != nil {
		errJson(c, 400, "无此学生的签到相关信息")
		return
	}
	dataJson(c, "修改成功", gin.H{
		"student": clockIn.Student,
		"state": clockIn.State,
	})
	return
}

func SearchAttendanceCountAndRateOfSection(c *gin.Context) {
	var clockIn models.ClockIn
	secId, err := strconv.ParseInt(c.Param("sectionId"), 10, 64)
	if err != nil {
		errJson(c, 400, "sectionId参数有误: " + err.Error())
		return
	}
	clockIn.Section.SectionId = secId

	//1-1 通过sectionId找到courseId，然后再找到classsId，找出班级人数
	err = clockIn.Section.SearchCourseIddById()
	if err != nil {
		errJson(c, 500, "数据库错误: " + err.Error())
		return
	}
	err = clockIn.Section.Course.SearchdClasssIdById()
	if err != nil {
		errJson(c, 500, "数据库错误: " + err.Error())
		return
	}
	err = clockIn.Section.Course.Classs.SearchStuCountById()
	if err != nil {
		errJson(c, 500, "数据库错误3: " + err.Error())
		return
	}

	////1-2 统计clock_in 表中section_id = secId 的有多少条记录，代表着应签到人数
	//err = clockIn.CountStuNumBySecId()
	//if err != nil {
	//	errJson(c, 500, "数据库错误: " + err.Error())
	//	return
	//}

	//2统计section_id = secId and state=1的人数，代表着已签到人数
	clockIn.State = 1
	attendanceCount := 0
	attendanceCount, err = clockIn.CountStuNumBySecIdAndState()
	if err != nil {
		errJson(c, 500, "数据库错误4: " + err.Error())
		return
	}
	fmt.Println("api: ", &attendanceCount)

	//3计算出勤率
	//attendanceRate := float64(attendanceCount) / float64(clockIn.Student.Classs.StudentCount)
	attendanceRate := float64(attendanceCount) / float64(clockIn.Section.Course.Classs.StudentCount)
	attendanceRate = util.ParseAndRemainFloat64(attendanceRate, 4)
	attendanceRate *= 100
	dataJson(c, "查询成功", gin.H{
		"attendanceCount": attendanceCount,
		"attendanceRate": attendanceRate,
	})
	return
}

func SearchAttendancesOfStudent(c *gin.Context) {
	var clockIn models.ClockIn
	clockIns := make([]models.ClockInByStu, 0)
	data, _ := c.GetRawData()
	page, _ := strconv.Atoi(c.Param("page"))
	limit, _ := strconv.Atoi(c.Param("limit"))

	offset, err := util.Page(page, limit)
	if err != nil {
		errJson(c, 500, err.Error())
		return
	}

	stuId, err := jsonparser.GetInt(data, "student", "studentId")
	if err != nil || stuId <= 0 {
		errJson(c, 400, "studentId参数有误")
		return
	}
	clockIn.Student.StudentId = stuId

	couId, err := jsonparser.GetInt(data, "course", "courseId")
	if err != nil || couId <= 0 {
		errJson(c, 400, "courseId参数有误")
		return
	}
	clockIn.Section.Course.CourseId = couId

	////查询签到状态
	//state, err1 := jsonparser.GetInt(data, "state")
	//if err1 == nil {
	//	if state < 0 || state > 3 {
	//		errJson(c, 400, "state参数有误")
	//		return
	//	}
	//	clockIn.State = state
	//}

	//
	sortField, err2 := jsonparser.GetString(data, "sortField")
	if err2 == nil {
		switch sortField {
		case "state":
			break
		case "startDate":
			break
		default:
			errJson(c, 400, "sortField参数有误")
			return
		}
	}

	direction, err3 := jsonparser.GetInt(data, "direction")
	if err3 == nil {
		switch direction {
		case 1:
			break
		case 2:
			break
		default:
			errJson(c, 400, "direction参数有误")
			return
		}
	}

	if err2 != nil {
		clockIns, err = clockIn.SearchSecDateAndClockStateByStuIdAndCourseId(offset, limit)
		if err != nil {
			errJson(c, 500, "数据库错误： " + err.Error())
			return
		}
	} else if err2 == nil && err3 != nil {
		switch sortField {
		case "state":
			clockIns, err = clockIn.SearchSecDateAndClockStateByStuIdAndCourseIdSortByClockStateDesc(offset, limit)
			if err != nil {
				errJson(c, 500, "数据库错误： " + err.Error())
				return
			}
		case "startDate":
			clockIns, err = clockIn.SearchSecDateAndClockStateByStuIdAndCourseIdSortBySecDateDesc(offset, limit)
			if err != nil {
				errJson(c, 500, "数据库错误： " + err.Error())
				return
			}
		}
	} else {
		if sortField == "state" {
			if direction == 2{
				clockIns, err = clockIn.SearchSecDateAndClockStateByStuIdAndCourseIdSortByClockStateDesc(offset, limit)
				if err != nil {
					errJson(c, 500, "数据库错误： " + err.Error())
					return
				}
			} else {
				clockIns, err = clockIn.SearchSecDateAndClockStateByStuIdAndCourseIdSortByClockStateAsc(offset, limit)
				if err != nil {
					errJson(c, 500, "数据库错误： " + err.Error())
					return
				}
			}
		} else {
			if direction == 2 {
				clockIns, err = clockIn.SearchSecDateAndClockStateByStuIdAndCourseIdSortBySecDateDesc(offset, limit)
				if err != nil {
					errJson(c, 500, "数据库错误： " + err.Error())
					return
				}
			} else {
				clockIns, err = clockIn.SearchSecDateAndClockStateByStuIdAndCourseIdSortBySecDateAsc(offset, limit)
				if err != nil {
					errJson(c, 500, "数据库错误： " + err.Error())
					return
				}
			}
		}
	}

	//fmt.Println("最后数据: ",clockIns)
	//var clo models.ClockInByStuList
	//clo.ClockInByStus = clockIns

	//var clo models.ClockInByStuRe
	//clos := make([]models.ClockInByStuRe, 0)
	//for _,v := range clockIns {
	//	clo.ClockInByStu = v
	//	clos = append(clos, clo)
	//}
	//dataJson(c, "查询成功", clos)
	totalElement, err := clockIn.CountSecDateAndClockStateByStuIdAndCourseId()
	if err != nil {
		errJson(c, 500, "数据库错误： " + err.Error())
		return
	}

	c.JSON(200, gin.H{
		"status": 0,
		"msg":    "查询成功",
		"data": gin.H{
			"content": clockIns,
			"pageable": gin.H{
				"totalElement": totalElement,
			},
		},
	})
	//dataJson(c, "查询成功", clockIns)
	return
}


