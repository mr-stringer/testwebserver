# testwebserver

A  basic static web server that prints the hostname that it runs on.

The version v1.0.0 includes a Linux AMD64 binary, a systemd unit file and cloud-init file, making it easy to integrate install during cloud provisioning and automatically run.

By default, testwebserver will listen on port 8000, but can be configured to run on any port with -p flag.
