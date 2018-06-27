package sender

import (
	"fmt"
	"net/http"
	"net/url"

	"gopkg.in/mgo.v2/bson"
)

func SendSMS(mobile string, smsbody string, values bson.M) {

	params := url.Values{
		"user":   {values["username"].(string)},
		"pwd":    {values["password"].(string)},
		"sender": {values["sender"].(string)},
		"mt":     {"2"},
		"msg":    {smsbody},
		"mobile": {mobile},
	}
	furl := fmt.Sprintf("%s?%s", values["url"].(string),
		params.Encode())
	fmt.Println(furl)
	client := &http.Client{}

	req, _ := http.NewRequest("GET", furl, nil)

	resp, _ := client.Do(req)
	fmt.Println(resp)

}
