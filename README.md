# A comparison of RPC methods

## Candidates

 - Native GRPC
 - GRPC-Gateway
 - GraphQL
 - HTTP + JSON + http.Handler
 - Twirp

## Service

The service is a simple echo server. You give it a message, it spits out that message

## Benchmark

To run the benchmark, run the following:

``` sh
make up
```

This is start up a docker-compose stack of the services, run the benchmarks
writing them to ./out/, and start a Jupyter notebook to visualize the results.

In the log for `make up` you will see something like the following:

``` sh
jupyter_1      |     To access the notebook, open this file in a browser:
jupyter_1      |         file:///home/jovyan/.local/share/jupyter/runtime/nbserver-7-open.html
jupyter_1      |     Or copy and paste one of these URLs:
jupyter_1      |         http://3d0e09d5942e:8888/?token=6d8c14dc988091f87dd2e9beb9504177c664db445337eb3a
jupyter_1      |      or http://127.0.0.1:8888/?token=6d8c14dc988091f87dd2e9beb9504177c664db445337eb3a
```

On MacOS, right click on the 127.0.0.1:8888 URL and click `Open URL`. You will
see a notebook called `services.ipynb`, click on it.

The report you see is from the previous run, to update the data, click on
`Kernel -> Restart & Run All` to render the report with the latest data.
