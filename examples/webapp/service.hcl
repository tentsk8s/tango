name = "web-frontend"
replicas = 100

// the web server
container {
    name = "arschles/frontend"
    tag = "v0.1.0"
    port = 3000
    env = ["STORAGE_TYPE=disk", "ATHENS_DISK_STORAGE_ROOT=/tmp"]
}

// sidecar for log aggregation
container {
    name = "arschles/logs"
    tag = "v0.1.1"
    env = ["BATCH_INTERVAL_SEC=10"]
    // this is where the dockerfile for the logs program lives.
    // optional - if it's missing, it's in the root dir
    dockerfile = "./logger"
    // this is where the root dir for the logs program live
    // optional - if it's missing, it's the root dir
    docker-root = "./logger"
}
