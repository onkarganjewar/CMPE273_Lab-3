package main

import (
	"fmt"
	"crypto/md5"
	"strconv"
	"encoding/hex"
	"sort"
	"net/http"
	"io/ioutil"
	 "encoding/json"
	 "httprouter"
)

type Pair struct{
	Key int `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

var cache = make(map[string]string)
var Data = make(map[int]string)
var Instances = []string{"3010","3011","3012"}
var sck []string

func get_c(text string)(string){
	hash := md5.Sum([]byte(text))
   	return hex.EncodeToString(hash[:])
}

func get_S(k int)(string){
	    ad := 0
		ac := 0
		index := 0
 		for index < len(sck){
 			if(ad != 1){
 				if(get_c(strconv.Itoa(k)) == sck[index]){
					ad = 1
				}
 			}else if(ad == 1){
 				if(strg(cache[sck[index]],Instances)){
 					ac = 1
 					break
 				}
 			}
 			if(index == len(sck)-1 && ac == 0){
 				index = 0
 			}else{
 				index += 1
 			}
		}
 		return cache[sck[index]]
}

func h_put(rw http.ResponseWriter, req *http.Request, p httprouter.Params){
	k,_ := strconv.Atoi(p.ByName("key"))
	value := p.ByName("value")
 	url := "http://localhost:"+get_S(k)+"/keys/"+strconv.Itoa(k)+"/"+value
	req, err := http.NewRequest("PUT", url, nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var data interface{}
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &data)
    var m = data.(interface{}).(float64)
 	fmt.Fprint(rw, m)
}

func get_H(rw http.ResponseWriter, req *http.Request, p httprouter.Params){
	k,_ := strconv.Atoi(p.ByName("key"))
	resp, err := http.Get("http://localhost:"+get_S(k)+"/keys/"+strconv.Itoa(k))
	if(err == nil) {
        var data interface{}
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &data)
	    var m = data.(map[string] interface{})
       	pair := new(Pair)
		pair.Key = int(m["Key"].(float64))
		pair.Value = m["Value"].(string)
		oj, err := json.Marshal(pair)
		if err != nil {

			return
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(200)
        fmt.Fprintf(rw, "%s", oj)

    } else {
        fmt.Println(err)
    }
}

func strg(str string, list []string) bool {
 	for _, v := range list {
 		if v == str {
 			return true
 		}
 	}
 	return false
 }

func main() {
	Data[1] = "a"
	Data[2] = "b"
	Data[3] = "c"
	Data[4] = "d"
	Data[5] = "e"
	Data[6] = "f"
	Data[7] = "g"
	Data[8] = "h"
	Data[9] = "i"
	Data[10] = "j"

	for _, each := range Instances {
    	cache[get_c(each)] = each
    }

	for k, _ := range Data {
		cache[get_c(strconv.Itoa(k))] = strconv.Itoa(k)
	}

	for k, _ :=range cache{
		sck = append(sck,k)
	}

	sort.Strings(sck)
	mux := httprouter.New()
    mux.PUT("/keys/:key/:value", h_put)
    mux.GET("/keys/:key", get_H)
    server := http.Server{
            Addr:        "0.0.0.0:8088",
            Handler: mux,
    }
    server.ListenAndServe()
}
