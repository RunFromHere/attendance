package util

import "errors"

//用于计算解析page 和 limit
//page		第几页
//limit   	每页几行
//返回起始行

func Page(page, limit int) (int, error){
	if page <= 0 || limit <= 0{
		return 0, errors.New("分页参数错误")
	}
	offset := (page-1) * limit

	return offset, nil
}