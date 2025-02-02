#!/bin/bash
set -e

cd "$(dirname "$0")"

k6 run calculator_http_stress_test.js
