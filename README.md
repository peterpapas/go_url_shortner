# URL shortner

This is a simple web application that allows users to shorten long URLs into shorter ones. It is written in Go and uses HTML and CSS for the user interface.

## Starting the Application

To start the application, simply run the following commands in your terminal:

``` go build shortener.go ```

Followed by:

```./shortener ```

Alternatively you can click the newly build ```shortener.exe```

The application will start running on http://localhost:8080.

## Using the Application

To shorten a URL, enter the long URL in the input field on the homepage and click the "Shorten" button. The application will generate a shorter URL, which you can use to redirect to the long URL.

You can access the shortened URL by typing http://localhost:8080/<shortened-url> in your browser's address bar. If the shortened URL is valid, you will be redirected to the long URL.

## Languages Used

* Go
* HTML
* CSS
