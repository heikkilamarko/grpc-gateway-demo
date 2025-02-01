#!/bin/bash
set -e

echo ----------

echo 🟢 valid add request:
nats request calculator.v1.add '{ "a": 100, "b": 20 }' --raw | jq
echo ----------

echo 🟢 valid subtract request:
nats request calculator.v1.subtract '{ "a": 100, "b": 20 }' --raw | jq
echo ----------

echo 🔴 divide operation not supported:
nats request calculator.v1.divide '{ "a": 100, "b": 20 }' --raw | jq
echo ----------

echo 🔴 property b missing:
nats request calculator.v1.add '{ "a": 100 }' --raw | jq
echo ----------

echo 🔴 properties a and b missing:
nats request calculator.v1.add '{}' --raw | jq
echo ----------

echo 🔴 additional property c is not allowed:
nats request calculator.v1.add '{ "a": 100, "b": 20, "c": 30 }' --raw | jq
echo ----------

echo 🔴 property a is not a number:
nats request calculator.v1.add '{ "a": "xyz", "b": 20 }' --raw | jq
echo ----------

echo 🔴 request body is not in json format:
nats request calculator.v1.add '<xml><a>100</a><b>20</b></xml>' --raw | jq
echo ----------

echo 🔴 request body is empty:
nats request calculator.v1.add '' --raw | jq
echo ----------
