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
docker run --rm -ti --volume "$PWD":/shared/ fgr "$@" -p 5801:3000
```

gets mapped to:

``` bash
  $ fgr -f config.yml
```

Considering that we are using Docker, it might be prudent to include build instructions for images.

## Basic API and YML configuration

This is the proposed method of charting from a YML configuration

1. Stacked bar example
```yaml
# stacked_bar_chart.yml

kind: bar
group_by:
  columns:
    - Overall Qual
    - Year Built
  sum: SalePrice
pivot:
  index: Year Built
  columns: SalePrice
stacked: true

```

``` bash
  $ fgr -f stacked_bar_chart.yml
```

This should output the following python code:

```python
df = pd.read_csv('../clean_train.csv')

bar_chart_data = df[["Overall Qual","Overall Cond", "SalePrice","Year Built","Year Remod/Add", "Yr Sold"]]

temp = bar_chart_data.groupby(['Year Built', 'Overall Qual'])[['SalePrice']].sum()

temp.reset_index(inplace=True)

stacked_bar_chart = temp.pivot(index='Year Built', 
           columns='Overall Qual', 
           values='SalePrice').plot(kind='bar', 
                                    figsize=(20,10),
                                    stacked=True)

fig = stacked_bar_chart.get_figure()
fig.savefig("./temp.png") 

```

This should output the following js code:

```node
// write JS code here

```

2. Multiple stacked bar charts example
```yaml
# multiple_stacked_bar_charts.yml

kind: bar
group_by:
  columns:
    overall: 
      - Overall Qual
      - Overall Cond
    year_data:
      - Year Built
      - Year Remod/Add
      - Yr Sold
  sum: SalePrice
pivot:
  index:
    - Year Built
    - Year Remod/Add
    - Yr Sold
  columns: SalePrice
stacked: true


```

``` bash
  $ fgr -f multiple_stacked_bar_charts.yml
```

This should output the following python code:
```python
import numpy as np
import matplotlib.pyplot as plt
import pandas as pd

# TODO - update this script

df = pd.read_csv('../clean_train.csv')

bar_chart_data = df[["Overall Qual","Overall Cond", "SalePrice","Year Built","Year Remod/Add", "Yr Sold"]]

temp = bar_chart_data.groupby(['Year Built', 'Overall Qual'])[['SalePrice']].sum()

temp.reset_index(inplace=True)

stacked_bar_chart = temp.pivot(index='Year Built', 
           columns='Overall Qual', 
           values='SalePrice').plot(kind='bar', 
                                    figsize=(20,10),
                                    stacked=True)

fig = stacked_bar_chart.get_figure()
fig.savefig("./temp.png") 

```

This should output the following js code:

```node
// write JS code here

```

## Running a server

``` bash
  $ fgr run -p 3000 -f multiple_stacked_bar_charts.yml
```