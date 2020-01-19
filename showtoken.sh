#!/bin/bash

env
echo token:${COVERALLS_TOKEN//-/,}
echo len:${#COVERALLS_TOKEN}
echo part:${COVERALLS_TOKEN:2:5}
