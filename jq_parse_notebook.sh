#!/bin/bash
cat ./jupyter_proofs/fgr_v0.0.1.ipynb | ~/Desktop/jq-win64.exe '[.cells[] | select(.cell_type | contains("code")) | .source | join("\n")] | join("\n")'