package main
import (
   	"fmt"
    "httprouter"
    "net/http"
    "encoding/json"
//    "log"
    "strconv"
)

type Pair struct{
	Key int
	Value string
}

type PairArr struct{
	Maps []Pair
}

var data = make(map[int]string)

func upd_K(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	key,_ := strconv.Atoi(p.ByName("key"))
	value := p.ByName("value")
	data[key] = value
	rw.WriteHeader(http.StatusCreated)
	fmt.Fprint(rw, "200")
}

func get_K(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	key,_ := strconv.Atoi(p.ByName("key"))
	pair := new(Pair)
	pair.Key = key
	pair.Value = data[key]
	oj, err := json.Marshal(pair)
	if err != nil {
		//log.Println(error.Error())
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	fmt.Fprint(rw, string(oj))
}

func getall_K(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	pairarr := new(PairArr)
	pairarr.Maps = []Pair{}
	for k, v := range data {
        fmt.Println("k:", k, "v:", v)
		pair := new(Pair)
		pair.Key = k
		pair.Value = v
		pairarr.Maps = append(pairarr.Maps, *pair)
	}
	oj, err := json.Marshal(pairarr)
	if err != nil {
		//log.Println(error.Error())
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	fmt.Fprint(rw, string(oj))
}

func main() {
    mux := httprouter.New()
    mux.PUT("/keys/:key/:value", upd_K)
    mux.GET("/keys/:key", get_K)
    mux.GET("/keys", getall_K)
    server := http.Server{
            Addr:        "0.0.0.0:3002",
            Handler: mux,
    }
    server.ListenAndServe()
}
