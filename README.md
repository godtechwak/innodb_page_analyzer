Analysis of InnoDB Page

How to use
===============
```
$ go mod init go_page_info
$ go mod tidy
$ go build ./main.go
$ ./main <innodb table file name>.ibd <ENG(default) or KR(=1)>
ex-ENG) ./main test.ibd
ex-KR) ./main test.ibd 1 
```

Examples
===============
##### ENG Version
```
$ ./main test.ibd 1

- InnoDB tablespace file: test0.ibd
- Pages: 43

page_type = File space header(0x0008)                page_num = 0        page_prev = 80300    page_next = 1        record_count = 0        tree_level = 0        index_id = 18446744069414649855
page_type = Insert buffer bitmap(0x0005)             page_num = 1        page_prev = 0        page_next = 0        record_count = 0        tree_level = 0        index_id = 0
page_type = Index node(0x0003)                       page_num = 2        page_prev = 0        page_next = 0        record_count = 0        tree_level = 0        index_id = 18446744069414649855
page_type = Tablespace SDI Index page(0x45bd)        page_num = 3        page_prev = -1       page_next = -1       record_count = 2        tree_level = 0        index_id = 18446744073709551615
page_type = B-tree node(0x45bf)                      page_num = 4        page_prev = -1       page_next = -1       record_count = 13       tree_level = 1        index_id = 2719
page_type = B-tree node(0x45bf)                      page_num = 5        page_prev = -1       page_next = -1       record_count = 8        tree_level = 1        index_id = 2720
page_type = B-tree node(0x45bf)                      page_num = 6        page_prev = -1       page_next = -1       record_count = 4        tree_level = 1        index_id = 2721
page_type = B-tree node(0x45bf)                      page_num = 7        page_prev = -1       page_next = -1       record_count = 8        tree_level = 1        index_id = 2722
page_type = B-tree node(0x45bf)                      page_num = 8        page_prev = -1       page_next = 10       record_count = 105      tree_level = 0        index_id = 2719
page_type = B-tree node(0x45bf)                      page_num = 9        page_prev = 10       page_next = 19       record_count = 210      tree_level = 0        index_id = 2719
page_type = B-tree node(0x45bf)                      page_num = 10       page_prev = 8        page_next = 9        record_count = 208      tree_level = 0        index_id = 2719
page_type = B-tree node(0x45bf)                      page_num = 11       page_prev = 19       page_next = 18       record_count = 3        tree_level = 0        index_id = 2719
page_type = B-tree node(0x45bf)                      page_num = 12       page_prev = -1       page_next = 34       record_count = 285      tree_level = 0        index_id = 2720
page_type = B-tree node(0x45bf)                      page_num = 13       page_prev = 35       page_next = 32       record_count = 312      tree_level = 0        index_id = 2720
page_type = B-tree node(0x45bf)                      page_num = 14       page_prev = -1       page_next = 37       record_count = 301      tree_level = 0        index_id = 2722
page_type = B-tree node(0x45bf)                      page_num = 15       page_prev = 39       page_next = 38       record_count = 298      tree_level = 0        index_id = 2722
page_type = B-tree node(0x45bf)                      page_num = 16       page_prev = -1       page_next = 28       record_count = 555      tree_level = 0        index_id = 2721
page_type = B-tree node(0x45bf)                      page_num = 17       page_prev = 28       page_next = 27       record_count = 638      tree_level = 0        index_id = 2721
page_type = B-tree node(0x45bf)                      page_num = 18       page_prev = 11       page_next = 20       record_count = 211      tree_level = 0        index_id = 2719
page_type = B-tree node(0x45bf)                      page_num = 19       page_prev = 9        page_next = 11       record_count = 212      tree_level = 0        index_id = 2719
page_type = B-tree node(0x45bf)                      page_num = 20       page_prev = 18       page_next = 24       record_count = 210      tree_level = 0        index_id = 2719
page_type = B-tree node(0x45bf)                      page_num = 21       page_prev = 34       page_next = 35       record_count = 303      tree_level = 0        index_id = 2720
page_type = B-tree node(0x45bf)                      page_num = 22       page_prev = 32       page_next = 33       record_count = 309      tree_level = 0        index_id = 2720
page_type = B-tree node(0x45bf)                      page_num = 23       page_prev = 37       page_next = 39       record_count = 304      tree_level = 0        index_id = 2722
page_type = B-tree node(0x45bf)                      page_num = 24       page_prev = 20       page_next = 26       record_count = 210      tree_level = 0        index_id = 2719
page_type = B-tree node(0x45bf)                      page_num = 25       page_prev = 38       page_next = 40       record_count = 299      tree_level = 0        index_id = 2722
page_type = B-tree node(0x45bf)                      page_num = 26       page_prev = 24       page_next = 29       record_count = 211      tree_level = 0        index_id = 2719
page_type = B-tree node(0x45bf)                      page_num = 27       page_prev = 17       page_next = -1       record_count = 640      tree_level = 0        index_id = 2721
page_type = B-tree node(0x45bf)                      page_num = 28       page_prev = 16       page_next = 17       record_count = 580      tree_level = 0        index_id = 2721
page_type = B-tree node(0x45bf)                      page_num = 29       page_prev = 26       page_next = 30       record_count = 210      tree_level = 0        index_id = 2719
page_type = B-tree node(0x45bf)                      page_num = 30       page_prev = 29       page_next = 31       record_count = 210      tree_level = 0        index_id = 2719
page_type = B-tree node(0x45bf)                      page_num = 31       page_prev = 30       page_next = 36       record_count = 210      tree_level = 0        index_id = 2719
page_type = B-tree node(0x45bf)                      page_num = 32       page_prev = 13       page_next = 22       record_count = 306      tree_level = 0        index_id = 2720
page_type = B-tree node(0x45bf)                      page_num = 33       page_prev = 22       page_next = -1       record_count = 309      tree_level = 0        index_id = 2720
page_type = B-tree node(0x45bf)                      page_num = 34       page_prev = 12       page_next = 21       record_count = 301      tree_level = 0        index_id = 2720
page_type = B-tree node(0x45bf)                      page_num = 35       page_prev = 21       page_next = 13       record_count = 288      tree_level = 0        index_id = 2720
page_type = B-tree node(0x45bf)                      page_num = 36       page_prev = 31       page_next = -1       record_count = 203      tree_level = 0        index_id = 2719
page_type = B-tree node(0x45bf)                      page_num = 37       page_prev = 14       page_next = 23       record_count = 310      tree_level = 0        index_id = 2722
page_type = B-tree node(0x45bf)                      page_num = 38       page_prev = 15       page_next = 25       record_count = 301      tree_level = 0        index_id = 2722
page_type = B-tree node(0x45bf)                      page_num = 39       page_prev = 23       page_next = 15       record_count = 301      tree_level = 0        index_id = 2722
page_type = B-tree node(0x45bf)                      page_num = 40       page_prev = 25       page_next = -1       record_count = 299      tree_level = 0        index_id = 2722
page_type = Freshly allocated page(0x0000)           page_num = 0        page_prev = 0        page_next = 0        record_count = 0        tree_level = 0        index_id = 0
page_type = Freshly allocated page(0x0000)           page_num = 0        page_prev = 0        page_next = 0        record_count = 0        tree_level = 0        index_id = 0
```
##### KR Version
```
$ ./main test.ibd 1

- InnoDB tablespace file: test0.ibd
- Pages: 43

page_type = 테이블 스페이스 헤더(0x0008)                      page_num = 0        page_prev = 80300    page_next = 1        record_count = 0        tree_level = 0        index_id = 18446744069414649855
page_type = 체인지 버퍼 비트맵(0x0005)                       page_num = 1        page_prev = 0        page_next = 0        record_count = 0        tree_level = 0        index_id = 0
page_type = 인덱스 노드(0x0003)                           page_num = 2        page_prev = 0        page_next = 0        record_count = 0        tree_level = 0        index_id = 18446744069414649855
page_type = 테이블스페이스 SDI 인덱스 페이지(0x45bd)              page_num = 3        page_prev = -1       page_next = -1       record_count = 2        tree_level = 0        index_id = 18446744073709551615
page_type = 인덱스 페이지(0x45bf)                          page_num = 4        page_prev = -1       page_next = -1       record_count = 13       tree_level = 1        index_id = 2719
page_type = 인덱스 페이지(0x45bf)                          page_num = 5        page_prev = -1       page_next = -1       record_count = 8        tree_level = 1        index_id = 2720
page_type = 인덱스 페이지(0x45bf)                          page_num = 6        page_prev = -1       page_next = -1       record_count = 4        tree_level = 1        index_id = 2721
page_type = 인덱스 페이지(0x45bf)                          page_num = 7        page_prev = -1       page_next = -1       record_count = 8        tree_level = 1        index_id = 2722
page_type = 인덱스 페이지(0x45bf)                          page_num = 8        page_prev = -1       page_next = 10       record_count = 105      tree_level = 0        index_id = 2719
page_type = 인덱스 페이지(0x45bf)                          page_num = 9        page_prev = 10       page_next = 19       record_count = 210      tree_level = 0        index_id = 2719
page_type = 인덱스 페이지(0x45bf)                          page_num = 10       page_prev = 8        page_next = 9        record_count = 208      tree_level = 0        index_id = 2719
page_type = 인덱스 페이지(0x45bf)                          page_num = 11       page_prev = 19       page_next = 18       record_count = 3        tree_level = 0        index_id = 2719
page_type = 인덱스 페이지(0x45bf)                          page_num = 12       page_prev = -1       page_next = 34       record_count = 285      tree_level = 0        index_id = 2720
page_type = 인덱스 페이지(0x45bf)                          page_num = 13       page_prev = 35       page_next = 32       record_count = 312      tree_level = 0        index_id = 2720
page_type = 인덱스 페이지(0x45bf)                          page_num = 14       page_prev = -1       page_next = 37       record_count = 301      tree_level = 0        index_id = 2722
page_type = 인덱스 페이지(0x45bf)                          page_num = 15       page_prev = 39       page_next = 38       record_count = 298      tree_level = 0        index_id = 2722
page_type = 인덱스 페이지(0x45bf)                          page_num = 16       page_prev = -1       page_next = 28       record_count = 555      tree_level = 0        index_id = 2721
page_type = 인덱스 페이지(0x45bf)                          page_num = 17       page_prev = 28       page_next = 27       record_count = 638      tree_level = 0        index_id = 2721
page_type = 인덱스 페이지(0x45bf)                          page_num = 18       page_prev = 11       page_next = 20       record_count = 211      tree_level = 0        index_id = 2719
page_type = 인덱스 페이지(0x45bf)                          page_num = 19       page_prev = 9        page_next = 11       record_count = 212      tree_level = 0        index_id = 2719
page_type = 인덱스 페이지(0x45bf)                          page_num = 20       page_prev = 18       page_next = 24       record_count = 210      tree_level = 0        index_id = 2719
page_type = 인덱스 페이지(0x45bf)                          page_num = 21       page_prev = 34       page_next = 35       record_count = 303      tree_level = 0        index_id = 2720
page_type = 인덱스 페이지(0x45bf)                          page_num = 22       page_prev = 32       page_next = 33       record_count = 309      tree_level = 0        index_id = 2720
page_type = 인덱스 페이지(0x45bf)                          page_num = 23       page_prev = 37       page_next = 39       record_count = 304      tree_level = 0        index_id = 2722
page_type = 인덱스 페이지(0x45bf)                          page_num = 24       page_prev = 20       page_next = 26       record_count = 210      tree_level = 0        index_id = 2719
page_type = 인덱스 페이지(0x45bf)                          page_num = 25       page_prev = 38       page_next = 40       record_count = 299      tree_level = 0        index_id = 2722
page_type = 인덱스 페이지(0x45bf)                          page_num = 26       page_prev = 24       page_next = 29       record_count = 211      tree_level = 0        index_id = 2719
page_type = 인덱스 페이지(0x45bf)                          page_num = 27       page_prev = 17       page_next = -1       record_count = 640      tree_level = 0        index_id = 2721
page_type = 인덱스 페이지(0x45bf)                          page_num = 28       page_prev = 16       page_next = 17       record_count = 580      tree_level = 0        index_id = 2721
page_type = 인덱스 페이지(0x45bf)                          page_num = 29       page_prev = 26       page_next = 30       record_count = 210      tree_level = 0        index_id = 2719
page_type = 인덱스 페이지(0x45bf)                          page_num = 30       page_prev = 29       page_next = 31       record_count = 210      tree_level = 0        index_id = 2719
page_type = 인덱스 페이지(0x45bf)                          page_num = 31       page_prev = 30       page_next = 36       record_count = 210      tree_level = 0        index_id = 2719
page_type = 인덱스 페이지(0x45bf)                          page_num = 32       page_prev = 13       page_next = 22       record_count = 306      tree_level = 0        index_id = 2720
page_type = 인덱스 페이지(0x45bf)                          page_num = 33       page_prev = 22       page_next = -1       record_count = 309      tree_level = 0        index_id = 2720
page_type = 인덱스 페이지(0x45bf)                          page_num = 34       page_prev = 12       page_next = 21       record_count = 301      tree_level = 0        index_id = 2720
page_type = 인덱스 페이지(0x45bf)                          page_num = 35       page_prev = 21       page_next = 13       record_count = 288      tree_level = 0        index_id = 2720
page_type = 인덱스 페이지(0x45bf)                          page_num = 36       page_prev = 31       page_next = -1       record_count = 203      tree_level = 0        index_id = 2719
page_type = 인덱스 페이지(0x45bf)                          page_num = 37       page_prev = 14       page_next = 23       record_count = 310      tree_level = 0        index_id = 2722
page_type = 인덱스 페이지(0x45bf)                          page_num = 38       page_prev = 15       page_next = 25       record_count = 301      tree_level = 0        index_id = 2722
page_type = 인덱스 페이지(0x45bf)                          page_num = 39       page_prev = 23       page_next = 15       record_count = 301      tree_level = 0        index_id = 2722
page_type = 인덱스 페이지(0x45bf)                          page_num = 40       page_prev = 25       page_next = -1       record_count = 299      tree_level = 0        index_id = 2722
page_type = 새로 할당된 페이지(0x0000)                       page_num = 0        page_prev = 0        page_next = 0        record_count = 0        tree_level = 0        index_id = 0
page_type = 새로 할당된 페이지(0x0000)                       page_num = 0        page_prev = 0        page_next = 0        record_count = 0        tree_level = 0        index_id = 0
```
