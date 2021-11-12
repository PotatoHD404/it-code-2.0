#!/bin/sh

total=`jq '.summary.total' $1`
passed=`jq '.summary.passed' $1`
failed=`jq '.summary.failed' $1`

jq -r '.testcases[] | if .passed? then .passed="Success" else .passed="Failed" end | "\(.passed): \(.title) - \(.text)"' $1

echo
echo Total: $total
echo Passed: $passed
echo Failed: $failed

if [ "$failed" -gt "0" ]; then
  exit 1
fi
