# Snowflake

Snowflake is a dependency free implementation of the Twitter Snowflake ID generator in Go. Twitter Snowflake is a 
unique ID generator that is distributed and scalable. See: https://en.wikipedia.org/wiki/Snowflake_ID

It is used to generate unique IDs for distributed systems. The ID is a 64-bit integer that is composed of a timestamp, 
a machine ID, and a sequence number. The timestamp is the number of milliseconds since a custom epoch. The machine ID 
is a unique number that identifies the machine that generates the ID. The sequence number is a number that is 
incremented when multiple IDs are generated within the same millisecond.

A 64 bits ID can be represented as an 11 character string in base64 encoding. This is a very efficient way to store the 
ID in a string format in XML, JSON, or a database, although I would store it as a binary uint64 if possible.

This package can be a livesaver if you need to generate unique IDs in a distributed system, otherwise you would have to
use a database or service to generate unique IDs, which would introduce latency, a bottleneck and a single point 
of failure.

All the packages I investigated are amazing, and I would not hold back to use any of them if you do not want to use 
this package. I think the Influx implementation is the best alternative, because it is fast. If you cannot use it 
because of the potential out of order ids or the lack of flexibility, then you could bwmarin/snowflake or 
godruoyi/go-snowflake.

## Motivation

I decided to make this implementation because I was not satisfied with the existing implementations, mostly because of
the lack of flexibility, code quality, and performance. I also wanted to have a package that is easy to use and has a 
good test coverage.

I think this package delivers on these points. It is easy to use, has a good test coverage, and is very fast. It is also
very flexible, allowing for different epochs and bit sizes per instance in a single application.

## Comparison


### Dependencies

| Package                           | Dependencies |
|-----------------------------------|--------------|
| github.com/crosscode-nl/snowflake | 0            | 
| github.com/influxdata/snowflake   | ~43 ns/op    | 
| github.com/bwmarrin/snowflake     | ~244 ns/op   | 
| github.com/godruoyi/go-snowflake  | ~244 ns/op   |

A low or zero dependency count is a good thing. It means the package is easy to use and has a low risk of breaking
because of a dependency update. The chance of a security issue (CVE) is lower.

### Performance

| Package                           | Performance      | IDs/s    | Blocking            | Synchronization | Coverage   |
|-----------------------------------|------------------|----------|---------------------|-----------------|------------|
| github.com/crosscode-nl/snowflake | ~43 / ~244 ns/op | 23M / 4M | optional, sleep     | CAS, unlimited  | 100%       |
| github.com/influxdata/snowflake   | ~43 ns/op        | 23M      | no                  | CAS, 100x, bug  | 88.9%      |
| github.com/bwmarrin/snowflake     | ~244 ns/op       | 4M       | yes, busy loop, bug | Mutex           | 92.9%      |
| github.com/godruoyi/go-snowflake  | ~244 ns/op       | 4M       | yes, busy loop      | CAS, unlimited  | 91.4%      |

With 244 ns/op you can generate 4 million IDs per second. With 43 ns/op you can generate 23 million IDs per second per 
instance. 

The biggest performance comes from allowing drift or not. When drift is allowed, the generator can generate IDs 
without blocking. Our implementation supports both blocking and non-blocking modes. The blocking mode will give
similar results to the other blocking implementations.

A thing I noticed is that it does not seem to matter a lot if the implementation uses a naive mutex implementation or a 
'lock-free' CAS loop. Supporting and allowing drift is the most important factor for performance on peak load. If you
do not care that the ids are out of order, you can use the non-blocking mode and have the best performance.

Also, changing the machine ID bits size will change the performance. When it becomes smaller the sequence becomes larger
, we will have less drift, and more IDs can be generated per millisecond in blocking mode. 

For example, changing machine ID bits from 10 to 9 will change the performance in blocking mode. Consider:

| Bits | Performance  | IDs/s |
|------|--------------|-------|
| 6    | ~44 ns/op    | 22M   |
| 7    | ~44 ns/op    | 22M   |
| 9    | ~122 ns/op   | 8M    |
| 10   | ~244 ns/op   | 4M    |
| 11   | ~488 ns/op   | 2M    |
| 16   | ~15613 ns/op | 64K   |

