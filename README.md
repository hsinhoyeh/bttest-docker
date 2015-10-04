##### bttest-docker
bttest[1] is a in-memory mock server for bigtable cluster.

It minic the behavior of bigtable cluster in a very accurate way and can be easily used for daily development.

bttest-docker dockerized bttest, where it can be used anywhere.
##### Build

```
docker build -t <repo>:bttest:latest .
```

##### Run
```
docker run <repo>:bttest:latest // where 9001 port is listened
```

##### Reference
[1] https://github.com/GoogleCloudPlatform/gcloud-golang/tree/master/bigtable/bttest
