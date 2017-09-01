# github-clone
GitHub Automated Clone

## Update Go in C9
https://community.c9.io/t/writing-a-go-app/1725
```
sudo rm -rf /opt/go
wget https://storage.googleapis.com/golang/go1.9.linux-amd64.tar.gz
sudo tar -C /opt -xvf go1.9.linux-amd64.tar.gz
go version
```
## Install Packages
```
go get -u golang.org/x/oauth2
```
## Create Config File
```
config.json
{
    "GITHUB_KEY" : "YOUR_KEY"
}
```
## Run App
```
go run main.go 
```

## Get Output Directory Size
```
du -sh output/
```

## List Counts of All File Types
```
cd output
find . -type f | sed 's/.*\.//' | sort | uniq -c
```

## Copy All Files to Another Directory
```
cd output
mkdir ps1
find . -type f -name "*.ps1" -exec cp {} ps1 \;
```

## How Many Files are in a Directory?
```
cd output
ls ps1 | wc -l
```