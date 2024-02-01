package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev";

func main()  {
	fmt.Println("hello")

	//make a response

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response);
	//Response is of type: *http.Response   ==> therefore, response is of type pointer
	//** this means that we are not getting any copy of the response we are actually getting the reference towards the responce


	defer response.Body.Close()	// it is caller's responsibility to close the connection

	//read the response

	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)

	fmt.Println(content)

//<!DOCTYPE html>
// <html lang="en">

// // <head>
    //<meta charset="UTF-8" />
    //<meta name="viewport" content="width=device-width, initial-scale=1.0" />
    //<meta name="theme-color" content="#0e101a">
    //<meta name="image" property="og:image" content="https://lco.dev/images/hitesh.png" />
    //<meta name="description" property="og:description" content="About Hitesh Choudhary and LearnCodeOnline | link in bio">
    //<meta name="author" content="Hitesh Choudhary">

    //<!-- CSS only -->
// <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.1/dist/css/bootstrap.min.css" rel="stylesheet" >
// // <link rel="stylesheet" href="short.css">
    //<title>LCO Dev</title>
// </head>

// <body>
    // <div class="row mt-3">
        // <div class="col-12 text-center">
            // <img src="./images/hitesh.png" width="300px" class="rounded mx-auto border border-light d-block" alt="Hitesh">
            // <h1 class="display-4 mt-2 heading">Hitesh Choudhary</h1>
            // <p class="text-white">I write code and make high quality tech videos</p>
        // </div>
        
    // </div>
    // <div class="row m-1">
        // <div class="col-sm-12  text-center">
            // <a href="https://courses.learncodeonline.in/learn">
            //  <div class="card border border-white bg-dark mx-auto  mb-3" style="max-width: 500px;">
                //  <div class="row g-0">
                //    <div class="col-4">
                    //  <img src="./images/logo-white.png" class="img-fluid rounded-start" alt="..." width="120px">
                //    </div>
                //    <div class="col-8">
                    //  <div class="card-body">
                    //    <h5 class="card-title">LearnCodeOnline</h5>
                    //    <p class="card-text">All of my courses that I teach.</p>
                    //  </div>
                //    </div>
                //    <p class="card-text my-2"><small class="text-white">https://courses.learncodeonline.in</small></p>
                //  </div>
            //    </div>
            //  </a>
        //  </div>
    // </div>//

	//.............................................................................................................

// </body>

// </html>
}