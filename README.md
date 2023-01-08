## mathematician-in-wonderland

HISCON 2022 MISC 문제

### 로컬에서 구동하기

```
$ git clone https://github.com/minpeter/mathematician-in-wonderland.git
$ cd mathematician-in-wonderland
$ go run .
```

### Docker로 구동하기

```
$ git clone https://github.com/minpeter/mathematician-in-wonderland.git
$ cd mathematician-in-wonderland
$ docker build -t mathematician-in-wonderland .
$ docker run -dp 5555:5555 mathematician-in-wonderland
# nc localhost 5555
```
