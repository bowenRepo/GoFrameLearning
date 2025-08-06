// testdb.go
package main

import (
	"context"
	"fmt"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2" // MySQL 驱动
	"github.com/gogf/gf/v2/frame/g"
)

func main() {
	db := g.DB()

	// 1) 先 Ping 验证连接
	if err := db.PingMaster(); err != nil {
		fmt.Println("❌ 数据库连接失败：", err)
		return
	}

	// 2) 直接执行原生 SQL：SELECT 1
	v, err := db.GetValue(context.Background(), "SELECT 1")
	if err != nil {
		fmt.Println("❌ 执行查询失败：", err)
		return
	}

	// 3) 打印结果
	fmt.Println("✅ 数据库连接成功，查询结果：", v.Int())
}
