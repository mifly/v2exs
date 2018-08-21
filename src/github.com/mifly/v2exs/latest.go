package v2exs

import (
	//"net/http"
	"fmt"
	//"io/ioutil"
	"encoding/json"
	"net/http"
	"io/ioutil"
)

func StartGetLatest() error {
	/*ticker := time.NewTicker(2 * time.Second)
	for {
		select {
		case <-ticker.C:
			getLatest()
		}
	}*/
	getLatest()
	//gl()
	return nil
}

type Node struct {
	AvatarLarge      string `json:"avatar_large"`
	Name             string `json:"name"`
	AvatarNormal     string `json:"avatar_normal"`
	Title            string `json:"title"`
	Url              string `json:"url"`
	Topics           int `json:"topics"`
	Footer           string `json:"footer"`
	Header           string `json:"header"`
	TitleAlternative string `json:"title_alternative"`
	AvatarMini       string `json:"avatar_mini"`
	Starts           string `json:"starts"`
	Root             bool `json:"root"`
	Id               int    `json:"id"`
	ParentNodeName   string `json:"parent_node_name"`
}
type Member struct {
	Username     string `json:"username"`
	Website      string `json:"website"`
	Github       string `json:"github"`
	Psn          string `json:"psn"`
	AvatarNormal string `json:"avatar_normal"`
	Bio          string `json:"bio"`
	Url          string `json:"url"`
	TagLine      string `json:"tagline"`
	Twitter      string `json:"twitter"`
	Created      int64  `json:"created"`
	AvatarLarge  string `json:"avatar_large"`
	AvatarMini   string `json:"avatar_mini"`
	Location     string `json:"location"`
	Btc          string `json:"btc"`
	Id           int    `json:"id"`
}

type Latest struct {
	Node            Node   `json:"node,omitempty"`
	Member          Member `json:"member,omitempty"`
	LastReplyBy     string `json:"last_reply_by,omitempty"`
	LastTouched     int64 `json:"last_touched,omitempty"`
	Title           string `json:"title,omitempty"`
	Url             string `json:"url,omitempty"`
	Created         int64  `json:"created,omitempty"`
	Content         string `json:"content,omitempty"`
	ContentRendered string `json:"content_rendered,omitempty"`
	LastModified    int64  `json:"last_modified,omitempty"`
	Replies         int    `json:"replies,omitempty"`
	Id              int    `json:"id,omitempty"`
}


func getLatest() {
	var err error
	fmt.Println("get latest begin")
	resp, err := http.Get("https://www.v2ex.com/api/topics/latest.json")
	if err != nil {
		fmt.Printf("get latest failed %s\n", err.Error())
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read latest failed %s\n", err.Error())
		return
	}
	fmt.Printf("get latest:%s\n", string(data))
	//var latests LatestResult
	//var lat map[interface{}]interface{}
	var ll []Latest
	err = json.Unmarshal(data, &ll)
	if err != nil {
		fmt.Printf("Unmarshal latest failed %s\n", err.Error())
		return
	}
	//fmt.Printf("%s\n", latests.LatestList[0].Content)
	fmt.Printf("get latest done, %s\n", ll[0].Title)
}

