# Macaddress Cli

This is simple cli tool to retrieve the vendor or a given mac address.


## Getting Started

These instructions will help you get your Macaddress Cli up and running on your local machine for development and testing purposes.

### Prerequisites

The project require docker to be installed and an API key to authenticate with macaddress.io. You can get docker [here](https://www.docker.com/get-started). Get your macaddress.io API key from your [account settings](https://macaddress.io/account/general) on macaddress.io. 

### Installing

##### Get Macaddress Cli
```
git clone [git-repo-url] macaddress_cli
cd macaddress_cli
```

##### Update API_KEY
Edit `Dockerfile` file to specify API_KEY obtained from macaddress.io. 

##### Create docker image
```
docker build --tag=macaddress_cli .
```

### Usage

#### Run docker image and execute macaddress_cli 
```
docker run -it --rm macaddress_cli bash
./macaddress_cli <mac_address> // Replace <mac_address> with actual mac address
```
