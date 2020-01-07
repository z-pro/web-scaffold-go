package utils

import (
	"fmt"
	"github.com/astaxie/beego"
	"math"
	"strings"
)

// LimitPage 分页组件
func LimitPage(currentPageNum, AllCount int, filterArgs, url string) (string, int, int) {

	if filterArgs!=""{
		filterArgs = fmt.Sprintf("?%v",filterArgs)
	}
	// currentPageNum 当前页
	// AllCount 总数据量
	// FilterArgs额外的url参数
	// url 分页按钮的url
	pageCount, _ := beego.GetConfig("Int", "pagesize", 10) //从配置文件获取每个分页的最大数据量
	showPage := 10  //分页按钮范围
	// 如果没有当前分页就会默认分页为第一页（传入数值需要处理在没有分页的时候会等于0）
	if AllCount < 1 {
		AllCount = 1
	}
	// 计算有多少分页整数
	allPage := AllCount / pageCount.(int)
	// 计算分页余数
	mod := math.Mod(float64(AllCount), float64(pageCount.(int)))
	if mod > 0 {
		//如果有余再加一页
		allPage++
	}
	//用来存放分页按钮的列表
	htmlList := []string{}
	pageHalf := (showPage - 1) / 2
	start := 0
	stop := 0
	var previous string
	var next string
	if allPage < showPage {
		// 如果总分页数小于展示的分页数那么分页按钮的结束为总分页数
		start = 1
		stop = allPage
	} else {
		// 如果当前页小于分页按钮的一半分页按钮的开始等于分页1结束等于总分页按钮
		if currentPageNum < pageHalf+1 {
			start = 1
			stop = showPage
		} else {
			if currentPageNum >= allPage-pageHalf {
				start = allPage - showPage
				stop = allPage
			} else {
				start = currentPageNum - pageHalf
				stop = currentPageNum + pageHalf
			}
		}
	}
	if currentPageNum <= 1 {
		// 如果当前页小于等于1那么上一页样式设置为不可用
		previous = "<li class='page-item disabled'><a href='#' class='page-link'>上一页</a></li>"
	} else {
		//as := "<li class='page-item'><a href='%v?pageNum=%v%v' class='page-link'  style='cursor:pointer;text-decoration:none;'>上一页<span aria-hidden='true'>&laquo;</span></a></li>"
		as := "<li class='page-item'><a href='%v/%v%v' class='page-link'  style='cursor:pointer;text-decoration:none;'>上一页<span aria-hidden='true'>&laquo;</span></a></li>"
		previous = fmt.Sprintf(as, url, currentPageNum-1, filterArgs)
	}
	htmlList = append(htmlList, previous)
	for i := start; i <= stop; i++ {
		temp := ""
		if currentPageNum == i {
			//temp = "<li class='page-item active'><a href='%v?page=%v%v' class='page-link' >%v</a></li>"
			temp = "<li class='page-item active'><a href='%v/%v%v' class='page-link' >%v</a></li>"
			temp = fmt.Sprintf(temp, url, i, filterArgs, i)
		} else {
			//temp = "<li class='page-item'><a href='%v?pageNum=%v%v' class='page-link' >%v</a></li>"
			temp = "<li class='page-item'><a href='%v/%v%v' class='page-link' >%v</a></li>"
			temp = fmt.Sprintf(temp, url, i, filterArgs, i)
		}
		htmlList = append(htmlList, temp)
	}
	if currentPageNum >= allPage {
		// 如果当前页大于等于总页数那么下一页样式设置为不可用
		next = "<li class='page-item disabled'><a href='#' class='page-link'>下一页</a></li>"
	} else {
		//as := "<li class='page-item'><a href='%v?pageNum=%v%v' class='page-link' >下一页</a></li>"
		as := "<li class='page-item'><a href='%v/%v%v' class='page-link' >下一页</a></li>"
		next = fmt.Sprintf(as, url, currentPageNum+1, filterArgs)
	}
	htmlList = append(htmlList, next)

	navHeader:=`<nav aria-label="..." style="float: right">
                                <ul class="pagination">`
	navFooter:=` </ul>
                   </nav>`
	//data := strings.Join(htmlList, "")
	data := fmt.Sprintf("%v%v%v",navHeader,strings.Join(htmlList, ""),navFooter)

	dataStart := 0
	dataStop := 0
	// dataStart 对应的是数据库的limit, dataStop对应的是数据库的offset
	dataStart = (currentPageNum - 1) * pageCount.(int)
	dataStop = pageCount.(int)
	return data, dataStart, dataStop
}