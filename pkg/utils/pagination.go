package utils

import (
    "github.com/Unknwon/com"
    "github.com/gin-gonic/gin"
    "iam/pkg/settings"
)

// 数据库分页用的页数
func GetPage(c *gin.Context, total int) int {
    newPageNum := 0
    page, _ := com.StrTo(c.Query("page_number")).Int()
    if page > 0 {
        pageSize := GetPageSize(c)
        newPageNum = (page - 1) * pageSize
    }
    return newPageNum
}

func GetPageSize(c *gin.Context) int {
    pageSize := settings.AppSetting.PageSize
    pSizeStr := c.Query("page_size")
    if pSizeStr != "" {
        pageSize, _ = com.StrTo(pSizeStr).Int()
    }
    return pageSize
}

// 页面显示的当前页数
func ShowPageNum(pageNum, pageSize, total int) int {
    if pageNum == 0 {
        pageNum = 1
    } else {
        pageNum = pageNum + 1
    }
    if pageSize * pageNum > total {
        pageNum = pageNum / pageSize + pageNum % pageSize
    }
    return pageNum
}