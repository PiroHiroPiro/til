#!/bin/sh

docker build -t data_structure_and_algorithm_python .
docker run -dit -v $(pwd)/source:/usr/src/source -w /usr/src/source --name data_structure_and_algorithm_python $(docker images | grep data_structure_and_algorithm_python | awk '{print $3}') /bin/bash
docker stop data_structure_and_algorithm_python
