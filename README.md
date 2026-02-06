# promeser

[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)  

Prometheus metric server scenarios


## About

The repository demonstrates a simple way of Prometheus metrics written in Go.

## Features:

- create random values which are used for a metric in a way of synchron (on each http call) and asynchron (concurrent)
- collect the values as an Gauge and a Counter 
- register all Prometheus metrics
- expose the metrics as a http service on http://localhost:8080/metrics

The repository will not fulfil all the functions of a complete project. 
In order to demonstrate only the Prometheus metrics, the following elements have been removed:

- unit tests
- command cli
- decoupling the main function from the rest
- documentation

## Prometheus documentation

For further information, please look at the official Prometheus page https://prometheus.io/docs/guides/go-application/

