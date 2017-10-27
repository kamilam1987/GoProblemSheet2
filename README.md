# GoProblemSheet2

# What is go (programming language)?
Go (often referred to as golang) is a programming language created at Google in 2007 by Robert Griesemer, Rob Pike, and Ken Thompson. It is a compiled, statically typed language in the tradition of Algol and C, with garbage collection, limited structural typing, memory safety features and CSP-style concurrent programming features added. The compiler and other language tools originally developed by Google are all free and open source.
# How to install go
Please visit this link: https://golang.org/doc/install?download=go1.5.windows-amd64.msi2
# How to clone a repository and run programs
Please visit links: https://help.github.com/articles/cloning-a-repository/

https://github.com/processing/processing/wiki/Build-Instructions
# This repository contains problem solving :
## 1. Guessing game
Create a web application in Go that responds with the text “Guessing game”. This should be the response body irrespective of what request is received. Explain in your README how to examine the response, including the headers, using curl.

## 2. Make the text a H1
Change your web application to make “Guessing game” a level 1 heading in HTML. Test your application, making sure that the HTML is rendered by your browser. If it isn’t, fix it.

## 3. Serve a page using Bootstrap
Change the web application to serve a web page rather than hard-coding the text into the web application executable. Use the Bootstrap starter template, changing the header to say “Guessing game”. Add a link on the page to the relative URL /guess with the text “New game”. Have this page served as the root resource in the web application.

## 4. Add a guess route
Create a new route in your application at /guess. Have it serve a new page called guess.html. Use the same Bootstrap code for the page as in index.html but add a level 2 heading with the text “Guess a number between 1 and 20”. Add a form, with a text input with the name “guess” and a submit button with the label “Guess”. The action of the form should be /guess and the method should be GET.

## 5. Turn the guess page into a template
Change the web application to use the guess.tmpl file as a template. Add a single variable to the template called Message and create a struct in Go containing a single string. Create an instance of the struct with the string set to “Guess a number between 1 and 20” and have the template render this in the H2 tag.

## 6. Check for a cookie
In the /guess handler check if the cookie called target is set. If it is not, generate a random number between 1 and 20 and set a cookie called target to that value in the response. Otherwise, leave the cookie at its current value.

## 7. Check for a variable
Have the /guess handler check if a URL encoded variable called guess has been set to an integer. If it has, have the text “You guessed {guess}.” inserted into the template where {guess} is replaced with the value of guess.

## 8.Compare the cookie and the guess
If the target cookie and the guess variable are both set, then have the /guess handler compare them. If they are equal, set the target cookie to another randomly generated integer, and have the template display a congratulations message and a link to create a new game. Otherwise, have the template display a message telling the user what their guess was and whether it was too high or too low.

## 9. Use the POST method
Change your HTML form in guess.html to use the POST method instead. Make sure your application still works, bug fixing it if necessary.

Resources:
https://en.wikipedia.org/wiki/Go_(programming_language)
https://golang.org/doc/install?download=go1.5.windows-amd64.msi2
https://help.github.com/articles/cloning-a-repository/
