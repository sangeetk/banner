Banner
======

The Banner package meets the following spectification:


## Specifications

* Each banner is associated with a promotion so it will only run for a specific period of time.
* The display period can be set individually for each banner.
* If the banner is within the display period, display the banner.
* Only one banner can be displayed at a time.
* If more than one banner are active at any time, the banner with early expiry time will be displayed.
* A banner either globally active or inactive.
* Once the banner displayed period is expired, it is no longer displayed.
* The banner package provides preview functions which can be used by internal QA team to preview the banner before it is active.


## Project Status
This project is in beta and need to be thorougly tested by QA before it can be used in the production.  There are no known bugs. 

## Concurrency
The project needs little modification in order to be used by go routines. This can be done easily by adding synchronization using Go "sync" package.


## Table of Contents

- [Getting Started](#getting-started)
 - [Installation](#installation)
 - [Documentation](#documentation)
 - [Example](#example)


## Getting Started

### Installation
To start using Banner, copy the **banner** directory to your source code and use that import path in your application.


### Documentation
The **godoc** tool generates the nice documentation of this banner package.  To generate the documentation, please run the following command after copying the banner code. 

```sh
$ godoc -http=:9999 
```

The documentation will be accessible by browser at [http://localhost:9999][http://localhost:9999]

[http://localhost:9999]: [http://localhost:9999]

### Example
The **example** directory contains an example which creates and schedules few banners and displays the schedule in the console.  It also starts an http server at localhost:8080. At any moment in time, refresh the browser to load the current banner.  This example runs for 180 seconds after which the user must press Ctrl+C to terminate the process.

Please refer to the Example's README.md file for more details.




