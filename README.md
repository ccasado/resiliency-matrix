# Resiliency Matrix


### Start the web server:

   ```
   revel run -a resiliency-matrix
   ```

   Go to http://localhost:9000/ 


## Code Layout

The directory structure of a generated Revel application:

    conf/             Configuration directory
        app.conf      Main app configuration file
        routes        Routes definition file

    app/              App sources
        init.go       Interceptor registration
        controllers/  App controllers go here
        views/        Templates directory

    messages/         Message files

    public/           Public static assets
        css/          CSS files
        js/           Javascript files
        images/       Image files

    tests/            Test suites


### API Methods

## Add Service

   ```
   curl --header "Content-Type: application/json" \
   --request POST \
   --data '{"name":"Show"}' \
   http://localhost:9000/service
   ```

## Delete Service

   ```
   curl -X DELETE http://localhost:9000/service/id
   ```

### Help

* The [Getting Started with Revel](http://revel.github.io/tutorial/gettingstarted.html).
* The [Revel guides](http://revel.github.io/manual/index.html).
* The [Revel sample apps](http://revel.github.io/examples/index.html).
* The [API documentation](https://godoc.org/github.com/revel/revel).
* [CRUD API using Revel](https://github.com/mustanish/revel-crud).
