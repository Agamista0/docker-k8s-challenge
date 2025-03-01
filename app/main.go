package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "os"
    "strconv"
    "github.com/go-redis/redis/v8"
    "github.com/gorilla/mux"
)

var ctx = context.Background()
var rdb *redis.Client

func db_checker() error {
    if err := rdb.Ping(ctx).Err(); err != nil {
        fmt.Println("Redis is not connected:", err)
        return err
    }
    fmt.Println("Redis is successfully connected!")
    return nil
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
    // Simple health check response
    w.WriteHeader(http.StatusOK)
    fmt.Fprintln(w, "Healthy")
}

func hello(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Wow! A new visit")

    visits, err := rdb.Get(ctx, "visits").Result()
    if err != nil {
        visits = "0"
    }

    newVisits, err := strconv.Atoi(visits)
    if err != nil {
        fmt.Println("Error converting visits to integer:", err)
        newVisits = 0
    }
    newVisits++

    rdb.Set(ctx, "visits", strconv.Itoa(newVisits), 0)

    fmt.Fprintf(w, "Hello, World! You are visitor number %d.", newVisits)
}

func main() {
    redisHost := os.Getenv("REDIS_HOST")
    redisPort := os.Getenv("REDIS_PORT")
    appPort := os.Getenv("APP_PORT")

    if redisHost == "" || redisPort == "" {
        log.Fatal("REDIS_HOST or REDIS_PORT is not set")
    }

    rdb = redis.NewClient(&redis.Options{
        Addr: fmt.Sprintf("%s:%s", redisHost, redisPort),
        Password: "",
        DB:       0,
    })

    if err := db_checker(); err != nil {
        log.Fatal(err)
    }

    router := mux.NewRouter()
    router.HandleFunc("/", hello)
    router.HandleFunc("/health", healthCheck)

    fmt.Println("Server is running on port", appPort)
    log.Fatal(http.ListenAndServe(":"+appPort, router))
}
