package main

import (
	"fmt"
	"net/http"
	"context"
	"github.com/redis/go-redis/v9"
)

func main() {
	client := redis.NewClient(&redis.Options{
        	Addr:	  "redis:6379",
        	Password: "",
        	DB:		  0,
    	})
    	
    	ctx := context.Background()
    	pong, err := client.Ping(ctx).Result()
    	fmt.Println(pong, err)
    	
    	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    		fmt.Fprintf(w, "hi there!")
    	})
    	
	http.HandleFunc("/set_key", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "set_key")
		r.ParseForm()
		for k, _ := range r.PostForm {
			fmt.Fprintf(w, k)
			b := r.PostFormValue(k)
			fmt.Fprintf(w, b)
			err = client.Set(ctx, k, b, 0).Err()
			if err != nil {
    				panic(err)
			}
		}
	})

	http.HandleFunc("/get_key", func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		val, err := client.Get(ctx, key).Result()
		if err != nil {
    			panic(err)
		}
		fmt.Fprintf(w, val)
	})
	
	http.HandleFunc("/del_key", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "del_key")
		r.ParseForm()
		key := r.PostFormValue("key")
		val, err := client.GetDel(ctx, key).Result()
		if err != nil {
    				panic(err)
			}
		fmt.Fprintf(w, val)	
	})

	http.ListenAndServe(":9990", nil)
}
