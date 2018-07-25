# CICD project (drone + docker + github)

star me

## Install Drone

Please see the [drone tutorial](https://github.com/go-training/drone-tutorial).

## Install docker-compose
Please see the [docker-compose](https://docs.docker.com/compose/install/#prerequisites)
## Run
```
make clean && make && make docker && docker-compose up -d && make test
```
## issue
if you can't get data. please see if set data into redis successful or not.
if not success. do :

```
redis-cli -h 127.0.0.1 -p 6379
SET NewYork Business
SET Michale Person
SET NewYork Business
SET Apple business
```

## Use

use postman to test :
![result](https://raw.githubusercontent.com/LeyouHong/telo_demo/master/result.png)

## CICD result
![result](https://raw.githubusercontent.com/LeyouHong/telo_demo/master/cicd.png)

## 14M file test result
![result](https://raw.githubusercontent.com/LeyouHong/telo_demo/master/14M_file_test_result.jpeg)
