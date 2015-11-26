package main
import (
   	"fmt"
    "httprouter"
    "net/http"
    "encoding/json"
//    "log"
    "strconv"
)

type MapData struct{
	Key int
	Value string
}

type MapDataArray struct{
	Maps []MapData
}

var data = make(map[int]string)

func updateMap(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	key,_ := strconv.Atoi(p.ByName("key"))
	value := p.ByName("value")
	data[key] = value
	rw.WriteHeader(http.StatusCreated)
	fmt.Fprint(rw, "200")
}

func getMapData(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	key,_ := strconv.Atoi(p.ByName("key"))
	mapData := new(MapData)
	mapData.Key = key
	mapData.Value = data[key]
	outgoingJSON, err := json.Marshal(mapData)
	if err != nil {
		//log.Println(error.Error())
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	fmt.Fprint(rw, string(outgoingJSON))
}

func getAllMapData(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	mapDataArray := new(MapDataArray)
	mapDataArray.Maps = []MapData{}
	for k, v := range data {
		mapData := new(MapData)
		mapData.Key = k
		mapData.Value = v
		mapDataArray.Maps = append(mapDataArray.Maps, *mapData)
	}
	outgoingJSON, err := json.Marshal(mapDataArray)
	if err != nil {
		//log.Println(error.Error())
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	fmt.Fprint(rw, string(outgoingJSON))
}

func main() {
    mux := httprouter.New()
    mux.PUT("/keys/:key/:value", updateMap)
    mux.GET("/keys/:key", getMapData)
    mux.GET("/keys", getAllMapData)
    server := http.Server{
            Addr:        "0.0.0.0:3011",
            Handler: mux,
    }
    server.ListenAndServe()
}
