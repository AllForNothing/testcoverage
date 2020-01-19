#!/bin/bash

env
echo token:${COVERALLS_TOKEN//-/,}
echo len:${#COVERALLS_TOKEN}
echo part1:${COVERALLS_TOKEN:0:5}
echo part2:${COVERALLS_TOKEN:1}
