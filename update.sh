#!/bin/sh

yarn run clean-triggers
scrapy runspider spider.py -a filename=maps.json
./parser/index.js maps lists
rm -r maps
