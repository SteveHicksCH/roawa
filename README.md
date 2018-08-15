# A Practice Go application
This is a vanilla web application to:
* enter the company number
* get the current registered address from the public API 
* change of registered address on a web form
* Maybe use the change of registed address service running locally to go the update

# Package Structure
Experiment using https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1 

# Building
make build

# Set up and configuration

## Configuration

| Environment variable |  Description | Default Value
| -------------------- | ----------- | --------------
| `ROA_WEBAPP_API_KEY` | API key authorising the application to access the CHS API |
| `ROAWA_LISTEN_ADDRESS`| Listen Address and Port | 127.0.0.1:8080
| `ROAWA_TEMPLATES_DIR`| Directory for application HTML templates | ${GOPATH}/src/github.com/shicks/roawa/http/templates


# Running
On Command line:
make run

In Browser:
http://localhost:8080/roawa/edit

For CHS vagrant deploy - http://chs-dev:4000/roawa/edit

Note for Firefox:
Check that in the Network settings is ticked "Bypass proxy server for your local address"
In some versions of FF, that Tick is not present apart for "Manual proxy configuration" (ie.
not present for "Auto-detect proxy settings"). In some versions of FF you can then fix it from "about:config" —> browser.urlbar.trimURLs=false
If that doesn't work for your FF version, then set for the time being  "Manual proxy configuration" where the disable wanted Tick becomes available.


# Web References
* https://developer.companieshouse.gov.uk/api/docs/index.html
* https://golang.org/doc/articles/wiki/ - Web Application how to
* https://blog.scottlogic.com/2017/02/28/building-a-web-app-with-go.html with git repository https://github.com/Rosalita/GoViolin 


# TODO
* New Branch feature/more-go-chs-frameworks
* New screen to add the company number
* Use alice/gorilla
* Use roa service in CHS to update registered address