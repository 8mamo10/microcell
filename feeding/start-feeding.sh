#!/bin/bash
NUM=`ps aux | grep python3 | grep -v grep | wc -l`
if [ $NUM -gt 0 ]; then
        echo "Already running."
        exit
fi
/usr/bin/python3 /home/pi/git/microcell/feeding/servo.py
