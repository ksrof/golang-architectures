<div id="top"></div>

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/ksrof/golang-architectures">
    <img src="docs/images/gopherworking.png" alt="Gopher working logo" width="160" height="160">
  </a>

<h3 align="center">🏗️ Golang Architectures</h3>

  <p align="center">
    Exploring different types of architectures using Golang. This repo is for everyone who wants to build a solid project, to learn about software architecture patterns or to improve already existing skills. Note that the following patterns can be easly applied using any other programming language as they are agnostic, so see the language as the example implementation, but focus more on what the architecture can do for you or your project!
  </p>

  <div>
    <p align="center">
      <a href="https://github.com/ksrof/golang-architectures">
        <img src="https://img.shields.io/badge/Made%20with-Go-1f425f.svg"/>
      </a>
      <a href="https://github.com/ksrof/golang-architectures">
        <img src="https://img.shields.io/github/go-mod/go-version/ksrof/golang-architectures?filename=layered-gin%2Fgo.mod"/>
      </a>
      <a href="https://github.com/ksrof/golang-architectures/LICENSE">
        <img src="https://img.shields.io/badge/license-MIT-blue.svg"/>
      </a>
      <a href="https://github.com/ksrof/golang-architectures">
        <img src="https://badgen.net/badge/Open%20Source%20%3F/Yes%21/blue?icon=github"/>
      </a>
    </p>
  </div>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#why-do-i-need-a-clean-architecture">Why Do I Need A Clean Architecture?</a>
      <a href="#built-with">Built With</a>
      <a href="#layered-architecture">Layered Architecture</a>
      <a href="#contributing">Contributing</a>
      <a href="#license">License</a>
      <a href="#contact">Contact</a>
    </li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->

## Why Do I Need A Clean Architecture

This repo is a reference for the Go community that aims to help developers choose the right architecture for their project. Doesn't matter if you are applying the following concepts to a personal project or as part of a larger team, having the know-how about which software architecture suits better your needs is an important skill to have. Establishing good patterns which are consistent and well-known can reduce the time a developer has to waste undestanding their own (or others) code.

<p align="right">(<a href="#top">back to top</a>)</p>

## Built With

- [Golang](https://golang.org)
- [Gin](https://github.com/gin/gin-gonic)
- [Gorm](https://gorm.io)

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- ARCHITECTURES -->

## Layered Architecture

In this pattern each layer plays a distinct role within the application, the request needs to pass through the layer down below in order to get to the next layer. Layers are only responsible for what happens inside them, so you can modify one layer without messing the other.

Let's illustrate this with an example of how it would handle a request

![Layered Architecture](docs/images/layered-architecture.svg)

**W.I.P**

### User Request

### Delivery Layer

### Use-Case Layer

### Repository Layer

<!-- CONTRIBUTING -->

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- LICENSE -->

## License

Distributed under the MIT License. See `LICENSE` for more information.

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- CONTACT -->

## Contact

Kevin Suñer Ramos - [ksrof](https://github.com/ksrof) - kevinsunercontacto@gmail.com

Project Link: [https://github.com/ksrof/yummybook-api](https://github.com/ksrof/yummybook-api)

<p align="right">(<a href="#top">back to top</a>)</p>
