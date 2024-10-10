package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type RequestData struct {
	TaskName        string `json:"name"`
	TaskDescription string `json:"description"`
}

type Task struct {
	TaskName        string
	TaskDescription string
}

// var tasks []Task

var ctx = context.Background()

func TaskHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "unalbe to read request Body", http.StatusBadRequest)
			return
		}

		var data RequestData

		err = json.Unmarshal(body, &data)

		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		if data.TaskName == "" || data.TaskDescription == "" {
			http.Error(w, "Need to have both description and name at least", http.StatusBadRequest)
			return
		}

		rdb := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})

		t := time.Now()
		keySet := "simpletask:" + strconv.FormatInt(t.UnixMilli(), 10)
		dberr := rdb.Set(ctx, keySet, body, 0).Err()
		if dberr != nil {
			panic(dberr)
		}

		// val, err := rdb.Get(ctx, "key").Result()
		// if err != nil {
		// 	panic(err)
		// }
		// fmt.Println("key", val)

		// val2, err := rdb.Get(ctx, "key2").Result()
		// if err == redis.Nil {
		// 	fmt.Println("key2 does not exist")
		// } else if err != nil {
		// 	panic(err)
		// } else {
		// 	fmt.Println("key2", val2)
		// }
		fmt.Fprintf(w, "added tasks %s", keySet)
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}

}
