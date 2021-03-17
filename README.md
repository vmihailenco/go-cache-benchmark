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
  tinylfu   |     0.34 |  603ms | 0.21MiB | 338748 | 661252
  clockpro  |     0.33 |  630ms | 0.40MiB | 326931 | 673069
  arc       |     0.32 | 1331ms | 0.49MiB | 320073 | 679927
  ristretto |     0.32 | 1193ms | 4.16MiB | 318215 | 681785
  two-queue |     0.32 | 1205ms | 0.33MiB | 317993 | 682007
  lru       |     0.28 |  838ms | 0.20MiB | 280385 | 719615


zipfian size=1000 samples=100000

    CACHE   | HIT RATE |  TIME  | MEMORY  |  HITS  | MISSES
------------+----------+--------+---------+--------+---------
  tinylfu   |     0.29 |  609ms | 0.18MiB | 289853 | 710147
  clockpro  |     0.29 |  668ms | 0.41MiB | 285895 | 714105
  arc       |     0.28 | 1337ms | 0.47MiB | 283090 | 716910
  ristretto |     0.24 | 1166ms | 4.14MiB | 237856 | 762144
  two-queue |     0.28 | 1348ms | 0.34MiB | 276322 | 723678
  lru       |     0.21 |  860ms | 0.22MiB | 207349 | 792651


zipfian size=10000 samples=100000

    CACHE   | HIT RATE |  TIME  | MEMORY  |  HITS  | MISSES
------------+----------+--------+---------+--------+---------
  tinylfu   |     0.41 |  671ms | 1.99MiB | 411619 | 588381
  clockpro  |     0.41 |  686ms | 3.53MiB | 407354 | 592646
  arc       |     0.40 | 1448ms | 4.18MiB | 395339 | 604661
  ristretto |     0.40 | 1051ms | 7.08MiB | 398055 | 601945
  two-queue |     0.40 | 1379ms | 3.12MiB | 398903 | 601097
  lru       |     0.37 |  972ms | 1.98MiB | 365356 | 634644


zipfian size=10000 samples=1000000

    CACHE   | HIT RATE |  TIME  | MEMORY  |  HITS  | MISSES
------------+----------+--------+---------+--------+---------
  tinylfu   |     0.37 |  701ms | 1.97MiB | 365988 | 634012
  clockpro  |     0.36 |  708ms | 3.59MiB | 362778 | 637222
  arc       |     0.36 | 1424ms | 4.17MiB | 362951 | 637049
  ristretto |     0.34 | 1062ms | 7.28MiB | 338210 | 661790
  two-queue |     0.36 | 1443ms | 3.20MiB | 358252 | 641748
  lru       |     0.30 | 1028ms | 1.98MiB | 303715 | 696285


zipfian size=100000 samples=1000000

    CACHE   | HIT RATE |  TIME  |  MEMORY  |  HITS  | MISSES
------------+----------+--------+----------+--------+---------
  tinylfu   |     0.46 |  742ms | 19.18MiB | 459283 | 540717
  clockpro  |     0.46 |  627ms | 29.57MiB | 458122 | 541878
  arc       |     0.46 | 1445ms | 36.00MiB | 460501 | 539499
  ristretto |     0.45 | 1177ms | 29.34MiB | 453806 | 546194
  two-queue |     0.46 | 1299ms | 28.79MiB | 460002 | 539998
  lru       |     0.45 |  796ms | 17.76MiB | 445184 | 554816
```

To run this benchmark:

```shell
go run *.go
```
