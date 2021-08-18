## stand-alone
### start daemon
```bash
docker run -d ttbb/kafka:stand-alone
```
### start it port expose
```bash
docker run -p 9092:9092 -it ttbb/kafka:stand-alone bash
```
### start daemon
```bash
docker run -p 9092:9092 -d ttbb/kafka:stand-alone
```
### start daemon mac port expose
```bash
docker run -e KAFKA_ADVERTISE_ADDRESS=localhost -p 9092:9092 -d ttbb/kafka:stand-alone
```
## raft
### start it port expose
```bash
docker run -p 9092:9092 -it ttbb/kafka:raft bash
```
### start daemon port expose
```bash
docker run -p 9092:9092 -d ttbb/kafka:raft
```
### start daemon mac port expose
```bash
docker run -e KAFKA_ADVERTISE_ADDRESS=localhost -p 9092:9092 -d ttbb/kafka:raft
```