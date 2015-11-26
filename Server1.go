package main
import (
   	"fmt"
    "httprouter"
    "net/http"
    "encoding/json"

    "strconv"
)

type Pair struct{
	Key int `json:"key,omitempty"`
	Val string `json:"value,omitempty"`
}

type Arr struct{
	Maps []Pair `json:"-"`
}

var data = make(map[int]string)

func updatePair(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	key,_ := strconv.Atoi(p.ByName("id"))
	Val := p.ByName("value")
	data[key] = Val
	rw.WriteHeader(200)
	fmt.Fprint(rw, "200")
}

func getPair(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	key,_ := strconv.Atoi(p.ByName("id"))
	Pair := new(Pair)
	Pair.Key = key
	Pair.Val = data[key]
	oj, err := json.Marshal(Pair)
	if err != nil {
		//log.Println(error.Error())
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
  fmt.Fprintf(rw, "%s", oj)


}

func getAllPair(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	Arr := new(Arr)
	Arr.Maps = []Pair{}
	for k, v := range data {
        fmt.Println("k:", k, "v:", v)
		Pair := new(Pair)
		Pair.Key = k
		Pair.Val = v
		Arr.Maps = append(Arr.Maps, *Pair)
	}
	oj, err := json.Marshal(Arr)
	if err != nil {
		//log.Println(error.Error())
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)

  fmt.Fprintf(rw, "%s", oj)

}

func main() {
    mux := httprouter.New()
    mux.PUT("/keys/:id/:value", updatePair)
    mux.GET("/keys/:id", getPair)
    mux.GET("/keys", getAllPair)
    server := http.Server{
            Addr:        "0.0.0.0:3010",
            Handler: mux,
    }
    server.ListenAndServe()

  }
