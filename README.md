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

<details>
  <summary>TLDR</summary>

- TinyLFU works best for small number of keys (~ 100k).
- ClockPro has significantly smaller memory usage with large number of keys (~ 1m).
</details>

```
zipfian cache=1000 keys=2000

    CACHE   | HIT RATE | TIME  | MEMORY  |  HITS  | MISSES
------------+----------+-------+---------+--------+---------
  tinylfu   |     0.62 | 493ms | 0.18MiB | 621277 | 378723
  clockpro  |     0.61 | 465ms | 0.24MiB | 606992 | 393008
  arc       |     0.61 | 989ms | 0.35MiB | 605108 | 394892
  ristretto |     0.62 | 962ms | 0.61MiB | 623534 | 376466
  two-queue |     0.60 | 862ms | 0.29MiB | 601954 | 398046
  lru       |     0.60 | 671ms | 0.19MiB | 600911 | 399089


zipfian cache=1000 keys=10000

    CACHE   | HIT RATE |  TIME  | MEMORY  |  HITS  | MISSES
------------+----------+--------+---------+--------+---------
  tinylfu   |     0.34 |  591ms | 0.19MiB | 337356 | 662644
  clockpro  |     0.33 |  661ms | 0.39MiB | 327584 | 672416
  arc       |     0.32 | 1286ms | 0.49MiB | 320490 | 679510
  ristretto |     0.32 | 1201ms | 4.01MiB | 317564 | 682436
  two-queue |     0.32 | 1290ms | 0.33MiB | 317449 | 682551
  lru       |     0.28 |  792ms | 0.22MiB | 280076 | 719924


zipfian cache=1000 keys=100000

    CACHE   | HIT RATE |  TIME  | MEMORY  |  HITS  | MISSES
------------+----------+--------+---------+--------+---------
  tinylfu   |     0.29 |  579ms | 0.19MiB | 288741 | 711259
  clockpro  |     0.29 |  667ms | 0.41MiB | 286495 | 713505
  arc       |     0.28 | 1361ms | 0.47MiB | 283529 | 716471
  ristretto |     0.24 | 1216ms | 4.13MiB | 236722 | 763278
  two-queue |     0.28 | 1308ms | 0.33MiB | 276821 | 723179
  lru       |     0.21 |  861ms | 0.18MiB | 207555 | 792445


zipfian cache=10000 keys=20000

    CACHE   | HIT RATE |  TIME  | MEMORY  |  HITS  | MISSES
------------+----------+--------+---------+--------+---------
  tinylfu   |     0.66 |  602ms | 1.92MiB | 658056 | 341944
  clockpro  |     0.65 |  572ms | 2.50MiB | 646465 | 353535
  arc       |     0.65 | 1138ms | 3.29MiB | 645151 | 354849
  ristretto |     0.66 |  971ms | 4.33MiB | 658222 | 341778
  two-queue |     0.64 | 1037ms | 2.78MiB | 644071 | 355929
  lru       |     0.64 |  794ms | 1.83MiB | 642418 | 357582


zipfian cache=10000 keys=100000

    CACHE   | HIT RATE |  TIME  | MEMORY  |  HITS  | MISSES
------------+----------+--------+---------+--------+---------
  tinylfu   |     0.41 |  690ms | 1.99MiB | 412261 | 587739
  clockpro  |     0.41 |  690ms | 3.53MiB | 407371 | 592629
  arc       |     0.40 | 1452ms | 4.13MiB | 395852 | 604148
  ristretto |     0.40 | 1056ms | 7.09MiB | 398740 | 601260
  two-queue |     0.40 | 1384ms | 3.08MiB | 398381 | 601619
  lru       |     0.37 |  986ms | 1.98MiB | 365012 | 634988


zipfian cache=10000 keys=1000000

    CACHE   | HIT RATE |  TIME  | MEMORY  |  HITS  | MISSES
------------+----------+--------+---------+--------+---------
  tinylfu   |     0.37 |  707ms | 1.97MiB | 367068 | 632932
  clockpro  |     0.36 |  722ms | 3.60MiB | 362857 | 637143
  arc       |     0.36 | 1409ms | 4.20MiB | 363637 | 636363
  ristretto |     0.34 | 1155ms | 7.04MiB | 338984 | 661016
  two-queue |     0.36 | 1423ms | 3.22MiB | 358408 | 641592
  lru       |     0.30 | 1055ms | 1.95MiB | 302911 | 697089


zipfian cache=100000 keys=200000

    CACHE   | HIT RATE |  TIME  |  MEMORY  |  HITS  | MISSES
------------+----------+--------+----------+--------+---------
  tinylfu   |     0.67 |  713ms | 18.82MiB | 668787 | 331213
  clockpro  |     0.67 |  597ms | 21.77MiB | 666420 | 333580
  arc       |     0.66 | 1220ms | 30.16MiB | 664228 | 335772
  ristretto |     0.67 | 1069ms | 27.78MiB | 667192 | 332808
  two-queue |     0.66 | 1007ms | 25.88MiB | 664851 | 335149
  lru       |     0.66 |  747ms | 17.02MiB | 663942 | 336058


zipfian cache=100000 keys=1000000

    CACHE   | HIT RATE |  TIME  |  MEMORY  |  HITS  | MISSES
------------+----------+--------+----------+--------+---------
  tinylfu   |     0.46 |  708ms | 19.16MiB | 459512 | 540488
  clockpro  |     0.46 |  605ms | 29.60MiB | 458900 | 541100
  arc       |     0.46 | 1407ms | 36.02MiB | 461843 | 538157
  ristretto |     0.45 | 1237ms | 29.53MiB | 454498 | 545502
  two-queue |     0.46 | 1389ms | 28.78MiB | 460941 | 539059
  lru       |     0.45 |  794ms | 17.76MiB | 445593 | 554407


zipfian cache=100000 keys=10000000

    CACHE   | HIT RATE |  TIME  |  MEMORY  |  HITS  | MISSES
------------+----------+--------+----------+--------+---------
  tinylfu   |     0.41 |  743ms | 18.75MiB | 410111 | 589889
  clockpro  |     0.41 |  638ms | 29.71MiB | 411372 | 588628
  arc       |     0.41 | 1436ms | 35.12MiB | 413655 | 586345
  ristretto |     0.41 | 1255ms | 29.95MiB | 405744 | 594256
  two-queue |     0.41 | 1405ms | 27.73MiB | 411990 | 588010
  lru       |     0.40 |  801ms | 17.90MiB | 395560 | 604440


zipfian cache=1000000 keys=2000000

    CACHE   | HIT RATE |  TIME  |  MEMORY   |  HITS  | MISSES
------------+----------+--------+-----------+--------+---------
  tinylfu   |     0.51 |  776ms | 131.43MiB | 507544 | 492456
  clockpro  |     0.51 |  641ms |  78.38MiB | 506999 | 493001
  arc       |     0.51 | 1011ms |  78.16MiB | 506886 | 493114
  ristretto |     0.51 | 1233ms | 198.00MiB | 505005 | 494995
  two-queue |     0.51 |  936ms |  78.19MiB | 506699 | 493301
  lru       |     0.51 |  890ms |  85.85MiB | 507559 | 492441


zipfian cache=1000000 keys=10000000

    CACHE   | HIT RATE |  TIME  |  MEMORY   |  HITS  | MISSES
------------+----------+--------+-----------+--------+---------
  tinylfu   |     0.45 |  797ms | 138.26MiB | 447974 | 552026
  clockpro  |     0.45 |  636ms |  84.37MiB | 446783 | 553217
  arc       |     0.45 | 1094ms |  98.32MiB | 448545 | 551455
  ristretto |     0.45 | 1238ms | 153.82MiB | 446663 | 553337
  two-queue |     0.45 | 1055ms |  98.34MiB | 448392 | 551608
  lru       |     0.45 |  895ms |  92.72MiB | 447545 | 552455
```

To run this benchmark:

```shell
go run *.go
```
