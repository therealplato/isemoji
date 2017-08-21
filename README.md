## isemoji
reads stdin until EOF, then exits 0 if emoji, exits 1 if no emoji

Output hinting and copies of any input lines containing emoji are output to stderr.

Emoji character codes are copied from unicode consortium documentation, version 6.0 (2017).
The following codes are present in the docs, but have been intentionally dropped from this tool because they aren't really emoji:
```
0023          ; Emoji                #  1.1  [1] (#️)       number sign
002A          ; Emoji                #  1.1  [1] (*️)       asterisk
0030..0039    ; Emoji                #  1.1 [10] (0️..9️)    digit zero..digit nine
```

## install
`go get -u github.com/therealplato/isemoji`
verify your path contains `$GOPATH/bin` 

## use
`echo "select name from users" | [sql](https://github.com/marianogappa/sql) production | isemoji`
