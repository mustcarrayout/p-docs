package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	i := 0
	for ; i <= 10; i++ {
		go func() {

			rep, err := http.Get("http://139.196.84.185:9010/api/tool/generateSignTag?reportNo=%E6%B2%AA%E8%81%94%E5%9F%8E%EF%BC%882022%EF%BC%89%EF%BC%88%E4%BC%B0%EF%BC%89%E5%AD%97%E7%AC%ACM12613%E5%8F%B7")

			if err != nil {
				log.Panicln(err)
			}
			defer rep.Body.Close()

			datas, err := ioutil.ReadAll(rep.Body)
			if err != nil {
				log.Panicln(err)
			}

			fmt.Println(string(datas))

		}()
	}

	select {}

}
