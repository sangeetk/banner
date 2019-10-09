# Banner Example

## Install
Download the banner source code and correct the import path inside the example's **main.go** and **debug.go** file.

## Compile

To compile the example, simily run **go build** inside the example directory.

```sh
➜  banner git:(master) ✗ cd example
➜  example git:(master) ✗ go build
➜  example git:(master) ✗
```

## Run
Simply run the binary without any arguments.  This will show the following output in the terminal and also start the http server at localhost:8080.  The output displays the Scheduler linked-list of Timeslots and a graph like output of scheduled banners.  You can enter the time (in sec) to see which banner is scheduled at that time.

White spaces in the **Scheduling** line represents the tiem when no banner is scheduled.  The default banner is shown as "**.**" (dot) which has duration of [0,180) in seconds from the current time.


```sh
➜  example git:(master) ./example
Testing scheduling of banners...
Head -> {'R', 0-20} -> nil
Timeline: (00)|----+----|----+----|(20)
Scheduling:   RRRRRRRRRRRRRRRRRRRR

Head -> {'R', 0-20} -> {'B', 40-60} -> nil
Timeline: (00)|----+----|----+----|----+----|----+----|----+----|----+----|(60)
Scheduling:   RRRRRRRRRRRRRRRRRRRR                    BBBBBBBBBBBBBBBBBBBB

Head -> {'R', 0-20} -> {'G', 25-40} -> {'G', 40-45} -> {'B', 45-60} -> nil
Timeline: (00)|----+----|----+----|----+----|----+----|----+----|----+----|(60)
Scheduling:   RRRRRRRRRRRRRRRRRRRR     GGGGGGGGGGGGGGGGGGGGBBBBBBBBBBBBBBB

Head -> {'R', 0-20} -> {'G', 25-40} -> {'G', 40-45} -> {'B', 45-60} -> {'Y', 70-100} -> nil
Timeline: (00)|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|(100)
Scheduling:   RRRRRRRRRRRRRRRRRRRR     GGGGGGGGGGGGGGGGGGGGBBBBBBBBBBBBBBB          YYYYYYYYYYYYYYYYYYYYYYYYYYYYYY

Head -> {'R', 0-20} -> {null, 25-40} -> {'B', 40-45} -> {'B', 45-60} -> {'Y', 70-100} -> nil
Timeline: (00)|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|(100)
Scheduling:   RRRRRRRRRRRRRRRRRRRR                    BBBBBBBBBBBBBBBBBBBB          YYYYYYYYYYYYYYYYYYYYYYYYYYYYYY

Head -> {'R', 0-20} -> {null, 25-40} -> {'B', 40-45} -> {'B', 45-60} -> {'Y', 70-90} -> {'Y', 90-100} -> {'O', 100-120} -> nil
Timeline: (00)|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|(120)
Scheduling:   RRRRRRRRRRRRRRRRRRRR                    BBBBBBBBBBBBBBBBBBBB          YYYYYYYYYYYYYYYYYYYYYYYYYYYYYYOOOOOOOOOOOOOOOOOOOO

Head -> {'R', 0-20} -> {null, 25-35} -> {'G', 35-40} -> {'B', 40-45} -> {'B', 45-60} -> {'G', 60-65} -> {'Y', 70-90} -> {'Y', 90-100} -> {'O', 100-120} -> nil
Timeline: (00)|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|(120)
Scheduling:   RRRRRRRRRRRRRRRRRRRR               GGGGGBBBBBBBBBBBBBBBBBBBBGGGGG     YYYYYYYYYYYYYYYYYYYYYYYYYYYYYYOOOOOOOOOOOOOOOOOOOO

Head -> {'R', 0-20} -> {null, 25-35} -> {'G', 35-40} -> {'B', 40-45} -> {'B', 45-60} -> {'G', 60-65} -> {'Y', 70-90} -> {'Y', 90-100} -> {'O', 100-120} -> {'P', 120-130} -> nil
Timeline: (00)|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|(130)
Scheduling:   RRRRRRRRRRRRRRRRRRRR               GGGGGBBBBBBBBBBBBBBBBBBBBGGGGG     YYYYYYYYYYYYYYYYYYYYYYYYYYYYYYOOOOOOOOOOOOOOOOOOOOPPPPPPPPPP

Head -> {'M', 0-10} -> {'R', 10-20} -> {null, 25-35} -> {'G', 35-40} -> {'B', 40-45} -> {'B', 45-60} -> {'G', 60-65} -> {'Y', 70-90} -> {'Y', 90-100} -> {'O', 100-120} -> {'P', 120-130} -> nil
Timeline: (00)|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|(130)
Scheduling:   MMMMMMMMMMRRRRRRRRRR               GGGGGBBBBBBBBBBBBBBBBBBBBGGGGG     YYYYYYYYYYYYYYYYYYYYYYYYYYYYYYOOOOOOOOOOOOOOOOOOOOPPPPPPPPPP

Head -> {'M', 0-10} -> {'R', 10-20} -> {null, 25-35} -> {'G', 35-40} -> {'B', 40-45} -> {'B', 45-60} -> {'G', 60-65} -> {'Y', 70-90} -> {'Y', 90-100} -> {'O', 100-120} -> {'P', 120-130} -> {'m', 160-180} -> nil
Timeline: (00)|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|(180)
Scheduling:   MMMMMMMMMMRRRRRRRRRR               GGGGGBBBBBBBBBBBBBBBBBBBBGGGGG     YYYYYYYYYYYYYYYYYYYYYYYYYYYYYYOOOOOOOOOOOOOOOOOOOOPPPPPPPPPP                              mmmmmmmmmmmmmmmmmmmm

Head -> {'M', 0-10} -> {'R', 10-20} -> {'.', 20-25} -> {'.', 25-35} -> {'G', 35-40} -> {'B', 40-45} -> {'B', 45-60} -> {'G', 60-65} -> {'.', 65-70} -> {'Y', 70-90} -> {'Y', 90-100} -> {'O', 100-120} -> {'P', 120-130} -> {'.', 130-160} -> {'m', 160-180} -> nil
Timeline: (00)|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|----+----|(180)
Scheduling:   MMMMMMMMMMRRRRRRRRRR...............GGGGGBBBBBBBBBBBBBBBBBBBBGGGGG.....YYYYYYYYYYYYYYYYYYYYYYYYYYYYYYOOOOOOOOOOOOOOOOOOOOPPPPPPPPPP..............................mmmmmmmmmmmmmmmmmmmm

2019/10/09 16:39:48 Running http server at localhost:8080

Preview time in sec : 10
Time [6 sec] => Active startTime : [M] , Active at 10 sec : [R]

Preview time in sec : 90
Time [9 sec] => Active startTime : [M] , Active at 90 sec : [Y]

Preview time in sec : 180
Time [11 sec] => Active startTime : [R] , Active at 180 sec : [none]

Preview time in sec : ^C
➜  example git:(master) ✗
```

To terminate the process, press Ctrl+C