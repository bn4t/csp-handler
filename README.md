# Web-Reports-Handler [WIP]
Handler that implements endpoints for the [W3C Reporting API](https://w3c.github.io/reporting)

Currently supports following report types:

- csp reports
- network error logging


# Installation
0. Install golang (>=1.15) and GNU make if you don't have them already
1. Clone the repository: `git clone https://git.bn4t.me/bn4t/csp-handler.git`
2. Checkout the latest stable tag
3. Run `make build` to build the csp-handler binary
4. Run `sudo make install` to install csp-handler on your system. This will create the directory `/etc/csp-handler` (config directory). Additionally the user `csp-handler` will be created.
5. If you have systemd installed you can run `sudo make install-systemd` to install the systemd service. Run `service csp-handler start` to start the csp-handler service. Csp-handler will automatically run as the `csp-handler` user.

Make sure you edit the config located at `/etc/csp-handler/config.toml` before running the service.

## Command line flags
- `-config <config file>` - The location of the config file to use. Defaults to `config.toml` in the working directory.

# Deinstallation
Run `sudo make uninstall` to uninstall csp-handler. This will remove `/etc/csp-handler` if the directory is empty.

Run `sudo make uninstall-systemd` to remove the systemd service.

# Usage
Include the `report-uri` directive in your content security policy:

`report-uri https://csp-report.example.com/report-uri/mydomain.com`

Replace `csp-report.example.com` with the domain on which csp-report is deployed and `mydomain.com` with the domain on which the *content security policy* is deployed.

## License
GPLv3
