#!/bin/bash
set -e

cd "$(dirname "$0")"

nats request calculator.v1.add '{ "a": 100, "b": 20 }' --count 1_000_000
