docker build -t data_structure_and_algorithm:0.1 .
docker run -it -v (pwd)/source:/usr/src/source -w /use/src/source --name data_structure_and_algorithm 27a801f37d4d /bin/bash
