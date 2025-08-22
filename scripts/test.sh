#! /bin/bash

rm -r ./*/allure-results/
go test -count 3 -timeout 5s ./...
cp -r ./allure-report/history ./allure-results/ && echo "Copied history."
./scripts/write_allure_env_props.sh
allure generate ./allure-results/ ./*/allure-results/ --clean
