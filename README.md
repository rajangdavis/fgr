# fgr - serve charts via configuration

**fgr** is an idea for utilizing existing open-source charting libraries to create a tool for
generating and displaying charts based on configuration files in a secure context.

## What it will need:

1. Some Arbitrary data to chart
2. A configuration file to feed info for charting or command line args - starting with YAML then command line args, JSON (from file and/or streaming) in the future
3. A command to the CLI

## What it should do:

1. Export charts using matplotlib
2. Provide built-in server and front-end for charting that can pipe in data from the configuration file:
3. Code export:
    1. Python:
        1. raw code
        2. jupyter notebook
    2. Javascript:
        1. raw d3 script
        2. jupyter notebook
        3. html/css/js
4. Config export?

[Proposal for API can be found here](./jupyter_proofs/fgr_v0.0.1.ipynb).

