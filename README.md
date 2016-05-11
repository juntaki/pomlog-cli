# Productivity logger

Small CLI tool for work time logger based on Pomodoro technique.

## How to use

~~~
$ go run main.go
What number are you doing?
[0: work1][1: work2][999: exit]
0
[======================================================================] 100.00%
Take 5 min rest!
[======================================================================] 100.00%
1 pomodoro
What number are you doing?
[0: work1][1: work2][999: exit]
999
~~~

## Log format

~~~
2016-05-02T00:24:30+09:00: work1 start
2016-05-02T00:24:46+09:00: work1 start
~~~
