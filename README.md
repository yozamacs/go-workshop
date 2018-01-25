# Go workshop - Tech Intersections 2018

## Purpose

The purpose of this workshop is to experiment with a simple go project and learn how to cross compile go

## Get started

### Install  and Set Up Go
* Download the appropiate go binary for your operating system
[here](https://golang.org/dl/)
* Set up your GOPATH like [this](https://github.com/golang/go/wiki/SettingGOPATH)
* Create go directory, `mkdir $GOPATH`
* Create the following directories underneath your GOPATH
  * `mkdir $GOPATH/src`
  * `mkdir $GOPATH/bin`
  * `mkdir $GOPATH/pkg`

### Get the project set up
* Run `go get -u github.com/yozamacs/go-workshop` to get this code
* Navigate to the above path and open up in your favorite code editor and play around!

### Build project
* From the project directory run `go build .`
* This will create an executable in the project directory
* Run the executable with the command `./go-workshop`
  * You can run the executable with various options including
    * `-image_url=<any png,jpg/jpeg,gif image link>`
    * `-filter_list=<one or more of the following: grayscale,invert,pixelate> `

#### Ideas for tinkering
* Add more filters, see the options [here](https://github.com/disintegration/gift#filter-examples)
* Add a command line option that lists possible filters for users
* Try loading the image from a file instead of a link

## Miscellaneous
### Fun Image Links
* https://i.imgur.com/5JWde5K.jpg -> African man on phone
* https://i.imgur.com/Ed4LdEW.jpg -> question mark guy
* https://i.imgur.com/FB5IGMg.png -> why you always lying

### Ideas for expanding the workshop
* Creating a way for the app to take requests to make it an easy backend for a front end framework
* Using the [imaging](https://github.com/disintegration/imaging) library to combine various pictures into one picture
* Find a library that does facial recognition to get more control over filters
