Banner
======

The Banner package meets the following spectification:


## Specifications

* Each banner is associated with a promotion so it will only run for a specific period of time.
* The display period can be set individually for each banner.
* If the banner is within the display period, it display the banner.
* Only one banner can be displayed at a time.
* If more than one banner are active at any time, the banner with early expiry time will be displayed.
* A banner is either globally active or inactive.
* Once the banner displayed period is expired, it is no longer displayed.
* The banner package provides preview functions which can be used by internal QA team to preview the banner before it is active.


## Project Status
This project is in beta and need to be thorougly tested by QA before it can be used in the production.  There are no known bugs. 

## Concurrency
The project needs little modification in order to be used by go routines. This can be done easily by adding synchronization using Go "sync" package.


## Table of Contents

- [Installation](#installation)
- [Documentation](#documentation)
- [Example](#example)


## Installation
To start using Banner, copy the **banner** directory to your source code and use that import path in your application. Please also modify the import path of the banner package in the example code.


## Documentation
The **godoc** tool generates the nice documentation of this banner package.  To generate the documentation, please run the following command after copying the banner code. The documentation will be accessible by browser at [http://localhost:9999][http://localhost:9999]

[http://localhost:9999]: http://localhost:9999

```sh
$ godoc -http=:9999 
```
### Banner

The Banner struct defines the basic fields of a banner and its duration in Unix style timestamp. The banner is active at time t when ActiveAt <= t < ExprieAt. The calling function has to convert the time.Time by calling Unix() function. Because a banner can be either globally active or inactive, the server time takes care of timezone issues automatically.

```go
// Banner image
type Banner struct {
	ID          string `json:"id"`
	Filename    string `json:"filename"`
	ActiveAt    int64  `json:"active_at"`
	ExpireAt    int64  `json:"expire_at"`
	Image       []byte `json:"image"`
	Description string `json:"description"`
}
```

### Collection

A Collection is a collection of scheduled banners. A user can create as many collections as needed.  This is useful while scheduling banners based on the target audience.  For example, a Japanese collection may contains japanese banners while an English collection may contains banner for English speaking countries. Every collection has its own scheduler.

```go
// Collection of scheduled banners
type Collection struct {
	ID        string
	DataStore DataStore
	Scheduler Scheduler
}
```

### DataStore
Datastore is an interface that declares the storage functions. The banner code contains an in-memory datastore **DataStoreMemory** as an example.

```go
// DataStore interface
type DataStore interface {
	Put(b Banner) error
	Get(id string) (Banner, error)
	List() []Banner
	Del(id string) error
}
```

### Scheduler

Scheduler maintains a linked-list of Timeslots. Each of the Timeslot node in the list contains its **duration** [T1, T2), list of **active** banners, that are sorted according to their expiry time. Nodes (Timeslots) from the front of the list are removed after they are expired.

```go
// Scheduler is a linked list of timeslots
type Scheduler struct {
	Head *Timeslot
}
```

At the time of scheduing a new banner, if any overlap is detected in the duration, then the node is split into two nodes.  This is demonstrated in the example code.

### Timeslot

```go
// Timeslot is valid for time t when  T1 <= t < T2
type Timeslot struct {
	Lock    bool
	T1      int64
	T2      int64
	Banners []Banner
	Next    *Timeslot
}

```



## Example
The **example** directory contains an example which creates and schedules few banners and displays the schedule in the console.  It also starts an http server at localhost:8080. At any moment in time, refresh the browser to load the current banner.  This example runs for 180 seconds after which the user must press Ctrl+C to terminate the process.

To preview the banner in the future time, please input the time in seconds (in the console).  The output will show both the current banner and banner at the specified time.  To see only the banner currently active, simply press ENTER.

Please refer to the Example's [README.md][example-readme] file for more details.

[example-readme]: example/README.md