**BUG: The influxdata/snowflake implementation has a bug with the CAS loop, which potentially could cause
a larger drift than necessary. It is very unlikely to experience this bug though.**

**BUG: The bwmarrin/snowflake could potentially cause an extra temporary millisecond when it blocks when the sequence is 
exhausted.**

**NOTE: All benchmarks are done on a 16GB 2020 M1 Mac Mini.**

### Generator features

| Package                           | Drift | Custom Epoch      | Change bit size   | Timestamp | Default epoch        |
|-----------------------------------|-------|-------------------|-------------------|-----------|----------------------|
| github.com/crosscode-nl/snowflake | yes   | yes, per instance | yes, per instance | 42 bits   | 2024-03-01 00:00:00Z |
| github.com/influxdata/snowflake   | yes   | no                | no                | 42 bits   | 2017-04-09 00:00:00Z | 
| github.com/bwmarrin/snowflake     | no    | yes, global       | yes, global       | 41 bits   | 2010-11-04 01:42:54Z | 
| github.com/godruoyi/go-snowflake  | no    | yes, global       | yes, global       | 41 bits   | 2008-11-10 23:00:00Z | 

We allow multiple instances of the generator to have different epochs and bit sizes. This allows for more flexibility
in the use of the generator, particularly in an environment where multiple systems use different snowflake configurations.

The influx implementation is not configurable.

The bwmarrin and godruoyi implementations are configurable via package global variables. This means only one 
configuration is possible for the entire application. It looks like bwmarin made a start to make the configuration
per instance, but it is not finished.

Twitter uses 41 bits for the timestamp, which allows for a maximum of 69 years of IDs since the epoch, but we allow
for 42 bits. This allows for a maximum of 138 years of IDs since the epoch. Influx also uses 42 bits for the timestamp.

Using 42 bits for the timestamp allows for a maximum of 138 years of IDs since the epoch.
Using 41 bits for the timestamp allows for a maximum of 69 years of IDs since the epoch.

Having the possibility to set the epoch to the start of the app initial build date allows a longer period of id then
when using the unix epoch. None of the packages do this, but Influx does not allow to set the epoch. However, its default 
epoch is: 1491696000000, which is 2017-04-09 00:00:00.000 UTC.

**NB: If you switch between these implementations, make sure to set the epoch to the same value of the original package.**

### ID features

| Package                           | ID      | Encoding                                                                     | Default          | Decode     |
|-----------------------------------|---------|------------------------------------------------------------------------------|------------------|------------|
| github.com/crosscode-nl/snowflake | uint64  | Base64(std,url,mime,influx), Influx64(std,url,mime,influx), Hex(Upper,Lower) | Hex(Upper)       | yes        |
| github.com/influxdata/snowflake   | uint64  | Influx64(influx)                                                             | Influx64(influx) | no         |
| github.com/bwmarrin/snowflake     | int64   | Decimal, Base2, Base32, Base36, Base58, Base64                               | Decimal          | deprecated | 
| github.com/godruoyi/go-snowflake  | uint64  | None                                                                         | None             | yes        | 

The encoding features are for convenience only, although our implementations are optimized for speed.

Decode means that the package has a function to decode the ID into a struct with the timestamp, machineID, and sequence. 
bwmarin/snowflake has a deprecated function to decode the ID, but is still works.

The encoding options are nice if you need to convert the ID to a string. The Influx64 encoding is Base64 encoding which
is not compatible with the standard Base64 encoding. It is very fast and has a low memory footprint. I would pick this
if it is possible to use it.

The Base64 and Influx64 encodings deliver the shortest strings (11 bytes). 
The most efficient option is to store the ID as a binary uint64 (8 bytes). 

*TIP: If your system uses strings, and you want to use a different epoch, then you could switch to an encoding 
with a different length if your system can handle larger or shorter ID strings. You could also choose to add padding to
your strings and change the padding character to a different character.*

## License 

This package is licensed under the MIT license. See: [LICENSE](LICENSE)
