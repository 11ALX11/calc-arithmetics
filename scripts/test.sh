#! /bin/bash

rm -r ./*/allure-results/
go test -count 3 -timeout 5s ./...
cp -r ./allure-report/history ./allure-results/ && echo "Copied history."
allure generate ./allure-results/ ./*/allure-results/ --clean
