# Simple command line Github Search

`ghs` is a simple command line tool which will open the corresponding url for your github search in your default web browser.

## INSTALLATION
```
go get github.com/ocakhasan/ghs
```

it will create a binary file in your `GOBIN` folder.

## USAGE
```
ghs <query> <query-type>
```

For example if I want to search users with name `hasan`, what I should do is 

```
ghs hasan u
```

If I want to make a general search

```
ghs "repo-name"
```

it will make a general search in Github.

## Query Types

| Flag    | Query Type |
| ------- | ---------- |
| c       | commits    |
| commits | commits    |
| u       | users      |
| users   | users      |
| w       | wikis      |
| wikis   | wikis      |
| code    | code       |
| i       | issues     |
| issues  | issues     |
| t       | topics     |
| topics  | topics     |
