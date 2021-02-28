<h1 align="center">qeg: Quick Estimate Generator</h1>
<p align="center"><i>Made with :heart: by <a href="https://github.com/GreatGodApollo">@GreatGodApollo</a></i></p>

Generate a text based estimate given a JSON representation

## Built With
- [attotto/clipboard](https://github.com/atotto/clipboard)
- [jawher/mow.cli](https://github.com/jawher/mow.cli)
- [mattn/go-runewidth](github.com/mattn/go-runewidth)

## Usage
```
Usage: qeg [-c] [-d] FILE

Quick Estimate Generator

Arguments:
  FILE            What file should it read from?

Options:
  -c, --copy      Copy to clipboard?
  -d, --discord   Surround in codeblock?
```

Example file:
```
{
  "title": "1001",
  "customer": "TESTING!",
  "items": [
    {
      "description": "Moderation Bot",
      "price": 10.5
    }
  ]
}
```

## Licensing

This project is licensed under the [MIT License](https://choosealicense.com/licenses/mit/)

## Authors

* [Brett Bender](https://github.com/GreatGodApollo)