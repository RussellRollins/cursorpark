# cursorpark

It is important for the cursors of those who are lurking on your Google Doc to feel safe and secure.

Give them a home for rest and relaxation in a cursor's natural space.

```
Cursor Park & Nature Preserve
🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳
🌳      🐌      🐐                        🌳
🌳                    🌴🐕                🌳
🌳                🐛                      🌳
🌳                                        🌳
🌳                        🐝              🌳
🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳
```

Note: Because of Unicode/Fonts being silly, you may need to add some spacing to your nature preserve. It seems to work correctly in Nerd Font Complete in my terminal, but other places not so much.

## Build cursorpark

Requires at least Go 1.10, tested on 1.12.1.

```
go build
```

## Flags

Generate some cursor parks with special properties with flags:

This will generate a very big park with lots of animals. Be careful cursors!
```
cursorpark --width=50 --height=12 --freq=0.30
```
