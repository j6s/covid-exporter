# COVID exporter

Exports current covid case data from https://covid19api.com in a prometheus format.

## Usage
### Binary
Download the binary from the latest release and run it.
```
$ ./covid-exporter
2021/01/11 22:05:58 Starting to listen on :9084
```

### Docker
```
$ docker run thej6s/covid-exporter
2021/01/11 22:05:58 Starting to listen on :9084
```

## Example metrics

```
covid_cases{country="AD",status="confirmed"} 8586
covid_cases{country="AD",status="dead"} 85
covid_cases{country="AD",status="recovered"} 7724
covid_cases{country="AE",status="confirmed"} 230578
covid_cases{country="AE",status="dead"} 708
covid_cases{country="AE",status="recovered"} 206114
covid_cases{country="AF",status="confirmed"} 53489
[...]
```
