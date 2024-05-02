# Hostless-Beego-Example

Beego is used for rapid development of enterprise application in Go, including RESTful APIs, web apps and backend services. Check out their official documentation [here](https://doc.meoying.com/en-US/beego/developing/bee/#installing-bee-tool)

## Deploy a Beego App

### Deployment Instructions

In this tutorial, we advice creating a Dockerfile and using Docker as the build system

1. Fork/Clone this [Hostless-Beego-Example](https://github.com/Hostless-Examples/Hostless-Beego-Example.git) repo from github
2. Click on 'Create New App'
3. Choose a suitable app name
4. Choose your github account
5. Choose the forked github repo/the cloned remote repo
6. Choose a build system

    1. 'Detect automatically with Nixpicks' - automatically detects the programming language and builds the app
    2. 'Docker' - looks for a Dockerfile in the root of the project and build based on the instructions

7. Input a start command(optional)
8. The PORT environment variable is set by Hostless

#### Sample configuration
![sample](https://res.cloudinary.com/do58rrxug/image/upload/v1714681372/Screenshot_2024-05-02_at_21.22.32_rziqun.png)

#### Example project
An example project is hosted on [https://hostless-beego-example.hostless.app/](https://hostless-beego-example.hostless.app/)

