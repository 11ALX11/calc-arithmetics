#! /bin/bash

go test ./app ./i18n || allure generate ./*/allure-results/ --clean
