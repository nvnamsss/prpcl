#!/bin/bash
GOPATH=${PWD}
RED=`tput setaf 1`
GRN=`tput setaf 2`
PUR=`tput setaf 13`
RESET=`tput sgr0`
cd src
TEST_RESULT_DIR=./test-results
mkdir -p ${TEST_RESULT_DIR}
echo "----------------"

go test -v -covermode=count ./services/... -coverprofile ${TEST_RESULT_DIR}/.testCoverage.txt | tee ${TEST_RESULT_DIR}/test.log; echo ${PIPESTATUS[0]} > ${TEST_RESULT_DIR}/test.out

cat ${TEST_RESULT_DIR}/test.log | go get -u github.com/jstemmer/go-junit-report > ${TEST_RESULT_DIR}/report.xml

echo "----------------"
echo "${GRN}Result:${RESET}"
go tool cover -func ${TEST_RESULT_DIR}/.testCoverage.txt
exit $(cat ${TEST_RESULT_DIR}/test.out)
