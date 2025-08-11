#! /bin/bash

go test ./... || allure generate ./*/allure-results/ --clean
