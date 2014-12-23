# www

**www** is a program that lets you send the output from one or more commands, in the terminal, to a web browser.

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

**After you run that, open up your web browser and go to (for example) `http://127.0.0.1:5555/` to view it.**

Of course, these arr simple (and rather useless) examples.

The power of the **www** command is when you get command line tools to output something complex.
Such as HTML or images.

For example:
```
 cat log.txt | grep ERROR | awk '{print $4}' | histogram | htmlplot | www
```

Or, also for example:
```
cat DOCS.md | markdown --html4tag | www
```
