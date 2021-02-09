package config

import (
	"github.com/astaxie/beego/cache"
)

var Cache cache.Cache

func init() {
	var err error
	Cache, err = cache.NewCache("file", `{"CachePath":"tmp/cache","FileSuffix":".config","EmbedExpiry":"60", "DirectoryLevel":"3"}`)
	if err != nil {
		panic(err)
	}
}
