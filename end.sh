#!/bin/bash

pid=`/bin/ps -fu $USER| grep "index" | grep -v "grep" | awk '{print $2}'`

sudo kill -9 $pid
