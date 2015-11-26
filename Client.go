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
	Key int
	Value string
}

var Cache = make(map[string]string)
var Data = make(map[int]string)
var Instances = []string{"3000","3001","3002"}
var S_keys []string

func g_p(text string)(string){
	hash := md5.Sum([]byte(text))
   	return hex.EncodeToString(hash[:])
}

func get_serv(k int)(string){
	    foundData := 0
		foundCache := 0
		index := 0
 		for index < len(S_keys){
 			if(foundData != 1){
 				if(g_p(strconv.Itoa(k)) == S_keys[index]){
					foundData = 1
				}
 			}else if(foundData == 1){
 				if(s_sl(Cache[S_keys[index]],Instances)){
 					foundCache = 1
 					break
 				}
 			}
 			if(index == len(S_keys)-1 && foundCache == 0){
 				index = 0
 			}else{
 				index += 1
 			}
		}
 		return Cache[S_keys[index]]
}

func put_h(rw http.ResponseWriter, req *http.Request, p httprouter.Params){
	k,_ := strconv.Atoi(p.ByName("key"))
	value := p.ByName("value")
 	url := "http://localhost:"+get_serv(k)+"/keys/"+strconv.Itoa(k)+"/"+value
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

func h_get(rw http.ResponseWriter, req *http.Request, p httprouter.Params){
	k,_ := strconv.Atoi(p.ByName("key"))
	resp, err := http.Get("http://localhost:"+get_serv(k)+"/keys/"+strconv.Itoa(k))
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
			//log.Println(error.Error())
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusCreated)
        fmt.Fprint(rw, string(oj))
    } else {
        fmt.Println(err)
    }
}

func s_sl(str string, list []string) bool {
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
    	Cache[g_p(each)] = each
    }

	for k, _ := range Data {
		Cache[g_p(strconv.Itoa(k))] = strconv.Itoa(k)
	}

	for k, _ :=range Cache{
		S_keys = append(S_keys,k)
	}

	sort.Strings(S_keys)
	mux := httprouter.New()
    mux.PUT("/keys/:key/:value", put_h)
    mux.GET("/keys/:key", h_get)
    server := http.Server{
            Addr:        "0.0.0.0:8187",
            Handler: mux,
    }
    server.ListenAndServe()
}
