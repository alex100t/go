package main

import (
	"fmt"
)

/*
	Some test code.

	File index.go contains function Handler as an entry point for Yandex Cloud Function

	Move to work dir:

		cd src/github.com/alex100t/doJD-bot

	Build and run tests:

		go build main.go do_JD.go tgbottype.go index.go
		go run main.go tgbottype.go index.go do_JD.go
		go test -v


	Bot name:

		t.me/do_jd_bot



*/

func main() {
	reqstr := `{"httpMethod":"POST","headers":{"Accept-Encoding":"gzip, deflate","Content-Length":"367","Content-Type":"application/json","X-Internal-Runtime":"golang116","X-Internal-Subject-Id":"","X-Internal-Subject-Type":"anonymous","X-Real-Remote-Address":"[91.108.6.75]:40082","X-Request-Id":"f94759e1-7a12-47b8-9636-b3033c98521d","X-Trace-Id":"0bafc97a-3200-4074-9c77-e4267ac62328"},"url":"","params":{},"multiValueParams":{},"pathParams":{},"multiValueHeaders":{"Accept-Encoding":["gzip, deflate"],"Content-Length":["367"],"Content-Type":["application/json"],"X-Internal-Runtime":["golang116"],"X-Internal-Subject-Id":[""],"X-Internal-Subject-Type":["anonymous"],"X-Real-Remote-Address":["[91.108.6.75]:40082"],"X-Request-Id":["f94759e1-7a12-47b8-9636-b3033c98521d"],"X-Trace-Id":["0bafc97a-3200-4074-9c77-e4267ac62328"]},"queryStringParameters":{},"multiValueQueryStringParameters":{},"requestContext":{"identity":{"sourceIp":"91.108.6.75","userAgent":""},"httpMethod":"POST","requestId":"f94759e1-7a12-47b8-9636-b3033c98521d","requestTime":"1/Sep/2021:07:29:06 +0000","requestTimeEpoch":1630481346},"body":"{\"update_id\":470427193,\n\"edited_message\":{\"message_id\":19,\"from\":{\"id\":195009075,\"is_bot\":false,\"first_name\":\"Alex\",\"last_name\":\"Butylkin\",\"language_code\":\"en\"},\"chat\":{\"id\":195009075,\"first_name\":\"Alex\",\"last_name\":\"Butylkin\",\"type\":\"private\"},\"date\":1630436758,\"edit_date\":1630436775,\"text\":\"/start zz 22\",\"entities\":[{\"offset\":0,\"length\":6,\"type\":\"bot_command\"}]}}","isBase64Encoded":false}`

	respstr, err := Handler([]byte(reqstr))
	if err != nil {
		fmt.Println(err)
		//return nil, err
	}

	fmt.Println(string(respstr))
	fmt.Println(" I \"do ^JD\" to convert dates between Julian and ISO date formats.\n Lets start!")

}
