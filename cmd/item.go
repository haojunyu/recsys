/*
Copyright © 2020 haojunyu <haojunyu2012@gmail.com>

图文item
*/

package cmd

type NewsList []NewsI

type NewsI struct {
	Itemid       int 
	Newsid       string `json:"newsid"`  
	Docid        int `json:"docid"`
	Title        string `json:"title"`
	Newstate     string `json:"newsState"`
	Edit_type    int `json:"edit_type"`
	Description  string `json:"description"`
	Source       int  `json:"source"`
	Body         string
	MediaId      int `json:"mediaId"`
	Type         int
	Bigimg       string `json:"bigImg"`
	Inner_source string `json:"inner_source"`
	From_url_md5 string `json:"from_url_md5"`
	Vulgar       int `json:"vulgar"`
	Only         int `json:"only"`
	Ctime        int `json:"ctime"`
	From         string `json:"from"`
	HttpsUrl     string `json:"httpsUrl"`
	Urllocal     string `json:"urlLocal"`
	Url3g        string `json:"url3g"`
	Editor       string `json:"editor"`
	Imgs         []string `json:"imgs"` 
	Publishtime  int `json:"publishTime"`
	Cate1        int `json:"cate1"`
	Cate2        int `json:"cate2"`
	Cate3        int  `json:"cate3"`
	Original     int	`json:"original"`
	Level        int `json:"level"`
	Author       string `json:"author"`
	Check_status int `json:"check_status"`
	Expire       int `json:"expire"`
	Region       string `json:"region"`
	Status       int `json:"status"`
	Mediatype    int `json:"mediaType"`
	Startkey     string

	// list
	Id int  `json:"id"` 
	FetchedUrl  string `json:"fetchhedUrl"`
	WifiFetchedUrl string `json:"wifiFetchedUrl"` 
	Img3 []string `json:"img3"`
	ImgSize int `json:"imgSize"`
	BodyLength  int  `json:"bodyLength"`
	Img1 string `json:"img1"`
	ts int `json:"ts"`
}




type NewsM struct {
	Itemid       int 
	Newsid       string `json:"newsid"`  
	Docid        int `json:"docid"`
	Title        string `json:"title"`
	Newstate     string `json:"newsState"`
	Edit_type    int `json:"edit_type"`
	Description  string `json:"description"`
	Source       int  `json:"source"`
	Body         string
	MediaId      int `json:"mediaId"`
	Type         int
	Bigimg       string `json:"bigImg"`
	Inner_source string `json:"inner_source"`
	From_url_md5 string `json:"from_url_md5"`
	Vulgar       int `json:"vulgar"`
	Only         int `json:"only"`
	Ctime        int `json:"ctime"`
	From         string `json:"from"`
	HttpsUrl     string `json:"httpsUrl"`
	Urllocal     string `json:"urlLocal"`
	Url3g        string `json:"url3g"`
	Editor       string `json:"editor"`
	Imgs 				string 
	Publishtime  int `json:"publishTime"`
	Cate1        int `json:"cate1"`
	Cate2        int `json:"cate2"`
	Cate3        int  `json:"cate3"`
	Original     int	`json:"original"`
	Level        int `json:"level"`
	Author       string `json:"author"`
	Check_status int `json:"check_status"`
	Expire       int `json:"expire"`
	Region       string `json:"region"`
	Status       int `json:"status"`
	Mediatype    int `json:"mediaType"`
	Startkey     string
	Dg_item_tag  string
	Dg_item_cate string
	Create_time  string
	Update_time  string
}