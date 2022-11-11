# RSS-Reader-service

This service uses the already created RSS Reader package. 
The RSS Reader package is imported in this service by using the command `go get github.com/MuhammadImran696/RSS-Reader-package`
The service contains an end point named "/getdata". While making a request on this endpoint we send an array of urls in the body of the request.
that array of urls is passed to the method "Parse()" which is present in the RSS Reader package.
The Parse() method makes http.Get request on these urls and parses the data asynchronously and returns an array of RssItems.

## How to run this service locally?
The first step is to clone the service. It can be done through the following command:
`git clone https://github.com/MuhammadImran696/RSS-Reader-service.git`

After this step, We can run this service by two methods <br/>
Method 1: <br/>
By running the command <br/>
`go run main.go`

Method 2: <br/>
by running the following two commands <br/>
`docker build -t rss-service .`
`docker run -dp 9000:9000 rss-service`
### Example
#### Request
`curl --location --request POST 'http://localhost:9000/getdata' \
--header 'Content-Type: application/json' \
--data-raw '{
    "urls": [
        "http://www.cbsnews.com/latest/rss/main"
    ]
}'`

##### Response
`{
    "items": [
        {
            "Title": "Trump expected to sit for deposition in E. Jean Carroll defamation case",
            "Source": "Home - CBSNews.com",
            "SourceURL": "",
            "Link": "https://www.cbsnews.com/news/trump-deposition-e-jean-carroll-defamation/",
            "PublishDate": "2022-10-19T11:13:57-04:00",
            "Description": "Writer E. Jean Carroll accused Trump of defamation for saying she was \"totally lying\" about a sexual assault allegation."
        },
        .
        .
        .
        .
        ]
}`
