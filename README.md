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
  tinylfu   |     0.34 |  617ms | 0.21MiB | 337948 | 662052
  clockpro  |     0.33 |  644ms | 0.39MiB | 328059 | 671941
  arc       |     0.31 | 1401ms | 0.48MiB | 308081 | 691919
  ristretto |     0.32 | 1151ms | 4.16MiB | 317917 | 682083
  two-queue |     0.32 | 1364ms |       - | 317408 | 682592
  lru       |     0.28 |  877ms | 0.22MiB | 280457 | 719543


zipfian size=1000 samples=100000

    CACHE   | HIT RATE |  TIME  | MEMORY  |  HITS  | MISSES
------------+----------+--------+---------+--------+---------
  tinylfu   |     0.29 |  632ms | 0.19MiB | 289045 | 710955
  clockpro  |     0.29 |  731ms | 0.41MiB | 287093 | 712907
  arc       |     0.28 | 1447ms | 0.46MiB | 283012 | 716988
  ristretto |     0.24 | 1155ms | 4.21MiB | 236766 | 763234
  two-queue |     0.28 | 1536ms |       - | 275998 | 724002
  lru       |     0.21 | 1067ms | 0.22MiB | 207481 | 792519


zipfian size=10000 samples=100000

    CACHE   | HIT RATE |  TIME  | MEMORY  |  HITS  | MISSES
------------+----------+--------+---------+--------+---------
  tinylfu   |     0.41 |  645ms | 1.99MiB | 412113 | 587887
  clockpro  |     0.41 |  708ms | 3.53MiB | 406873 | 593127
  arc       |     0.40 | 1417ms | 4.16MiB | 395633 | 604367
  ristretto |     0.40 | 1079ms | 7.08MiB | 398586 | 601414
  two-queue |     0.40 | 1159ms |       - | 399329 | 600671
  lru       |     0.36 |  808ms | 1.98MiB | 364253 | 635747


zipfian size=10000 samples=1000000

    CACHE   | HIT RATE |  TIME  | MEMORY  |  HITS  | MISSES
------------+----------+--------+---------+--------+---------
  tinylfu   |     0.37 |  573ms | 1.97MiB | 367362 | 632638
  clockpro  |     0.36 |  630ms | 3.59MiB | 363366 | 636634
  arc       |     0.36 | 1236ms | 4.20MiB | 362079 | 637921
  ristretto |     0.34 | 1086ms | 7.15MiB | 339500 | 660500
  two-queue |     0.36 | 1117ms |       - | 357873 | 642127
  lru       |     0.30 |  753ms | 1.95MiB | 304063 | 695937


zipfian size=100000 samples=1000000

    CACHE   | HIT RATE |  TIME  |  MEMORY  |  HITS  | MISSES
------------+----------+--------+----------+--------+---------
  tinylfu   |     0.46 |  728ms | 19.04MiB | 457675 | 542325
  clockpro  |     0.46 |  606ms | 29.57MiB | 459292 | 540708
  arc       |     0.46 | 1461ms | 35.95MiB | 459788 | 540212
  ristretto |     0.45 | 1179ms | 29.32MiB | 454081 | 545919
  two-queue |     0.46 | 1349ms | 25.34MiB | 460084 | 539916
  lru       |     0.45 |  766ms | 17.75MiB | 445461 | 554539
```

To run this benchmark:

```shell
go run *.go
```
