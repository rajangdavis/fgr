# fgr - cli for generating charts via configuration

# PLEASE NOTE THAT THE API IS STILL UNDER CONSTRUCTION; THIS IS A DESIGN DOCUMENT; NOTHING BELOW WORKS, IT'S ALL HYPOTHETICAL

fgr will have a few different components:
1. A CLI to run everything
2. a configuration file to feed info for charting or command line args - starting with YAML then command line args, JSON (from file and/or streaming) in the future
3. exports charts using matplotlib or D3
4. Needs to be security focused - CLI is run in container, should only allow specific access to host files and 

Pie-in-the-sky stuff:
1. a built-in server for charting that can pipe in data from the configuration file
	1. Figure out some protocol for real-time data streaming
	2. Expose server via docker
2. code export
	1. Python
		1. raw code
		2. jupyter notebook
	2. Javascript
		1. raw d3 script
		2. jupyter notebook
		3. html/css/html

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

4. (OPTIONAL) Save the file as executable and add file to $PATH or alias in .bash_profile for standalone CLI
``` bash
#!/bin/bash
docker run --rm -ti --volume "$PWD":/shared/ fgr "$@"
```
``` bash
  $ fgr -f config.yml
```

Considering that we are using Docker, it might be prudent to include build instructions for images.

## Basic API and YML configuration

This is the proposed method of charting from a YML configuration

1. 
```yaml
type: bar
x_axis: Time 
y_axis: Money
stacked: true

```

``` bash
  $ fgr -f <config-file>
```
