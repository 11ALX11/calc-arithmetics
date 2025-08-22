#! /bin/bash

# Get the Go version
go_version=$(go version 2>/dev/null)
if [ $? -ne 0 ]; then
    go_version="Go not installed or command failed"
fi

# Get system information
operating_system=$(uname -o)
os_name=$(uname -s)
os_release=$(uname -r)
os_version=$(uname -v)
hardware_platform=$(uname -p)

# Write to the environment.properties file
cat <<EOF > ./allure-results/environment.properties
go_version = ${go_version}
operating_system = ${operating_system}
os_name = ${os_name}
os_release = ${os_release}
os_version = ${os_version}
hardware_platform = ${hardware_platform}
EOF
