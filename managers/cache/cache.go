package cache

/*import "github.com/astaxie/beego"

var (
	urllist *beego.BeeCache
)

func init() {
	urllist = beego.NewBeeCache()
	urllist.Every = 0 //不过期
	urllist.Start()
}

func (this *ShortController) Post() {
	var result ShortResult
	longurl := this.Input().Get("longurl")
	beego.Info(longurl)
	result.UrlLong = longurl
	urlmd5 := models.GetMD5(longurl)
	beego.Info(urlmd5)
	if urllist.IsExist(urlmd5) {
		result.UrlShort = urllist.Get(urlmd5).(string)
	} else {
		result.UrlShort = models.Generate()
		err := urllist.Put(urlmd5, result.UrlShort, 0)
		if err != nil {
			beego.Info(err)
		}
		err = urllist.Put(result.UrlShort, longurl, 0)
		if err != nil {
			beego.Info(err)
		}
	}
	this.Data["json"] = result
	this.ServeJson()
}*/
