# BadUrlShortener

a very bad url shortener

## Why is this bad?

cuz very bad code

doesnt have mutexes and stores the urls in a poorly formatted text file on the server's hardware (local storage)

has the ability to add new shortened urls by using a get request but the problem is that its very unsafe and may not always sync due to the lack of mutexes

## Can I use this?

if u want then sure
