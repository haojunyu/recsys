/*
Copyright © 2020 haojunyu <haojunyu2012@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
)

// 全局变量
const (
	Rcd    = "data/sync_news_mysql.rcd"
	Url    = "http://news-open.51y5.net"
	SqlUrl = "thirdparty:Thirdparty123@tcp(rm-2zew4iw684p3rw391o.mysql.rds.aliyuncs.com:3306)/thirdparty?charset=utf8"
)

var Db *sqlx.DB

type Item struct {
	Itemid       int
	Newsid       string
	Docid        string
	Title        string
	Newstate     string
	Edit_type    int
	Description  string
	Source       string
	Body         string
	Mediaid      string
	Type         int
	Bigimg       string
	Inner_source string
	From_url_md5 string
	Vulgar       int
	Only         int
	Ctime        int
	From         string
	Httpsurl     string
	Urllocal     string
	Url3g        string
	Editor       string
	Imgs         string
	Publishtime  int
	Cate1        string
	Cate2        string
	Cate3        string
	Original     int
	Level        string
	Author       string
	Check_status int
	Expire       int
	Region       string
	Status       int
	Mediatype    int
	Startkey     string
	Dg_item_tag  string
	Dg_item_cate string
	Create_time  string
	Update_time  string
}

// cmsCmd represents the cms command
var cmsCmd = &cobra.Command{
	Use:   "cms",
	Short: "推荐系统的内容管理系统",
	Long:  `对接推荐系统的内容：用户信息user，物品信息item和行为日志action`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cms called")
		syncNewsMysql()
	},
}

func init() {
	rootCmd.AddCommand(cmsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cmsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cmsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// loadStartKey
func loadStartKey() int {
	var count int = 0
	var startKey int = 0
	fin, err := os.OpenFile(Rcd, os.O_RDONLY, 0)
	if err != nil {
		fmt.Println("Open file ", Rcd, " error!", err)
		panic(err)
	}
	defer fin.Close()

	sc := bufio.NewScanner(fin)
	/*default split the file use '\n'*/
	for sc.Scan() {
		count++
		fields := strings.Split(sc.Text(), "\t")
		if len(fields) != 4 {
			fmt.Printf("line format error: %s\n", sc.Text())
		}
		fmt.Printf("%v %T\n", fields, fields)

		intParse, err := strconv.Atoi(fields[1])
		if err != nil {
			fmt.Printf("startKey %s convert error\n", fields[1])
		}
		startKey = intParse

		fmt.Printf("the line %d: %s\n", count, sc.Text())
	}
	if err := sc.Err(); err != nil {
		fmt.Println("An error has happened")
		panic(err)
	}
	fmt.Println("sync news mysql")
	return startKey
}

// getItemByApi 通过api获取items数据
func getItemByApi() {

}

// connect 连接数据库
func connect() *sqlx.DB {
	db, err := sqlx.Open("mysql", SqlUrl)

	if err != nil {
		fmt.Println("open mysql failed,", err)
		return nil
	}
	return db
}

func writeMysql() {
	if Db == nil {
		Db = connect()
	}
	var items []Item
	err := Db.Select(&items, "select * from item_wifi_news_info_rt limit 10")
	if err != nil {
		fmt.Println("Select error", err)
	}

	/*
		//Begin函数内部会去获取连接
			tx,_ := db.Begin()
			for i := 1301;i<=1400;i++{
					//每次循环用的都是tx内部的连接，没有新建连接，效率高
					tx.Exec("INSERT INTO user(uid,username,age) values(?,?,?)",i,"user"+strconv.Itoa(i),i-1000)
			}
			//最后释放tx内部的连接
			tx.Commit()
	*/
	for i, v := range items {
		fmt.Printf("%d %s %s %s\n", i, v.Docid, v.Title, v.Create_time)
	}
	Db.Close()
}

// syncNewsMysql 图文同步，cms->mysql
func syncNewsMysql() {
	fmt.Println("sync news mysql")
	// 加载start_key
	startKey := loadStartKey()
	fmt.Println(startKey)

	// 获取数据
	//getItemByApi()

	// 写入mysql数据库
	writeMysql()

}
