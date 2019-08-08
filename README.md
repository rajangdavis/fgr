# fgr - serve charts via configuration

## What will **fgr** *do*?

### PLEASE NOTE THAT THE API IS STILL UNDER CONSTRUCTION; THIS IS A DESIGN DOCUMENT; NOTHING BELOW WORKS, IT'S ALL HYPOTHETICAL

**fgr** will have a few different components:

1. A CLI to run everything
2. a configuration file to feed info for charting or command line args - starting with YAML then command line args, JSON (from file and/or streaming) in the future
3. exports charts using matplotlib or D3
4. **Needs to be security focused** - CLI is run in container, should only allow specific access to host files and users especially in a server and file creation context

Pie-in-the-sky stuff:

1. a built-in server for charting that can pipe in data from the configuration file:
    1. Figure out some protocol for real-time data streaming
    2. Expose server via docker
2. code export:
    1. Python:
        1. raw code
        2. jupyter notebook
    2. Javascript:
        1. raw d3 script
        2. jupyter notebook
        3. html/css/html

## Installation (this is a proposal, not 100% set in stone)

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
docker run --rm -ti --volume "$PWD":/shared/ fgr "$@" -p 5801:3000
```

Considering that we are using Docker, it might be prudent to include build instructions for images.

## Basic API and YML configuration

[Proposal for API can be found here](./jupyter_proofs/fgr%20v0.0.1.ipynb).

## Running a server

``` bash
  $ fgr run -p 3000 -f multiple_stacked_bar_charts.yml
```