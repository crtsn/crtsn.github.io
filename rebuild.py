#!/usr/bin/env python3

import subprocess
import time
from pathlib import Path
import datetime

directory = Path("_ABSOLUTE_CARROT").resolve()
mtimes = {}
command = "GOOS=js GOARCH=wasm go build -o ../main.wasm"

while True:
    go_files = list(directory.rglob("*.go"))
    changed = False

    for f in go_files:
        try:
            mtime = f.stat().st_mtime
        except FileNotFoundError:
            continue

        if f not in mtimes or mtimes[f] != mtime:
            changed = True
            mtimes[f] = mtime

    if changed:
        time_str = datetime.datetime.fromtimestamp(time.time()).strftime('%H:%M:%S')
        print(f"{time_str}: Running {command}")
        subprocess.run(command, shell=True, cwd="_ABSOLUTE_CARROT")

    time.sleep(2)
