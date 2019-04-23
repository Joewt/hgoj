package controllers


func PageCal(totalNum int64,pageNo int32,pageSize int) (bool,int32,int32) {
	isPage := true
	if int(totalNum) < pageSize {
		isPage = false
	}
	pagePrev := pageNo
	pageNext := pageNo + 2
	temp := int(totalNum) / pageSize
	if int(pageNo) == temp {
		pageNext = pageNo + 1
	}
	if pageNo == 0 {
		pagePrev = pageNo + 1
	}
	return isPage,pagePrev,pageNext
}
