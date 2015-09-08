# Fruit HTML scraping
This repo provides a basic script to scrape html doc and build JSON out of useful info present within.

###Instructions to run:  
1. Setup `$GOPATH` environment variable to some directory  
2. On your terminal create the relevant directories and navigate to it by running the following command  
  > $ mkdir -p $GOPATH/src/github.com/sim4life && cd $_

3. Now clone the repo in that directory by  
  > $ git clone git@github.com:sim4life/fruit_scraping.git

4. Then download dependencies by  
  > $ go get golang.org/x/net/publicsuffix  
  > $ go get github.com/PuerkitoBio/goquery  
  > $ go get github.com/djimenez/iconv-go

4. To run  
  > $ go run fruitsJSON.go

5. To test  
  > $ go test

I hope you enjoy the show.
