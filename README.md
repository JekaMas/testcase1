# Campaign import and search service
    Usage `go run ./main.go -port=3000`

## Performance
    rk ­c 64 ­d 10s http://localhost:3000/search_auto

  Output on X=50, Y=10, Z=1000:

    Running 10s test @ http://localhost:3000/search_auto
      2 threads and 64 connections
      Thread Stats   Avg      Stdev     Max   +/- Stdev
        Latency     3.77ms    4.10ms  72.57ms   89.92%
        Req/Sec    10.79k     1.20k   13.61k    75.50%
      214974 requests in 10.02s, 29.21MB read
    Requests/sec: 748868.53
    Transfer/sec:  606.33MB