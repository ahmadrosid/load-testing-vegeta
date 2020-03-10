# Vegeta load testing

## Basic
In this basic vegeta load testing we are doing a simple load testing by running 5 seconds attack and get simple report.
```bash
echo "GET http://localhost:8080/world" | vegeta attack -duration=5s | vegeta report
```

## Running spesific rate
In this basic vegeta load testing we are doing a simple load testing by running 5 concurrent request every seconds and get simple report.
```bash
echo "GET http://localhost:8080/world" | vegeta attack -rate=100 -duration=5s | vegeta report
```

## Get report html
```bash
echo "GET http://localhost:8080/world" | vegeta attack -rate=5 -duration=5s | tee results.bin | vegeta report
cat results.bin | vegeta plot > plot.html
```

## Get report json
```bash
echo "GET http://localhost:8080/world" | vegeta attack -rate=5 -duration=5s | tee results.bin | vegeta report
vegeta report -type=json results.bin > metrics.json
```
