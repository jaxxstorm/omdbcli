# omdbcli

A cli tool to search and retrieve movie details from the OMDb API

# Usage

## Search

Simply specify a search string! Example:

```
./omdbcli search friday
+--------------------------------+-----------+-----------+--------+
|             TITLE              |   YEAR    |  IMDB ID  |  TYPE  |
+--------------------------------+-----------+-----------+--------+
| Freaky Friday                  |      2003 | tt0322330 | movie  |
| Friday the 13th                |      1980 | tt0080761 | movie  |
| Friday                         |      1995 | tt0113118 | movie  |
| Friday the 13th                |      2009 | tt0758746 | movie  |
| Friday Night Lights            |      2004 | tt0390022 | movie  |
| Friday Night Lights            | 2006–2011 | tt0758745 | series |
| His Girl Friday                |      1940 | tt0032599 | movie  |
| Friday the 13th Part 2         |      1981 | tt0082418 | movie  |
| Friday the 13th Part III       |      1982 | tt0083972 | movie  |
| Friday the 13th: The Final     |      1984 | tt0087298 | movie  |
| Chapter                        |           |           |        |
+--------------------------------+-----------+-----------+--------+
```

You can narrow down your search, by year:

```
./omdbcli search friday -y 1984
+--------------------------------+------+-----------+-------+
|             TITLE              | YEAR |  IMDB ID  | TYPE  |
+--------------------------------+------+-----------+-------+
| Friday the 13th: The Final     | 1984 | tt0087298 | movie |
| Chapter                        |      |           |       |
+--------------------------------+------+-----------+-------+
```

or by type:

```
./omdbcli search friday -t series 
+--------------------------------+-----------+-----------+--------+
|             TITLE              |   YEAR    |  IMDB ID  |  TYPE  |
+--------------------------------+-----------+-----------+--------+
| Friday Night Lights            | 2006–2011 | tt0758745 | series |
| Friday Night Dinner            | 2011–     | tt1844923 | series |
| Friday Night with Jonathan     | 2001–2010 | tt0299302 | series |
| Ross                           |           |           |        |
| The Friday Night Project       | 2005–     | tt0445876 | series |
| Friday Night Tykes             | 2014–     | tt3467018 | series |
| TFI Friday                     | 1996–2015 | tt0115383 | series |
| Oye! It's Friday!              | 2008–     | tt1373021 | series |
| Friday: The Animated Series    | 2007–     | tt1077170 | series |
| Friday Download                | 2011–     | tt1949706 | series |
| Friday Night                   | 1983–2000 | tt0165583 | series |
+--------------------------------+-----------+-----------+--------+
```

## Retrieve Details

You can retrieve details on an entry using the get command:

```
./omdbcli get friday night lights
+---------------------+------+------------+--------------------------------+--------------------------------+--------------------------------+-------+
|        TITLE        | YEAR |  DIRECTOR  |             WRITER             |             ACTORS             |              PLOT              | TYPE  |
+---------------------+------+------------+--------------------------------+--------------------------------+--------------------------------+-------+
| Friday Night Lights | 2004 | Peter Berg | Buzz Bissinger (book), David   | Billy Bob Thornton, Lucas      | Odessa, Texas, is a small,     | movie |
|                     |      |            | Aaron Cohen (screenplay),      | Black, Garrett Hedlund, Derek  | town in Texas. Racially        |       |
|                     |      |            | Peter Berg (screenplay)        | Luke                           | divided and economically       |       |
|                     |      |            |                                |                                | dying, there is one night that |       |
|                     |      |            |                                |                                | gives the town something to    |       |
|                     |      |            |                                |                                | live for: Friday Night. The    |       |
|                     |      |            |                                |                                | Permian Panthers have a big    |       |
|                     |      |            |                                |                                | winning tradition in Texas     |       |
|                     |      |            |                                |                                | high school football, led by   |       |
|                     |      |            |                                |                                | QB Mike Winchell and superstar |       |
|                     |      |            |                                |                                | tailback Boobie Miles, but all |       |
|                     |      |            |                                |                                | is not well, as Boobie suffers |       |
|                     |      |            |                                |                                | a career-ending injury in      |       |
|                     |      |            |                                |                                | the first game of the season.  |       |
|                     |      |            |                                |                                | Hope is lost among citizens    |       |
|                     |      |            |                                |                                | in Odessa, and for the team,   |       |
|                     |      |            |                                |                                | but Coach Gary Gaines, who     |       |
|                     |      |            |                                |                                | believes that "Perfection      |       |
|                     |      |            |                                |                                | is being able to look your     |       |
|                     |      |            |                                |                                | friends in the eye and know    |       |
|                     |      |            |                                |                                | you did everything you could   |       |
|                     |      |            |                                |                                | not to let them down", is      |       |
|                     |      |            |                                |                                | somehow able to help the team  |       |
|                     |      |            |                                |                                | rise up from the ashes and     |       |
|                     |      |            |                                |                                | make a huge season comeback.   |       |
|                     |      |            |                                |                                | Now on their way to state,     |       |
|                     |      |            |                                |                                | the Panthers must go out and   |       |
|                     |      |            |                                |                                | be perfect, because they may   |       |
|                     |      |            |                                |                                | never matter this much for the |       |
|                     |      |            |                                |                                | rest of their lives.           |       |
+---------------------+------+------------+--------------------------------+--------------------------------+--------------------------------+-------+
```

Optionally, you can also specify the retrieval of the rotten tomatoes rating. Not all entries have one:

```
./omdbcli get thor -r
+-------+------+-----------------+--------------------------------+--------------------------------+--------------------------------+-------+
| TITLE | YEAR |    DIRECTOR     |             WRITER             |             ACTORS             |              PLOT              | TYPE  |
+-------+------+-----------------+--------------------------------+--------------------------------+--------------------------------+-------+
| Thor  | 2011 | Kenneth Branagh | Ashley Miller (screenplay),    | Chris Hemsworth, Natalie       | The warrior Thor (Hemsworth)   | movie |
|       |      |                 | Zack Stentz (screenplay), Don  | Portman, Tom Hiddleston,       | is cast out of the fantastic   |       |
|       |      |                 | Payne (screenplay), J. Michael | Anthony Hopkins                | realm of Asgard by his         |       |
|       |      |                 | Straczynski (story), Mark      |                                | father Odin (Hopkins) for his  |       |
|       |      |                 | Protosevich (story), Stan Lee  |                                | arrogance and sent to Earth    |       |
|       |      |                 | (comic book), Larry Lieber     |                                | to live among humans. Falling  |       |
|       |      |                 | (comic book), Jack Kirby       |                                | in love with scientist Jane    |       |
|       |      |                 | (comic book)                   |                                | Foster (Portman) teaches Thor  |       |
|       |      |                 |                                |                                | much-needed lessons, and his   |       |
|       |      |                 |                                |                                | new-found strength comes into  |       |
|       |      |                 |                                |                                | play as a villain from his     |       |
|       |      |                 |                                |                                | homeland sends dark forces     |       |
|       |      |                 |                                |                                | toward Earth.                  |       |
+-------+------+-----------------+--------------------------------+--------------------------------+--------------------------------+-------+
|                                                                                                        ROTTEN TOMATOES RATING     |  77%  |
+-------+------+-----------------+--------------------------------+--------------------------------+--------------------------------+-------+
```

# Building

The app uses [glide](http://glide.sh/) for its dependencies, so ensure you have it installed and it's in your `$PATH`.

From there, run `glide install` to download all the apps dependencies

Once they're installed into your `vendor` directory, you can build it as normal using `go build`:

```
go build -o omdbcli main.go
```
