##This is my golang code example written in 2 days for test on interview.

[Test task description](task.pdf)

# Start service
    Usage `go run ./main.go -port=3000`

# Campaign Data Generator
    curl -XGET 'http://host:3000/campaign?x={number}&y={number}&z={number}'
    The generator has parameters: X (X <= 100) , Y (Y =< 26), Z (Z <= 10000)

# Generate user
    curl -XGET 'http://host:3000/user'

# Campaign import
    curl -XPOST 'http://host:3000/import_camp' -d '{CAMPAIGN_JSON_DATA}'
 
# Campaign searching by user info (user targets matched campaign attributes)
    curl -XPOST 'http://host:3000/search' -d '{USER_JSON_DATA}'

# Create users and do campaign search
    curl -XGET 'http://host:3000/search_auto'

## Performance testing
  Test data: 1000 campaigns with up to 10 target lists with up to 50 attributes.

    curl -XGET 'http://host:3000/campaign?x=50&y=10&z=1000'

  Start test (wrk should be installed):

    wrk http://localhost:3000/search_auto -c64 -d10s

  Output on model A1502 (Apple MacBook Pro A1502 (October, 2013)):

    Running 10s test @ http://localhost:3000/search_auto
      2 threads and 64 connections
      Thread Stats   Avg      Stdev     Max   +/- Stdev
        Latency     3.77ms    4.10ms  72.57ms   89.92%
        Req/Sec    10.79k     1.20k   13.61k    75.50%
      214974 requests in 10.02s, 29.21MB read
    Requests/sec: 748868.53
    Transfer/sec:  606.33MB