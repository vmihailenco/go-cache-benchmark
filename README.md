<p align="center">
  <a href="https://uptrace.dev/?utm_source=gh-redis&utm_campaign=gh-redis-banner1">
    <img src="https://raw.githubusercontent.com/uptrace/roadmap/master/banner1.png" alt="All-in-one tool to optimize performance and monitor errors & logs">
  </a>
</p>

# Cache benchmark for Go

This benchmark compares cache algorithms using scrambled zipfian distribution (a few occur very
often while many others occur rarely). Other distributions are supported too, but they produce the
same results

The following libraries are supported:

- https://github.com/dgryski/go-tinylfu
- https://github.com/dgryski/go-clockpro
- https://github.com/dgraph-io/ristretto
- https://github.com/hashicorp/golang-lru (LRU, ARC, TwoQueue)

And the results are:

```
zipfian size=1000 samples=10000

    CACHE   | HIT RATE |  TIME  | MEMORY  |  HITS  | MISSES
------------+----------+--------+---------+--------+---------
  tinylfu   |     0.34 |  615ms | 0.21MiB | 338072 | 661928
  clockpro  |     0.33 |  661ms | 0.40MiB | 327255 | 672745
  arc       |     0.31 | 1416ms | 0.48MiB | 308497 | 691503
  ristretto |     0.32 | 1168ms | 4.08MiB | 317563 | 682437
  two-queue |     0.32 | 1396ms |       - | 318059 | 681941
  lru       |     0.28 |  902ms | 0.22MiB | 280055 | 719945


zipfian size=1000 samples=100000

    CACHE   | HIT RATE |  TIME  | MEMORY  |  HITS  | MISSES
------------+----------+--------+---------+--------+---------
  tinylfu   |     0.29 |  642ms | 0.19MiB | 288596 | 711404
  clockpro  |     0.29 |  707ms | 0.41MiB | 285814 | 714186
  arc       |     0.28 | 1496ms | 0.46MiB | 283052 | 716948
  ristretto |     0.24 | 1111ms | 4.11MiB | 235774 | 764226
  two-queue |     0.28 | 1565ms |       - | 277095 | 722905
  lru       |     0.21 | 1022ms | 0.20MiB | 207190 | 792810


zipfian size=10000 samples=100000

    CACHE   | HIT RATE |  TIME  | MEMORY  |  HITS  | MISSES
------------+----------+--------+---------+--------+---------
  tinylfu   |     0.41 |  660ms | 1.98MiB | 412308 | 587692
  clockpro  |     0.41 |  685ms | 3.52MiB | 406986 | 593014
  arc       |     0.40 | 1363ms | 4.29MiB | 398631 | 601369
  ristretto |     0.40 | 1041ms | 7.11MiB | 398656 | 601344
  two-queue |     0.40 | 1134ms |       - | 398116 | 601884
  lru       |     0.37 |  767ms | 1.98MiB | 365826 | 634174


zipfian size=10000 samples=1000000

    CACHE   | HIT RATE |  TIME  | MEMORY  |  HITS  | MISSES
------------+----------+--------+---------+--------+---------
  tinylfu   |     0.37 |  568ms | 1.97MiB | 366694 | 633306
  clockpro  |     0.36 |  590ms | 3.60MiB | 362823 | 637177
  arc       |     0.36 | 1253ms | 4.19MiB | 363335 | 636665
  ristretto |     0.34 | 1049ms | 7.66MiB | 338908 | 661092
  two-queue |     0.36 | 1143ms |       - | 358100 | 641900
  lru       |     0.30 |  726ms | 1.96MiB | 303493 | 696507


zipfian size=100000 samples=1000000

    CACHE   | HIT RATE |  TIME  |  MEMORY  |  HITS  | MISSES
------------+----------+--------+----------+--------+---------
  tinylfu   |     0.46 |  751ms | 19.11MiB | 459360 | 540640
  clockpro  |     0.46 |  627ms | 29.60MiB | 459233 | 540767
  arc       |     0.46 | 1425ms | 35.95MiB | 460739 | 539261
  ristretto |     0.45 | 1127ms | 29.74MiB | 454384 | 545616
  two-queue |     0.46 | 1329ms | 25.30MiB | 459803 | 540197
  lru       |     0.44 |  810ms | 17.75MiB | 444974 | 555026
```

To run this benchmark:

```shell
go run *.go
```
