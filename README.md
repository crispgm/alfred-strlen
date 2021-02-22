# Alfred Strlen

Show string length with Alfred.

<!-- ![screenshot](/screenshot.png) -->

## Installation

### Install via Packal

To be released.
<!-- [Alfred Strlen](http://www.packal.org/workflow/alfred-strlen) -->

### Install manually

Download and import [alfred-strlen.alfredworkflow](https://github.com/crispgm/alfred-strlen/raw/master/alfred-strlen.alfredworkflow).

## Dependencies

- [Alfred](https://www.alfredapp.com/)
- [Alfred Powerpack](https://www.alfredapp.com/powerpack/)

## Usage

Show string length:

```shell
len hello, world

Output:
hello, world
12
```

Show number of characters (runes):

```shell
# ASCII characters
count hello, wolrd

Output:
hello, world
12

# UTF-8 characters
count 你好, 世界

Output:
你好, 世界
14
```

## Credit

- [deanishe/awgo](https://github.com/deanishe/awgo)

## License

MIT
