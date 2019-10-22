package main

import (
	"./api/DomainRequest"
	"./api/DomainResponse"
	"./api/URLDB"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

const (
	Exists    = "EXISTS"
	NotExists = "NOT_EXISTS"
	CACHED    = "CACHED"
	NotCached = "NotCached"
)

var domains []string                // mocked storage for real time data
var cache = make(map[string]string) // md5 key based cache for returning previous domain fetching results

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(domains)
}

func ValidateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		_ = r.ParseForm()
		body, err := ioutil.ReadAll(r.Body) //get payload body as a byte array
		if err != nil {
			log.Println(err)
		}
		//easyjson object deserialization
		reqObj := &DomainRequest.DomainRequest{}
		e := reqObj.UnmarshalJSON(body)
		if e != nil {
			log.Println("error parsing request body")
		}
		//check for oayload correctness
		if validErrs := ValidateParams(reqObj); len(validErrs) > 0 {
			err := map[string]interface{}{"validationError": validErrs}
			w.Header().Set("Content-type", "applciation/json")
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(err)
		} else {
			var response = HandleDomainRequest(reqObj, body)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			str, err := response.MarshalJSON()
			if err == nil {
				_, _ = w.Write(str)
			} else {
				_ = json.NewEncoder(w).Encode(err)
			}
		}
	} else {
		http.Error(w, "Invalid request method for this end point", http.StatusMethodNotAllowed)
	}
}

func HandleDomainRequest(reqObj *DomainRequest.DomainRequest, body []byte) *DomainResponse.DomainResponse {
	u := reqObj.Domain + reqObj.Path
	response := &DomainResponse.DomainResponse{}
	// check if we have a cached result
	var cacheState, val = GetResultFromCache(u)
	if cacheState == CACHED {
		GetFormattedRes(val, u, response)
	} else { // get real time data from db
		GetDBFormatedResponse(u, response, body)
	}
	return response
}

func GetDBFormatedResponse(url string, response *DomainResponse.DomainResponse, body []byte) {
	var key = url // + string(body)
	if !(IsURLExistsInDB(url)) {
		_ = response.UnmarshalJSON([]byte(fmt.Sprintf(`{"location":%s }`, "")))
		CachePut(GetMD5Hash(key), NotExists)
	} else {
		msg := fmt.Sprintf(`{"location":%s }`, strconv.Quote(url))
		_ = response.UnmarshalJSON([]byte(msg))

		CachePut(GetMD5Hash(key), Exists)
	}

}

func GetFormattedRes(val string, url string, response *DomainResponse.DomainResponse) {
	if val == Exists {
		msg := fmt.Sprintf(`{"location":%s }`, strconv.Quote(url))
		_ = response.UnmarshalJSON([]byte(msg))

	} else if val == NotExists {
		_ = response.UnmarshalJSON([]byte(fmt.Sprintf(`{"location":%s }`, "")))
	}
}

func ValidateParams(data *DomainRequest.DomainRequest) url.Values {

	errs := url.Values{}
	// check if the domain and path are  empty or missing
	if data.Domain == "" {
		errs.Add("error", "domain param is missing")
	}
	if data.Path == "" {
		errs.Add("error", "path param is missing")
	}
	return errs
}

// Open our jsonFile mock and store in memory
func loadMock() {
	db := &URLDB.URLDB{}
	byteValue, _ := ioutil.ReadFile("./configs/db.json")
	_ = db.UnmarshalJSON(byteValue)
	domains = db.Domains
}

// used to transform the cache url key to md5 encrypted key, to reduce size
func GetMD5Hash(url string) string {
	hash := md5.New()
	hash.Write([]byte(url))
	return hex.EncodeToString(hash.Sum(nil))
}

func GetResultFromCache(url string) (string, string) {
	res := cache[GetMD5Hash(url)]
	if res != "" {
		log.Println("cache hit for:" + url)
		return CACHED, res
	} else {
		log.Println("cache miss for:" + url)
		return NotCached, ""
	}

}

//this is our in-memory data, prod real db fetching logic will be applied here
func IsURLExistsInDB(url string) bool {
	for i := 0; i < len(domains); i += 1 {
		if url == domains[i] {
			return true
		}
	}
	return false
}

// This should be managed with LRU approach in real world
// with some eviction methodology , as discussed in our theoretical talk on site interview
func CachePut(key string, val string) {
	cache[key] = val
}

func main() {
	loadMock()
	http.HandleFunc("/", DefaultHandler)
	http.HandleFunc("/validate", ValidateHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
