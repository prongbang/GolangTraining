package utils

func CalcPageNo(pageNo int, pageCount int) int {
	if pageNo == 0 {
		pageNo = 1
	}

	if pageNo < 0 {
		pageNo = pageCount
	} else if pageNo > pageCount {
		pageNo = pageCount
	}

	return pageNo
}

func CalcStartRow(pageNo int, pageSize int) int  {

	return (pageNo - 1) * pageSize + 1
}

func CalcEndRow(startRow int, pageSize int) int  {

	return startRow + pageSize - 1
}