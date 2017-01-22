# perl-Ping-vs-golang-Ping
use Anyevent::Ping VS use Golang goroutine concurrency , both timeout is 5 second and 100 concurrency

## use
```
time ./PingAnyevent.pl 
...
Result(192.168.6.249): OK in 0.0119359493255615 seconds
Result(192.168.6.250): OK in 0.0114510059356689 seconds
Result(192.168.6.253): OK in 0.0110869407653809 seconds
Result(192.168.6.113): TIMEOUT in 5.0010199546814 seconds
Result(192.168.6.115): TIMEOUT in 5.00020503997803 seconds
...

time go run ./pingtest.go
...
192.168.6.217 true 21 <nil>
192.168.6.138 false 5000 read ip4 192.168.218.129->192.168.6.138: i/o timeout
192.168.6.135 false 5000 read ip4 192.168.218.129->192.168.6.135: i/o timeout
192.168.6.136 false 5000 read ip4 192.168.218.129->192.168.6.136: i/o timeout
192.168.6.137 false 5000 read ip4 192.168.218.129->192.168.6.137: i/o timeout
...
```

##benchmark
config timeout both are 5 second, 100 concurrency scan 254 iplists 

- ae 10秒
```
real    0m10.052s
user    0m0.054s
sys     0m0.056s

cpu 60144 root      20   0  145656   5940   2232 S  2.3  0.6   0:00.07 PingAnyevent.pl 
```

- goroutine 10秒 
```
real    0m10.378s
user    0m0.275s
sys     0m0.145s

cpu 60233 root      20   0  101124   2600   1232 S  1.0  0.3   0:00.03 pingtest   
```

##AUTHOR
zhangqi Jason asd1986_n@126.com
