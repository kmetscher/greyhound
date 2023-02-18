![Gopher Greyhound, see link on Renee French](./logo.png)

Greyhound (noun): Something fast, at least ostensibly, like the dog or the bus line, or drank fast, like the cocktail.

The Greyhound first appeared in the Savoy Cocktail Book. Shake:

- 1 part [gin](https://github.com/gin-gonic/gin)
- 4 parts [grapefruit juice](https://github.com/vuejs/)
- Ice

Strain into a highball. Garnish with a lime or a lemon. Salt the rim for a Salty Dog.

## Motivation

I've extensively used [Laravel](https://github.com/laravel/laravel) alongside [Inertia](https://github.com/inertiajs/inertia) to flexibly build modern web applications with
an expressive and powerful server side and a single-page application client side with React and Vue. Done tastefully, this allows you to engage in anything from progressively enhanced
websites to full-stack apps without the need for writing an API or using Node-ecosystem libraries for data retrieval.

Greyhound is one part Gin adapter for Inertia itself, and aims to be another part complete full stack framework in its own right. Simply implementing the Inertia protocol for Gin already
opens up many doors for rapid development (and I will likely extract the Inertia implementation into its own repository for people to use how they see fit), but adding the other bells and whistles
of a more complete server-side MVC in Go will make it possible to build applications in a lingua franca while still being able to interface natively with the backend services that Go is 
used and loved for. Being able to use concurrency as a primitive while constructing controllers and helper functions, as well as its use within Gin itself, opens up doors for ridiculously
fast applications; combine this with Go and Gin's HTTP libraries, and you have the option to compile a web application into a self-hosting binary that can just be hucked onto a server, or 
deployed intelligibly within cloud-native architectures like clusters and containers.

## Links

- [Blog](https://kylemetscher.com)
- [LinkedIn](https://www.linkedin.com/in/c0w80yd4n)
- [Renee French's gopher](https://go.dev/blog/gopher)
- [IPChicken](https://ipchicken.com)
