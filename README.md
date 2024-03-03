☦️ This software was written in the name of the __Father__ and of the __Son__ and of the __Holy Spirit__; Amen. 

About
===
This repo contains the source code for my personal website. 

There are a couple of goals I had when creating the website:
1. Simple. The website must be simple and elegant. Absolutely NO BLOAT!
2. Fast. I wanted load times to be 0. Everything should render IMMEDIATELY!
3. Extensible. Not necessarily scalable - after all this is just my small static website.
I did want to be able to add new things quickly and seamlessly though.

Seeing as I wanted to use my website as a learning opportunity, and keeping the previous 
goals in mind, I opted to use the following technologies:
1. Go for serving static content using it's HTML template library. This allows me to embed 
structs and basic functions into the HTML template before it's rendered.
2. Tailwind for styling. I've never really enjoyed fiddling with custom CSS classes. 
Tailwind makes my life a bit easier. 
3. HTMX, because Javascript sux XD

Building
=== 

Dependencies
=====

There are a few dependencies required to build the project from source:
* Go
* Tailwind
* Tailwind animation delay

Running
=====

```console
$ ./run.sh
```
