# www

**www** is a command line tool that lets you send the output from one or more commands, in the terminal, to a web browser.

## Usage

Very simple usage is:

```
echo 'apple banana cherry' | www
```

or:
```
echo some_photo.jpeg | www
```

or:
```
echo report.html | www
```

or:
```
echo '<html><body><h1>Hello world!</h1></body></html>' | www
```

or (in a shell that support *process substitution*, such as bash and zsh):
```
www <(echo 'apple banana cherry')
```

**After you run that, open up your web browser and go to (for example) `http://127.0.0.1:5555/` to view it.**

Of course, these are simple (and rather useless) examples.

The power of the **www** command is when you want to view the output that results from a complex set of
command line calculations. Especially if the final result *is* or *can be transformed into* HTML or images.

For example:
```
 cat log.txt | grep ERROR | awk '{print $4}' | histogram | htmlplot | www
```

Or, also for example:
```
cat DOCS.md | markdown --html4tag | www
```


You can also override what TCP port the HTTP server uses. You can do that with code like:
```
... | www --port=2323
```
