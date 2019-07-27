# fgr - cli for generating charts via configuration

# PLEASE NOTE THAT THE API IS STILL UNDER CONSTRUCTION
# NOTHING BELOW WORKS AS OF 7/26/19

fgr a few different components:
1. A CLI to run everything
2. a configuration file to feed info for charting or command line args - either JSON or YAML
3. exports charts using matplotlib or D3
4. a built-in server for charting that can pipe in data from the configuration file

## Installation

1. Clone from github
``` bash
  $ git clone https://github.com/rajangdavis/fgr.git
```

2. Go to project directory and build with docker
``` bash
  $ cd ./fgr && docker build -t fgr .
```

3. Run the CLI
``` bash
  $ docker run --rm -ti fgr -f data.csv
```

4. (OPTIONAL) Save the file as executable and add file to path for standalone CLI
``` bash
#!/bin/bash
# gnu-date - a wrapper script for invoking `date(1)` from within a Docker image
docker run --rm -ti fgr "$@"
```
``` bash
  $ fgr -f config.yml
```

## Basic API

``` bash
  $ fgr -f <config-file>
```